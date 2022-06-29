package model

type BiRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Group            string  `json:"group"`
		BusinessId       int     `json:"business_id"`
		RefreshRowFactor float64 `json:"refresh_row_factor"`
		RefreshRate      int     `json:"refresh_rate"`
		MaxDelay         int     `json:"max_delay"`
		Token            string  `json:"token"`
		HostList         []struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			WssPort int    `json:"wss_port"`
			WsPort  int    `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

type RoomInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		RoomInfo struct {
			Uid            int    `json:"uid"`
			RoomId         int    `json:"room_id"`
			ShortId        int    `json:"short_id"`
			Title          string `json:"title"`
			Cover          string `json:"cover"`
			Tags           string `json:"tags"`
			Background     string `json:"background"`
			Description    string `json:"description"`
			LiveStatus     int    `json:"live_status"`
			LiveStartTime  int    `json:"live_start_time"`
			LiveScreenType int    `json:"live_screen_type"`
			LockStatus     int    `json:"lock_status"`
			LockTime       int    `json:"lock_time"`
			HiddenStatus   int    `json:"hidden_status"`
			HiddenTime     int    `json:"hidden_time"`
			AreaId         int    `json:"area_id"`
			AreaName       string `json:"area_name"`
			ParentAreaId   int    `json:"parent_area_id"`
			ParentAreaName string `json:"parent_area_name"`
			Keyframe       string `json:"keyframe"`
			SpecialType    int    `json:"special_type"`
			UpSession      string `json:"up_session"`
			PkStatus       int    `json:"pk_status"`
			IsStudio       bool   `json:"is_studio"`
			Pendants       struct {
				Frame struct {
					Name  string `json:"name"`
					Value string `json:"value"`
					Desc  string `json:"desc"`
				} `json:"frame"`
			} `json:"pendants"`
			OnVoiceJoin int `json:"on_voice_join"`
			Online      int `json:"online"`
			RoomType    struct {
				Field1 int `json:"3-21"`
			} `json:"room_type"`
		} `json:"room_info"`
		AnchorInfo struct {
			BaseInfo struct {
				Uname        string `json:"uname"`
				Face         string `json:"face"`
				Gender       string `json:"gender"`
				OfficialInfo struct {
					Role     int    `json:"role"`
					Title    string `json:"title"`
					Desc     string `json:"desc"`
					IsNft    int    `json:"is_nft"`
					NftDmark string `json:"nft_dmark"`
				} `json:"official_info"`
			} `json:"base_info"`
			LiveInfo struct {
				Level        int    `json:"level"`
				LevelColor   int    `json:"level_color"`
				Score        int    `json:"score"`
				UpgradeScore int    `json:"upgrade_score"`
				Current      []int  `json:"current"`
				Next         []int  `json:"next"`
				Rank         string `json:"rank"`
			} `json:"live_info"`
			RelationInfo struct {
				Attention int `json:"attention"`
			} `json:"relation_info"`
			MedalInfo interface{} `json:"medal_info"`
		} `json:"anchor_info"`
		NewsInfo struct {
			Uid     int    `json:"uid"`
			Ctime   string `json:"ctime"`
			Content string `json:"content"`
		} `json:"news_info"`
		RankdbInfo struct {
			Roomid    int    `json:"roomid"`
			RankDesc  string `json:"rank_desc"`
			Color     string `json:"color"`
			H5Url     string `json:"h5_url"`
			WebUrl    string `json:"web_url"`
			Timestamp int    `json:"timestamp"`
		} `json:"rankdb_info"`
		AreaRankInfo struct {
			AreaRank struct {
				Index int    `json:"index"`
				Rank  string `json:"rank"`
			} `json:"areaRank"`
			LiveRank struct {
				Rank string `json:"rank"`
			} `json:"liveRank"`
		} `json:"area_rank_info"`
		BattleRankEntryInfo struct {
			FirstRankImgUrl string `json:"first_rank_img_url"`
			RankName        string `json:"rank_name"`
			ShowStatus      int    `json:"show_status"`
		} `json:"battle_rank_entry_info"`
		TabInfo struct {
			List []struct {
				Type      string `json:"type"`
				Desc      string `json:"desc"`
				IsFirst   int    `json:"isFirst"`
				IsEvent   int    `json:"isEvent"`
				EventType string `json:"eventType"`
				ListType  string `json:"listType"`
				ApiPrefix string `json:"apiPrefix"`
				RankName  string `json:"rank_name"`
			} `json:"list"`
		} `json:"tab_info"`
		ActivityInitInfo struct {
			EventList []interface{} `json:"eventList"`
			WeekInfo  struct {
				BannerInfo interface{} `json:"bannerInfo"`
				GiftName   interface{} `json:"giftName"`
			} `json:"weekInfo"`
			GiftName interface{} `json:"giftName"`
			Lego     struct {
				Timestamp int    `json:"timestamp"`
				Config    string `json:"config"`
			} `json:"lego"`
		} `json:"activity_init_info"`
		VoiceJoinInfo struct {
			Status struct {
				Open        int    `json:"open"`
				AnchorOpen  int    `json:"anchor_open"`
				Status      int    `json:"status"`
				Uid         int    `json:"uid"`
				UserName    string `json:"user_name"`
				HeadPic     string `json:"head_pic"`
				Guard       int    `json:"guard"`
				StartAt     int    `json:"start_at"`
				CurrentTime int    `json:"current_time"`
			} `json:"status"`
			Icons struct {
				IconClose    string `json:"icon_close"`
				IconOpen     string `json:"icon_open"`
				IconWait     string `json:"icon_wait"`
				IconStarting string `json:"icon_starting"`
			} `json:"icons"`
			WebShareLink string `json:"web_share_link"`
		} `json:"voice_join_info"`
		AdBannerInfo struct {
			Data []struct {
				Id                   int         `json:"id"`
				Title                string      `json:"title"`
				Location             string      `json:"location"`
				Position             int         `json:"position"`
				Pic                  string      `json:"pic"`
				Link                 string      `json:"link"`
				Weight               int         `json:"weight"`
				RoomId               int         `json:"room_id"`
				UpId                 int         `json:"up_id"`
				ParentAreaId         int         `json:"parent_area_id"`
				AreaId               int         `json:"area_id"`
				LiveStatus           int         `json:"live_status"`
				AvId                 int         `json:"av_id"`
				IsAd                 bool        `json:"is_ad"`
				AdTransparentContent interface{} `json:"ad_transparent_content"`
				ShowAdIcon           bool        `json:"show_ad_icon"`
			} `json:"data"`
		} `json:"ad_banner_info"`
		SkinInfo struct {
			Id          int    `json:"id"`
			SkinName    string `json:"skin_name"`
			SkinConfig  string `json:"skin_config"`
			ShowText    string `json:"show_text"`
			SkinUrl     string `json:"skin_url"`
			StartTime   int    `json:"start_time"`
			EndTime     int    `json:"end_time"`
			CurrentTime int    `json:"current_time"`
		} `json:"skin_info"`
		WebBannerInfo struct {
			Id               int    `json:"id"`
			Title            string `json:"title"`
			Left             string `json:"left"`
			Right            string `json:"right"`
			JumpUrl          string `json:"jump_url"`
			BgColor          string `json:"bg_color"`
			HoverColor       string `json:"hover_color"`
			TextBgColor      string `json:"text_bg_color"`
			TextHoverColor   string `json:"text_hover_color"`
			LinkText         string `json:"link_text"`
			LinkColor        string `json:"link_color"`
			InputColor       string `json:"input_color"`
			InputTextColor   string `json:"input_text_color"`
			InputHoverColor  string `json:"input_hover_color"`
			InputBorderColor string `json:"input_border_color"`
			InputSearchColor string `json:"input_search_color"`
		} `json:"web_banner_info"`
		LolInfo struct {
			LolActivity struct {
				Status     int    `json:"status"`
				GuessCover string `json:"guess_cover"`
				VoteCover  string `json:"vote_cover"`
				VoteH5Url  string `json:"vote_h5_url"`
				VoteUseH5  bool   `json:"vote_use_h5"`
			} `json:"lol_activity"`
		} `json:"lol_info"`
		PkInfo         interface{} `json:"pk_info"`
		BattleInfo     interface{} `json:"battle_info"`
		SilentRoomInfo struct {
			Type       string `json:"type"`
			Level      int    `json:"level"`
			Second     int    `json:"second"`
			ExpireTime int    `json:"expire_time"`
		} `json:"silent_room_info"`
		SwitchInfo struct {
			CloseGuard   bool `json:"close_guard"`
			CloseGift    bool `json:"close_gift"`
			CloseOnline  bool `json:"close_online"`
			CloseDanmaku bool `json:"close_danmaku"`
		} `json:"switch_info"`
		RecordSwitchInfo struct {
			RecordTab bool `json:"record_tab"`
		} `json:"record_switch_info"`
		RoomConfigInfo struct {
			DmText string `json:"dm_text"`
		} `json:"room_config_info"`
		GiftMemoryInfo struct {
			List interface{} `json:"list"`
		} `json:"gift_memory_info"`
		NewSwitchInfo struct {
			RoomSocket           int `json:"room-socket"`
			RoomPropSend         int `json:"room-prop-send"`
			RoomSailing          int `json:"room-sailing"`
			RoomInfoPopularity   int `json:"room-info-popularity"`
			RoomDanmakuEditor    int `json:"room-danmaku-editor"`
			RoomEffect           int `json:"room-effect"`
			RoomFansMedal        int `json:"room-fans_medal"`
			RoomReport           int `json:"room-report"`
			RoomFeedback         int `json:"room-feedback"`
			RoomPlayerWatermark  int `json:"room-player-watermark"`
			RoomRecommendLiveOff int `json:"room-recommend-live_off"`
			RoomActivity         int `json:"room-activity"`
			RoomWebBanner        int `json:"room-web_banner"`
			RoomSilverSeedsBox   int `json:"room-silver_seeds-box"`
			RoomWishingBottle    int `json:"room-wishing_bottle"`
			RoomBoard            int `json:"room-board"`
			RoomSupplication     int `json:"room-supplication"`
			RoomHourRank         int `json:"room-hour_rank"`
			RoomWeekRank         int `json:"room-week_rank"`
			RoomAnchorRank       int `json:"room-anchor_rank"`
			RoomInfoIntegral     int `json:"room-info-integral"`
			RoomSuperChat        int `json:"room-super-chat"`
			RoomTab              int `json:"room-tab"`
			RoomHotRank          int `json:"room-hot-rank"`
			FansMedalProgress    int `json:"fans-medal-progress"`
			GiftBayScreen        int `json:"gift-bay-screen"`
			RoomEnter            int `json:"room-enter"`
			RoomMyIdol           int `json:"room-my-idol"`
			RoomTopic            int `json:"room-topic"`
			FansClub             int `json:"fans-club"`
		} `json:"new_switch_info"`
		SuperChatInfo struct {
			Status      int           `json:"status"`
			JumpUrl     string        `json:"jump_url"`
			Icon        string        `json:"icon"`
			RankedMark  int           `json:"ranked_mark"`
			MessageList []interface{} `json:"message_list"`
		} `json:"super_chat_info"`
		OnlineGoldRankInfoV2 struct {
			List []struct {
				Uid        int    `json:"uid"`
				Face       string `json:"face"`
				Uname      string `json:"uname"`
				Score      string `json:"score"`
				Rank       int    `json:"rank"`
				GuardLevel int    `json:"guard_level"`
			} `json:"list"`
		} `json:"online_gold_rank_info_v2"`
		DmBrushInfo struct {
			MinTime     int `json:"min_time"`
			BrushCount  int `json:"brush_count"`
			SliceCount  int `json:"slice_count"`
			StorageTime int `json:"storage_time"`
		} `json:"dm_brush_info"`
		DmEmoticonInfo struct {
			IsOpenEmoticon   int `json:"is_open_emoticon"`
			IsShieldEmoticon int `json:"is_shield_emoticon"`
		} `json:"dm_emoticon_info"`
		DmTagInfo struct {
			DmTag           int           `json:"dm_tag"`
			Platform        []interface{} `json:"platform"`
			Extra           string        `json:"extra"`
			DmChronosExtra  string        `json:"dm_chronos_extra"`
			DmMode          []interface{} `json:"dm_mode"`
			DmSettingSwitch int           `json:"dm_setting_switch"`
			MaterialConf    interface{}   `json:"material_conf"`
		} `json:"dm_tag_info"`
		TopicInfo struct {
			TopicId   int    `json:"topic_id"`
			TopicName string `json:"topic_name"`
		} `json:"topic_info"`
		GameInfo struct {
			GameStatus int `json:"game_status"`
		} `json:"game_info"`
		WatchedShow struct {
			Switch       bool   `json:"switch"`
			Num          int    `json:"num"`
			TextSmall    string `json:"text_small"`
			TextLarge    string `json:"text_large"`
			Icon         string `json:"icon"`
			IconLocation int    `json:"icon_location"`
			IconWeb      string `json:"icon_web"`
		} `json:"watched_show"`
		TopicRoomInfo struct {
			InteractiveH5Url string `json:"interactive_h5_url"`
			Watermark        int    `json:"watermark"`
		} `json:"topic_room_info"`
		ShowReserveStatus bool `json:"show_reserve_status"`
		SecondCreateInfo  struct {
			ClickPermission int `json:"click_permission"`
		} `json:"second_create_info"`
		PlayTogetherInfo struct {
			Switch   int `json:"switch"`
			IconList []struct {
				Icon    string `json:"icon"`
				Title   string `json:"title"`
				JumpUrl string `json:"jump_url"`
				Status  int    `json:"status"`
			} `json:"icon_list"`
		} `json:"play_together_info"`
		VideoConnectionInfo interface{} `json:"video_connection_info"`
		PlayerThrottleInfo  struct {
			Status              int `json:"status"`
			NormalSleepTime     int `json:"normal_sleep_time"`
			FullscreenSleepTime int `json:"fullscreen_sleep_time"`
			TabSleepTime        int `json:"tab_sleep_time"`
			PromptTime          int `json:"prompt_time"`
		} `json:"player_throttle_info"`
		GuardInfo struct {
			Count                   int `json:"count"`
			AnchorGuardAchieveLevel int `json:"anchor_guard_achieve_level"`
		} `json:"guard_info"`
		HotRankInfo struct {
			Rank      int    `json:"rank"`
			Trend     int    `json:"trend"`
			Countdown int    `json:"countdown"`
			Timestamp int    `json:"timestamp"`
			Url       string `json:"url"`
			Icon      string `json:"icon"`
			AreaName  string `json:"area_name"`
			NewData   struct {
				Rank      int    `json:"rank"`
				Trend     int    `json:"trend"`
				Countdown int    `json:"countdown"`
				Timestamp int    `json:"timestamp"`
				Url       string `json:"url"`
				Icon      string `json:"icon"`
				AreaName  string `json:"area_name"`
				RankDesc  string `json:"rank_desc"`
			} `json:"new_data"`
		} `json:"hot_rank_info"`
	} `json:"data"`
}
