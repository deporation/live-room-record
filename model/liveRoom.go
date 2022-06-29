package model

// Message 弹幕消息
type Message struct {
	Mode           int                    //弹幕显示模式（滚动、顶部、底部）
	FontSize       int                    //字体尺寸
	Color          int                    //颜色
	Timestamp      int                    //时间戳（毫秒）
	Rnd            int                    //随机数，前端叫作弹幕ID，可能是去重用的
	UidCrc32       int                    //用户ID文本的CRC32
	MsgType        int                    //是否礼物弹幕（节奏风暴）
	Bubble         int                    //右侧评论栏气泡
	DmType         int                    //弹幕类型，0文本，1表情，2语音
	EmoticoOptions interface{}            //表情参数
	VoiceConfig    interface{}            //语音参数
	ModeInfo       map[string]interface{} //一些附加参数

	Msg string //弹幕内容

	Uid          int    //用户ID
	Uname        string //用户名
	Admin        int    //是否房管
	Vip          int    //是否月费老爷
	Svip         int    //是否年费老爷
	Urank        int    //用户身份，用来判断是否正式会员，猜测非正式会员为5000，正式会员为10000
	MobileVerify int    //是否绑定手机
	UnameColor   string //用户名颜色

	MedalLevel   string //勋章等级
	MedalName    string //勋章名
	Runame       string //勋章房间主播名
	MedalRoomId  int    //勋章房间ID
	Mcolor       int    //勋章颜色
	SpecialMedal string //特殊勋章

	UserLevel   int    //用户等级
	UlevelColor int    //用户等级颜色
	UlevelRank  string //用户等级排名，>50000时为'>50000'

	OldTiltle string //旧头衔
	Title     string //头衔

	PrivilegeType int //舰队类型，0非舰队，1总督，2提督，3舰长
}

// GiftMessage 礼物消息
type GiftMessage struct {
	GiftName   string //礼物名
	Num        int    //数量
	Uname      string //用户名
	Face       string //用户头像url
	GuardLevel int    //舰队等级，0非舰队，1总督，2提督，3舰长
	Uid        int    //用户ID
	Timestamp  int    //时间戳
	GiftId     int    //礼物ID
	GiftType   int    //礼物类型（未知）
	Action     string //目前遇到的有'喂食'、'赠送'
	Price      int    //礼物单价瓜子数
	Rnd        string //随机数，可能是去重用的。有时是时间戳+去重ID，有时是UUID
	CoinType   string //瓜子类型，'silver'或'gold'，1000金瓜子 = 1元
	TotalCoin  int    //总瓜子数
	Tid        string //可能是事务ID，有时和rnd相同
}

// GuardBuyMessage 上舰消息
type GuardBuyMessage struct {
	Uid        int    //用户id
	UserName   string //用户名
	GuardLevel int    //舰队等级，0非舰队，1总督，2提督，3舰长
	Num        int    //数量
	Price      int    //单价金瓜子数
	GiftId     int    //礼物id
	GiftName   string //礼物名
	StartTime  int    //开始时间戳
	EndTime    int    //结束时间戳
}

// SuperChatMessage sc醒目留言
type SuperChatMessage struct {
	Price                 int    `json:"price"`                   //price: 价格（人民币）
	Meaage                string `json:"meaage"`                  //消息
	MessageTrams          string `json:"message_trams"`           //消息日文翻译（目前只出现在SUPER_CHAT_MESSAGE_JPN）
	StarTime              int    `json:"star_time"`               //开始时间戳
	EndTime               int    `json:"end_time"`                //结束时间戳
	Time                  int    `json:"time"`                    //剩余时间（约等于 结束时间戳 - 开始时间戳）
	Id                    int    `json:"id"`                      //醒目留言ID，删除时用
	GiftId                int    `json:"gift_id"`                 //礼物id
	GiftName              string `json:"gift_name"`               //礼物名
	Uid                   int    `json:"uid"`                     //用户id
	Uname                 string `json:"uname"`                   //用户名
	Face                  string `json:"face"`                    //用户头像url
	GuardLevel            int    `json:"guard_level"`             //舰队等级，0非舰队，1总督，2提督，3舰长
	UserLevel             int    `json:"user_level"`              //用户等级
	BackgroundBottomColor string `json:"background_bottom_color"` //底部颜色
	BackgroundColor       string `json:"background_color"`        //背景色
	BackgroundIcon        string `json:"background_icon"`         //背景图标
	BackgroundImage       string `json:"background_image"`        //背景图url
	BackgroundPriceColor  string `json:"background_price_color"`  //背景价格颜色
}

//Certificate 认证包
type Certificate struct {
	Uid      int    `json:"uid"`
	RoomId   int64  `json:"room_id"`
	Protover int    `json:"protover"`
	Platform string `json:"platform"`
	Type     int    `json:"type"`
	Key      string `json:"key"`
}

// HeartbeatMessage 心跳消息
type HeartbeatMessage struct {
	Popularity int32 //人气
}

type Auth struct {
	UID       uint8  `json:"uid"`
	Roomid    uint32 `json:"roomid"`
	Protover  uint8  `json:"protover"`
	Platform  string `json:"platform"`
	Clientver string `json:"clientver"`
	Type      uint8  `json:"type"`
	Key       string `json:"key"`
}
