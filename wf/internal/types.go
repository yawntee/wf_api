package internal

import "errors"

var (
	ErrNoLogin = errors.New("您还未登录")
)

type GameUser struct {
	Uid      int64  `json:"Uid,omitempty"`
	Sid      string `json:"Sid,omitempty"`
	Username string `json:"Username,omitempty"`
	Token    string `json:"Token,omitempty"`
}

type GameResp[T any] struct {
	DataHeaders struct {
		DeviceId    int        `msgpack:"device_id"`
		ForceUpdate bool       `msgpack:"force_update"`
		ResultCode  ResultCode `msgpack:"result_code"`
		Servertime  int        `msgpack:"servertime"`
		ShortUdid   int        `msgpack:"short_udid"`
		ViewerId    int        `msgpack:"viewer_id"`
	} `msgpack:"data_headers"`
	Data T `msgpack:"data"`
}

type GameUserInfo struct {
	UserInfo struct {
		Stamina               int           `msgpack:"stamina"`
		StaminaHealTime       int           `msgpack:"stamina_heal_time"`
		BoostPoint            int           `msgpack:"boost_point"`
		BossBoostPoint        int           `msgpack:"boss_boost_point"`
		TransitionState       int           `msgpack:"transition_state"`
		Role                  int           `msgpack:"role"`
		Name                  string        `msgpack:"name"`
		LastLoginTime         string        `msgpack:"last_login_time"`
		Comment               string        `msgpack:"comment"`
		Vmoney                int           `msgpack:"vmoney"`
		FreeVmoney            int           `msgpack:"free_vmoney"`
		RankPoint             int           `msgpack:"rank_point"`
		StarCrumb             int           `msgpack:"star_crumb"`
		BondToken             int           `msgpack:"bond_token"`
		ExpPool               int           `msgpack:"exp_pool"`
		ExpPooledTime         int           `msgpack:"exp_pooled_time"`
		LeaderCharacterId     int           `msgpack:"leader_character_id"`
		PartySlot             int           `msgpack:"party_slot"`
		DegreeId              int           `msgpack:"degree_id"`
		Birth                 int           `msgpack:"birth"`
		FreeMana              int           `msgpack:"free_mana"`
		PaidMana              int           `msgpack:"paid_mana"`
		MonthCardRemainDays   int           `msgpack:"month_card_remain_days"`
		WeeklyBonusRemainDays int           `msgpack:"weekly_bonus_remain_days"`
		IsBoughtFundLaite     bool          `msgpack:"is_bought_fund_laite"`
		IsBoughtFundMainQuest bool          `msgpack:"is_bought_fund_main_quest"`
		IsBoughtFundExQuest   bool          `msgpack:"is_bought_fund_ex_quest"`
		MonthlyPaymentTotal   int           `msgpack:"monthly_payment_total"`
		RenewalGiftRemainDays []interface{} `msgpack:"renewal_gift_remain_days"`
	} `msgpack:"user_info"`
	PaymentRebateInfo struct {
		Status      int `msgpack:"status"`
		ExpiredTime int `msgpack:"expired_time"`
	} `msgpack:"payment_rebate_info"`
	UserDailyChallengePointList []struct {
		Id           int           `msgpack:"id"`
		Point        int           `msgpack:"point"`
		CampaignList []interface{} `msgpack:"campaign_list"`
	} `msgpack:"user_daily_challenge_point_list"`
	BonusIndexList []struct {
		BonusGroupId   string `msgpack:"bonus_group_id"`
		BonusGroupType string `msgpack:"bonus_group_type"`
		Index          int    `msgpack:"index"`
	} `msgpack:"bonus_index_list"`
	LoginBonusReceivedAt int `msgpack:"login_bonus_received_at"`
	UserNoticeList       []struct {
		Kind   int `msgpack:"kind"`
		Status int `msgpack:"status"`
	} `msgpack:"user_notice_list"`
	UserTriggeredTutorial []int `msgpack:"user_triggered_tutorial"`
	UserCharacterList     []struct {
		ViewerId          int         `msgpack:"viewer_id"`
		CharacterId       int         `msgpack:"character_id"`
		EntryCount        int         `msgpack:"entry_count"`
		ActionSkill1      interface{} `msgpack:"action_skill_1"`
		ActionSkill2      interface{} `msgpack:"action_skill_2"`
		EpisodeLearnCount int         `msgpack:"episode_learn_count"`
		EvolutionLevel    int         `msgpack:"evolution_level"`
		EvolutionImgLevel int         `msgpack:"evolution_img_level"`
		OverLimitStep     int         `msgpack:"over_limit_step"`
		Protection        bool        `msgpack:"protection"`
		CreateTime        string      `msgpack:"create_time"`
		UpdateTime        string      `msgpack:"update_time"`
		DeleteTime        interface{} `msgpack:"delete_time"`
		Exp               int         `msgpack:"exp"`
		ExpTotal          int         `msgpack:"exp_total"`
		Stack             int         `msgpack:"stack"`
		BondTokenList     []struct {
			ManaBoardIndex int `msgpack:"mana_board_index"`
			Status         int `msgpack:"status"`
		} `msgpack:"bond_token_list"`
		ManaBoardIndex int `msgpack:"mana_board_index"`
	} `msgpack:"user_character_list"`
	UserCharacterManaNodeList []interface{} `msgpack:"user_character_mana_node_list"`
	UserPartyGroupList        []struct {
		Null              int `msgpack:"null"`
		PartyGroupId      int `msgpack:"party_group_id"`
		PartyGroupColorId int `msgpack:"party_group_color_id"`
		PartyList         []struct {
			PartyGroupId       int    `msgpack:"party_group_id"`
			PartyId            int    `msgpack:"party_id"`
			PartyName          string `msgpack:"party_name"`
			CharacterIds       []int  `msgpack:"character_ids"`
			UnisonCharacterIds []int  `msgpack:"unison_character_ids"`
			EquipmentIds       []int  `msgpack:"equipment_ids"`
			PartyEdited        bool   `msgpack:"party_edited"`
			PartyType          int    `msgpack:"party_type"`
		} `msgpack:"party_list"`
	} `msgpack:"user_party_group_list"`
	ItemList          map[int]int `msgpack:"item_list"`
	UserEquipmentList []struct {
		Null        int  `msgpack:"null"`
		ViewerId    int  `msgpack:"viewer_id"`
		EquipmentId int  `msgpack:"equipment_id"`
		Protection  bool `msgpack:"protection"`
		Level       int  `msgpack:"level"`
		Stack       int  `msgpack:"stack"`
	} `msgpack:"user_equipment_list"`
	UserCharacterFromTownHistory []interface{} `msgpack:"user_character_from_town_history"`
	QuestProgress                map[int][]struct {
		QuestId   int  `msgpack:"quest_id"`
		Finished  bool `msgpack:"finished"`
		HighScore int  `msgpack:"high_score,omitempty"`
		ClearRank int  `msgpack:"clear_rank,omitempty"`
	} `msgpack:"quest_progress"`
	LastMainQuestId int `msgpack:"last_main_quest_id"`
	GachaInfoList   []struct {
		GachaId            int  `msgpack:"gacha_id"`
		IsDailyFirst       bool `msgpack:"is_daily_first"`
		IsAccountFirst     bool `msgpack:"is_account_first"`
		DailyOneCount      int  `msgpack:"daily_one_count"`
		DailyTenCount      int  `msgpack:"daily_ten_count"`
		CrazyDrawCount     int  `msgpack:"crazy_draw_count"`
		GachaExchangePoint int  `msgpack:"gacha_exchange_point,omitempty"`
	} `msgpack:"gacha_info_list"`
	AvailableAssetVersion            string          `msgpack:"available_asset_version"`
	ShouldPromptTakeoverRegistration bool            `msgpack:"should_prompt_takeover_registration"`
	HasUnreadNewsItem                bool            `msgpack:"has_unread_news_item"`
	FundReceiveList                  []interface{}   `msgpack:"fund_receive_list"`
	CrazyGachaResultList             [][]interface{} `msgpack:"crazy_gacha_result_list"`
	LastCrazyGachaDrawResult         []interface{}   `msgpack:"last_crazy_gacha_draw_result"`
	MonthlyChargeBonusInfo           struct {
		InitTime    int `msgpack:"init_time"`
		BonusDays   int `msgpack:"bonus_days"`
		ExpiredTime int `msgpack:"expired_time"`
	} `msgpack:"monthly_charge_bonus_info"`
	SimplePaymentItemList []struct {
		StoreProductId string  `msgpack:"store_product_id"`
		StartTime      float64 `msgpack:"start_time"`
		EndTime        float64 `msgpack:"end_time"`
	} `msgpack:"simple_payment_item_list"`
	MonthlyTip bool `msgpack:"monthly_tip"`
	UserOption struct {
		GachaPlayNoRarityUpMovie bool `msgpack:"gacha_play_no_rarity_up_movie"`
		AutoPlay                 bool `msgpack:"auto_play"`
		NumberNotationSymbol     bool `msgpack:"number_notation_symbol"`
		PaymentAlert             bool `msgpack:"payment_alert"`
		RoomNumberHidden         bool `msgpack:"room_number_hidden"`
		AttentionSoundEffect     bool `msgpack:"attention_sound_effect"`
		AttentionVibration       bool `msgpack:"attention_vibration"`
		AttentionEnableInBattle  bool `msgpack:"attention_enable_in_battle"`
		SimpleAbilityDescription bool `msgpack:"simple_ability_description"`
		Stamina                  bool `msgpack:"stamina"`
		ServerPush               bool `msgpack:"server_push"`
	} `msgpack:"user_option"`
	DrawnQuestList []struct {
		CategoryId int `msgpack:"category_id"`
		QuestId    int `msgpack:"quest_id"`
		OddsId     int `msgpack:"odds_id"`
	} `msgpack:"drawn_quest_list"`
	MailArrived               bool `msgpack:"mail_arrived"`
	MissionTips               bool `msgpack:"mission_tips"`
	ClearedRegularMissionList []struct {
		MissionId int `msgpack:"mission_id"`
		Stage     int `msgpack:"stage"`
	} `msgpack:"cleared_regular_mission_list"`
	AllActiveMissionList []struct {
		MissionId     int `msgpack:"mission_id"`
		ProgressValue int `msgpack:"progress_value"`
		Stages        []struct {
			Stage    int  `msgpack:"stage"`
			Received bool `msgpack:"received"`
		} `msgpack:"stages"`
	} `msgpack:"all_active_mission_list"`
	StartDashExchangeCampaignList []interface{} `msgpack:"start_dash_exchange_campaign_list"`
	Config                        struct {
		AttentionRecruitmentIntervalSeconds      int     `msgpack:"attention_recruitment_interval_seconds"`
		AttentionRecruitmentRedeliverLimit       int     `msgpack:"attention_recruitment_redeliver_limit"`
		AttentionPollingIntervalSecondsNormal    int     `msgpack:"attention_polling_interval_seconds_normal"`
		AttentionPollingIntervalSecondsBattle    int     `msgpack:"attention_polling_interval_seconds_battle"`
		MultiAttentionLifetimeSeconds            int     `msgpack:"multi_attention_lifetime_seconds"`
		ContributionScoreRateToParasite          float64 `msgpack:"contribution_score_rate_to_parasite"`
		AttentionLogIntervalSeconds              int     `msgpack:"attention_log_interval_seconds"`
		DisableFinishDurationSeconds             int     `msgpack:"disable_finish_duration_seconds"`
		DisableDeclineCountSeconds               int     `msgpack:"disable_decline_count_seconds"`
		DisableDeclineCountLimit                 int     `msgpack:"disable_decline_count_limit"`
		DisableDeclineDurationSeconds            int     `msgpack:"disable_decline_duration_seconds"`
		DisableIntentDisconnectDurationSeconds   int     `msgpack:"disable_intent_disconnect_duration_seconds"`
		DisableUnintentDisconnectDurationSeconds int     `msgpack:"disable_unintent_disconnect_duration_seconds"`
		DisableRemoteErrorDurationSeconds        int     `msgpack:"disable_remote_error_duration_seconds"`
		AttentionAnimationTimeSeconds            int     `msgpack:"attention_animation_time_seconds"`
		DisableExpireCountLimit                  int     `msgpack:"disable_expire_count_limit"`
		DisableExpireDurationSeconds             int     `msgpack:"disable_expire_duration_seconds"`
		PollingDelayNormalSecondsRangeMin        int     `msgpack:"polling_delay_normal_seconds_range_min"`
		PollingDelayNormalSecondsRangeMax        int     `msgpack:"polling_delay_normal_seconds_range_max"`
		PollingDelayBattleSecondsRangeMin        int     `msgpack:"polling_delay_battle_seconds_range_min"`
		PollingDelayBattleSecondsRangeMax        int     `msgpack:"polling_delay_battle_seconds_range_max"`
	} `msgpack:"config"`
	LoginInfo struct {
		Sign          string `msgpack:"sign"`
		CreateDate    string `msgpack:"createDate"`
		RoleName      string `msgpack:"roleName"`
		RoleId        int    `msgpack:"roleId"`
		ServerName    string `msgpack:"serverName"`
		ServerId      string `msgpack:"serverId"`
		TimeUsed      int    `msgpack:"timeUsed"`
		AccountName   string `msgpack:"accountName"`
		LoginMode     int    `msgpack:"loginMode"`
		LoginType     int    `msgpack:"loginType"`
		NewAccount    int    `msgpack:"newAccount"`
		CreditAccount int    `msgpack:"creditAccount"`
		PhysicalValue int    `msgpack:"physicalValue"`
		RoleLevel     int    `msgpack:"roleLevel"`
		Ip            string `msgpack:"ip"`
	} `msgpack:"login_info"`
	CnCrashUrl            string `msgpack:"cn_crash_url"`
	SurveyUrl             string `msgpack:"survey_url"`
	QqGroupUrl            string `msgpack:"qq_group_url"`
	BugReportUrl          string `msgpack:"bug_report_url"`
	EnableGift            bool   `msgpack:"enable_gift"`
	EnableCustomerService bool   `msgpack:"enable_customer_service"`
	EnableRename          bool   `msgpack:"enable_rename"`
	ExecuteInBackground   bool   `msgpack:"execute_in_background"`
	EnableDeleteFile      bool   `msgpack:"enable_delete_file"`
}
