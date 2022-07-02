package model

import "github.com/tidwall/gjson"

const (
	RoomId = 411318
)

func ParaseMessage(src []gjson.Result) *Message {
	var message Message
	base := src[0].Array()
	message.Mode = int(base[1].Int())
	message.FontSize = int(base[2].Int())
	message.Color = int(base[3].Int())
	message.Timestamp = base[4].Int()
	message.Rnd = int(base[5].Int())
	message.UidCrc32 = int(base[7].Int())
	message.MsgType = int(base[9].Int())
	message.Bubble = int(base[10].Int())
	message.DmType = int(base[12].Int())
	message.EmoticoOptions = int(base[13].Int())
	message.VoiceConfig = int(base[14].Int())
	message.ModeInfo = base[15].Map()

	message.Msg = src[1].String()

	userInfo := src[2].Array()
	message.Uid = int(userInfo[0].Int())
	message.Uname = userInfo[1].String()
	message.Admin = int(userInfo[2].Int())
	message.Vip = int(userInfo[3].Int())
	message.Svip = int(userInfo[4].Int())
	message.Urank = int(userInfo[5].Int())
	message.MobileVerify = int(userInfo[6].Int())
	message.UnameColor = userInfo[7].String()

	medal := src[3].Array()
	if len(medal) > 0 {
		message.MedalLevel = medal[0].String()
		message.MedalName = medal[1].String()
		message.Runame = medal[2].String()
		message.MedalRoomId = int(medal[3].Int())
		message.Mcolor = int(medal[4].Int())
		message.SpecialMedal = medal[5].String()
	}

	message.UserLevel = int(src[4].Array()[0].Int())
	message.UlevelColor = int(src[4].Array()[2].Int())
	message.UlevelRank = src[4].Array()[3].String()

	message.OldTiltle = src[5].Array()[0].String()
	message.Title = src[5].Array()[1].String()

	message.PrivilegeType = int(src[7].Int())

	return &message
}

// Message 弹幕消息
type Message struct {
	Mode           int                     //弹幕显示模式（滚动、顶部、底部）
	FontSize       int                     //字体尺寸
	Color          int                     //颜色
	Timestamp      int64                   //时间戳（毫秒）
	Rnd            int                     //随机数，前端叫作弹幕ID，可能是去重用的
	UidCrc32       int                     //用户ID文本的CRC32
	MsgType        int                     //是否礼物弹幕（节奏风暴）
	Bubble         int                     //右侧评论栏气泡
	DmType         int                     //弹幕类型，0文本，1表情，2语音
	EmoticoOptions interface{}             //表情参数
	VoiceConfig    interface{}             //语音参数
	ModeInfo       map[string]gjson.Result //一些附加参数

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

//GiftMessage 礼物消息
type GiftMessage struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Draw              int         `json:"draw"`
		Gold              int         `json:"gold"`
		Silver            int         `json:"silver"`
		Num               int         `json:"num"`        //数量
		TotalCoin         int         `json:"total_coin"` //总瓜子数
		Effect            int         `json:"effect"`
		BroadcastID       int         `json:"broadcast_id"`
		CritProb          int         `json:"crit_prob"`
		GuardLevel        int         `json:"guard_level"` //舰队等级，0非舰队，1总督，2提督，3舰长
		Rcost             int         `json:"rcost"`
		UID               int         `json:"uid"`       //用户ID
		Timestamp         int         `json:"timestamp"` //时间戳
		GiftID            int         `json:"giftId"`    //礼物ID
		GiftType          int         `json:"giftType"`  //礼物类型（未知）
		Super             int         `json:"super"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		Remain            int         `json:"remain"`
		Price             int         `json:"price"` //礼物单价瓜子数
		BeatID            string      `json:"beatId"`
		BizSource         string      `json:"biz_source"`
		Action            string      `json:"action"`    //目前遇到的有'喂食'、'赠送'
		CoinType          string      `json:"coin_type"` //瓜子类型，'silver'或'gold'，1000金瓜子 = 1元
		Uname             string      `json:"uname"`     //用户名
		Face              string      `json:"face"`      //用户头像url
		BatchComboID      string      `json:"batch_combo_id"`
		Rnd               string      `json:"rnd"`      //随机数，可能是去重用的。有时是时间戳+去重ID，有时是UUID
		GiftName          string      `json:"giftName"` //礼物名
		ComboSend         interface{} `json:"combo_send"`
		BatchComboSend    interface{} `json:"batch_combo_send"`
		TagImage          string      `json:"tag_image"`
		TopList           interface{} `json:"top_list"`
		SendMaster        interface{} `json:"send_master"`
		IsFirst           bool        `json:"is_first"`
		Demarcation       int         `json:"demarcation"`
		ComboStayTime     int         `json:"combo_stay_time"`
		ComboTotalCoin    int         `json:"combo_total_coin"`
		Tid               string      `json:"tid"` //可能是事务ID，有时和rnd相同
		EffectBlock       int         `json:"effect_block"`
		IsSpecialBatch    int         `json:"is_special_batch"`
		ComboResourcesID  int         `json:"combo_resources_id"`
		Magnification     float64     `json:"magnification"`
		NameColor         string      `json:"name_color"`
		MedalInfo         struct {
			TargetID         int    `json:"target_id"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"` //舰队等级，0非舰队，1总督，2提督，3舰长
		} `json:"medal_info"`
		SvgaBlock int `json:"svga_block"`
	} `json:"data"`
}

type InteractWord struct {
	Cmd string `json:"cmd"`

	Data struct {
		UID        int    `json:"uid"`
		Uname      string `json:"uname"`
		UnameColor string `json:"uname_color"`
		Identities []int  `json:"identities"`
		MsgType    int    `json:"msg_type"`
		Roomid     int    `json:"roomid"`
		Timestamp  int    `json:"timestamp"`
		Score      int64  `json:"score"`
		FansMedal  struct {
			TargetID         int    `json:"target_id"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			Score            int    `json:"score"`
		} `json:"fans_medal"`
		IsSpread     int    `json:"is_spread"`
		SpreadInfo   string `json:"spread_info"`
		Contribution struct {
			Grade int `json:"grade"`
		} `json:"contribution"`
		SpreadDesc string `json:"spread_desc"`
		TailIcon   int    `json:"tail_icon"`
	} `json:"data"`
}

type ComboSend struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID           int         `json:"uid"`
		Ruid          int         `json:"ruid"`
		Uname         string      `json:"uname"`
		RUname        string      `json:"r_uname"`
		ComboNum      int         `json:"combo_num"`
		GiftID        int         `json:"gift_id"`
		GiftNum       int         `json:"gift_num"`
		BatchComboNum int         `json:"batch_combo_num"`
		GiftName      string      `json:"gift_name"`
		Action        string      `json:"action"`
		ComboID       string      `json:"combo_id"`
		BatchComboID  string      `json:"batch_combo_id"`
		IsShow        int         `json:"is_show"`
		SendMaster    interface{} `json:"send_master"`
		NameColor     string      `json:"name_color"`
		TotalNum      int         `json:"total_num"`
		MedalInfo     struct {
			TargetID         int    `json:"target_id"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
		} `json:"medal_info"`
		ComboTotalCoin int `json:"combo_total_coin"`
	} `json:"data"`
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
	Cmd  string `json:"cmd"`
	Data struct {
		BackgroundBottomColor string  `json:"background_bottom_color"`
		BackgroundColor       string  `json:"background_color"`
		BackgroundColorEnd    string  `json:"background_color_end"`
		BackgroundColorStart  string  `json:"background_color_start"`
		BackgroundIcon        string  `json:"background_icon"`
		BackgroundImage       string  `json:"background_image"`
		BackgroundPriceColor  string  `json:"background_price_color"`
		ColorPoint            float64 `json:"color_point"`
		Dmscore               int     `json:"dmscore"`
		EndTime               int     `json:"end_time"`
		Gift                  struct {
			GiftId   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			Num      int    `json:"num"`
		} `json:"gift"`
		Id          int `json:"id"`
		IsRanked    int `json:"is_ranked"`
		IsSendAudit int `json:"is_send_audit"`
		MedalInfo   struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconId           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       string `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetId         int    `json:"target_id"`
		} `json:"medal_info"`
		Message          string `json:"message"`
		MessageFontColor string `json:"message_font_color"`
		MessageTrans     string `json:"message_trans"`
		Price            int    `json:"price"`
		Rate             int    `json:"rate"`
		StartTime        int    `json:"start_time"`
		Time             int    `json:"time"`
		Token            string `json:"token"`
		TransMark        int    `json:"trans_mark"`
		Ts               int    `json:"ts"`
		Uid              int    `json:"uid"`
		UserInfo         struct {
			Face       string `json:"face"`
			FaceFrame  string `json:"face_frame"`
			GuardLevel int    `json:"guard_level"`
			IsMainVip  int    `json:"is_main_vip"`
			IsSvip     int    `json:"is_svip"`
			IsVip      int    `json:"is_vip"`
			LevelColor string `json:"level_color"`
			Manager    int    `json:"manager"`
			NameColor  string `json:"name_color"`
			Title      string `json:"title"`
			Uname      string `json:"uname"`
			UserLevel  int    `json:"user_level"`
		} `json:"user_info"`
	} `json:"data"`
	Roomid int `json:"roomid"`
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
