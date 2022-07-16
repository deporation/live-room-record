package orm

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoClient struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (m *MongoClient) Connect(ip string, port int, baseName string, collectionName string) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/?connect=direct;authSource=admin", ip, port))
	connect, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
		return
	}
	m.Client = connect
	collection := m.Client.Database(baseName).Collection(collectionName)
	m.Collection = collection
}

func (m *MongoClient) InsertedOne(ctx context.Context, c string, e interface{}) (error, interface{}) {
	var err error

	cid, err := m.Collection.InsertOne(ctx, e)
	if err == nil {
		return nil, cid.InsertedID
	}

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return err, 0
}
func (m *MongoClient) UpdateOne(ctx context.Context, c string, f interface{}, s interface{}) (int64, error) {
	var err error
	cid, err := m.Collection.UpdateOne(ctx, f, s)
	if err == nil {
		return cid.MatchedCount, nil
	}

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return 0, err
}
func (m *MongoClient) UpdateOneUpsert(ctx context.Context, c string, f interface{}, s interface{}) (int64, error) {
	var err error
	cid, err := m.Collection.UpdateOne(ctx, f, s, options.Update().SetUpsert(true))
	if err == nil {
		return cid.MatchedCount, nil
	}

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return 0, err
}
func (m *MongoClient) UpdateAll(ctx context.Context, c string, f interface{}, s interface{}) (err error) {
	_, err = m.Collection.UpdateMany(ctx, f, s)

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return
}
func (m *MongoClient) FindOneAndUpdate(ctx context.Context, c string, f interface{}, s interface{}) (result *mongo.SingleResult) {
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result = m.Collection.FindOneAndUpdate(ctx, f, s, &opt)

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return result
}
func (m *MongoClient) FindOne(ctx context.Context, c string, f interface{}) (result *mongo.SingleResult) {
	result = m.Collection.FindOne(ctx, f)

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return result
}

func (m *MongoClient) Find(ctx context.Context, c string, f interface{}) (result *mongo.Cursor, err error) {
	result, err = m.Collection.Find(ctx, f)

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return result, err
}
func (m *MongoClient) FindPaging(ctx context.Context, c string, skip int64, limit int64, sort map[string]interface{}, f interface{}) (err error, result *mongo.Cursor) {
	result, err = m.Collection.Find(ctx, f, options.Find().SetSkip(skip), options.Find().SetLimit(limit), options.Find().SetSort(sort))

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return err, result
}
func (m *MongoClient) Count(ctx context.Context, c string, f interface{}) (err error, result int64) {
	if f == nil {
		f = bson.M{}
	}
	result, err = m.Collection.CountDocuments(ctx, f)

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return err, result
}
func (m *MongoClient) GenerateNextId(seqName string) int64 {
	seqData := make(map[string]int64)
	err, count := m.Count(context.TODO(), "seq", nil)
	if count == 0 {
		seqData[seqName] = 1
		seqData["_id"] = 1 // 自定义 id 容易接数据
		err, _ := m.InsertedOne(context.TODO(), "seq", seqData)
		if err == nil {
			return seqData[seqName]
		}
	}
	result := m.FindOne(context.TODO(), "seq", bson.M{})
	err = result.Decode(&seqData)
	if err != nil {
		fmt.Println(err)
	}
	capital, ok := seqData[seqName]
	if ok {
		seqData[seqName] = capital + 1
		MatchedCount, err := m.UpdateOne(context.TODO(), "seq", bson.M{}, bson.M{"$set": seqData})
		if err == nil && MatchedCount == 1 {
			return seqData[seqName]
		}
	} else {
		seqData[seqName] = 1
		MatchedCount, err := m.UpdateOne(context.TODO(), "seq", bson.M{}, bson.M{"$set": seqData})
		if err == nil && MatchedCount == 1 {
			return seqData[seqName]
		}
	}

	defer func(Client *mongo.Client) {
		err := Client.Disconnect(context.TODO())
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client)

	return 0
}

func (m *MongoClient) CLose(ctx context.Context) error {
	err := m.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}(m.Client, ctx)

	return nil
}
