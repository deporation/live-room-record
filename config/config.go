package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

type Config struct {
	LiveRoom []int `yaml:"live_room"`

	Momngo struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"momngo"`
}

var conf *Config
var once sync.Once

func getConfig() *Config {
	viperLocal := viper.New()
	work, _ := os.Getwd()
	viperLocal.AddConfigPath(work + "/config")
	viperLocal.SetConfigName("env")
	viperLocal.SetConfigType("yaml")
	err := viperLocal.ReadInConfig() //找到并读取配置文件
	if err != nil {                  // 捕获读取中遇到的error
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	//var config Config
	room := viperLocal.Get("live_room")
	log.Println(room)
	err = viperLocal.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	interfaces := room.([]interface{})
	for _, val := range interfaces {
		conf.LiveRoom = append(conf.LiveRoom, val.(int))
	}
	return conf
}

func GetInstance() *Config {
	once.Do(func() {
		getConfig()
	})
	return conf
}
