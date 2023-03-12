package internal

import (
	"errors"
)

var (
	ErrNoLogin = errors.New("您还未登录")
)

type GameUser struct {
	Uid      uint64 `json:"uid,omitempty"`
	Sid      string `json:"sid,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}

type GameUpdateData struct {
	Info struct {
		ClientAssetVersion         bool   `mapstructure:"client_asset_version"`
		TargetAssetVersion         string `mapstructure:"target_asset_version"`
		EventualTargetAssetVersion string `mapstructure:"eventual_target_asset_version"`
		IsInitial                  bool   `mapstructure:"is_initial"`
		LatestMajFirstVersion      string `mapstructure:"latest_maj_first_version"`
	} `mapstructure:"info"`
	Full struct {
		Version string `mapstructure:"version"`
		Archive []struct {
			Location string `mapstructure:"location"`
			Size     int    `mapstructure:"size"`
			Sha256   string `mapstructure:"sha256"`
		} `mapstructure:"archive"`
	} `mapstructure:"full"`
	Diff []struct {
		Version         string `mapstructure:"version"`
		OriginalVersion string `mapstructure:"original_version"`
		Archive         []struct {
			Location string `mapstructure:"location"`
			Size     int    `mapstructure:"size"`
			Sha256   string `mapstructure:"sha256"`
		} `mapstructure:"archive"`
	} `mapstructure:"diff"`
	AssetVersionHash string `mapstructure:"asset_version_hash"`
}

type GameUserInfo struct {
	UserInfo struct {
		Stamina               int           `mapstructure:"stamina"`
		StaminaHealTime       int           `mapstructure:"stamina_heal_time"`
		BoostPoint            int           `mapstructure:"boost_point"`
		BossBoostPoint        int           `mapstructure:"boss_boost_point"`
		TransitionState       int           `mapstructure:"transition_state"`
		Role                  int           `mapstructure:"role"`
		Name                  string        `mapstructure:"name"`
		LastLoginTime         string        `mapstructure:"last_login_time"`
		Comment               string        `mapstructure:"comment"`
		Vmoney                int           `mapstructure:"vmoney"`
		FreeVmoney            int           `mapstructure:"free_vmoney"`
		RankPoint             int           `mapstructure:"rank_point"`
		StarCrumb             int           `mapstructure:"star_crumb"`
		BondToken             int           `mapstructure:"bond_token"`
		ExpPool               int           `mapstructure:"exp_pool"`
		ExpPooledTime         int           `mapstructure:"exp_pooled_time"`
		LeaderCharacterId     int           `mapstructure:"leader_character_id"`
		PartySlot             int           `mapstructure:"party_slot"`
		DegreeId              int           `mapstructure:"degree_id"`
		Birth                 int           `mapstructure:"birth"`
		FreeMana              int           `mapstructure:"free_mana"`
		PaidMana              int           `mapstructure:"paid_mana"`
		MonthCardRemainDays   int           `mapstructure:"month_card_remain_days"`
		WeeklyBonusRemainDays int           `mapstructure:"weekly_bonus_remain_days"`
		IsBoughtFundLaite     bool          `mapstructure:"is_bought_fund_laite"`
		IsBoughtFundMainQuest bool          `mapstructure:"is_bought_fund_main_quest"`
		IsBoughtFundExQuest   bool          `mapstructure:"is_bought_fund_ex_quest"`
		MonthlyPaymentTotal   int           `mapstructure:"monthly_payment_total"`
		RenewalGiftRemainDays []interface{} `mapstructure:"renewal_gift_remain_days"`
	} `mapstructure:"user_info"`
	PaymentRebateInfo struct {
		Status      int `mapstructure:"status"`
		ExpiredTime int `mapstructure:"expired_time"`
	} `mapstructure:"payment_rebate_info"`
	UserDailyChallengePointList []struct {
		Id           int `mapstructure:"id"`
		Point        int `mapstructure:"point"`
		CampaignList []struct {
			CampaignId      int `mapstructure:"campaign_id"`
			AdditionalPoint int `mapstructure:"additional_point"`
		} `mapstructure:"campaign_list"`
	} `mapstructure:"user_daily_challenge_point_list"`
	BonusIndexList        []interface{} `mapstructure:"bonus_index_list"`
	UserNoticeList        []interface{} `mapstructure:"user_notice_list"`
	UserTriggeredTutorial []int         `mapstructure:"user_triggered_tutorial"`
	UserCharacterList     []struct {
		ViewerId          int    `mapstructure:"viewer_id"`
		CharacterId       int    `mapstructure:"character_id"`
		EntryCount        int    `mapstructure:"entry_count"`
		EpisodeLearnCount int    `mapstructure:"episode_learn_count"`
		EvolutionLevel    int    `mapstructure:"evolution_level"`
		EvolutionImgLevel int    `mapstructure:"evolution_img_level"`
		OverLimitStep     int    `mapstructure:"over_limit_step"`
		Protection        bool   `mapstructure:"protection"`
		AbilitySoulSlot1  int    `mapstructure:"ability_soul_slot_1,omitempty"`
		CreateTime        string `mapstructure:"create_time"`
		UpdateTime        string `mapstructure:"update_time"`
		Exp               int    `mapstructure:"exp"`
		ExpTotal          int    `mapstructure:"exp_total"`
		Stack             int    `mapstructure:"stack"`
		BondTokenList     []struct {
			ManaBoardIndex int `mapstructure:"mana_board_index"`
			Status         int `mapstructure:"status"`
		} `mapstructure:"bond_token_list"`
		ManaBoardIndex int `mapstructure:"mana_board_index"`
	} `mapstructure:"user_character_list"`
	UserCharacterManaNodeList map[int][]struct {
		ManaNodeMultipliedId int `mapstructure:"mana_node_multiplied_id"`
	} `mapstructure:"user_character_mana_node_list"`
	UserPartyGroupList []struct {
		Null              int `mapstructure:"null"`
		PartyGroupId      int `mapstructure:"party_group_id"`
		PartyGroupColorId int `mapstructure:"party_group_color_id"`
		PartyList         []struct {
			PartyGroupId       int    `mapstructure:"party_group_id"`
			PartyId            int    `mapstructure:"party_id"`
			PartyName          string `mapstructure:"party_name"`
			CharacterIds       []int  `mapstructure:"character_ids"`
			UnisonCharacterIds []int  `mapstructure:"unison_character_ids"`
			EquipmentIds       []int  `mapstructure:"equipment_ids"`
			PartyEdited        bool   `mapstructure:"party_edited"`
			PartyType          int    `mapstructure:"party_type"`
		} `mapstructure:"party_list"`
	} `mapstructure:"user_party_group_list"`
	ItemList          map[int]int `mapstructure:"item_list"`
	UserEquipmentList []struct {
		Null        int  `mapstructure:"null"`
		ViewerId    int  `mapstructure:"viewer_id"`
		EquipmentId int  `mapstructure:"equipment_id"`
		Protection  bool `mapstructure:"protection"`
		Level       int  `mapstructure:"level"`
		Stack       int  `mapstructure:"stack"`
	} `mapstructure:"user_equipment_list"`
	UserCharacterFromTownHistory []struct {
		CharacterId int `mapstructure:"character_id"`
	} `mapstructure:"user_character_from_town_history"`
	QuestProgress map[int][]struct {
		QuestId      int  `mapstructure:"quest_id"`
		Finished     bool `mapstructure:"finished"`
		HighScore    int  `mapstructure:"high_score"`
		ClearRank    int  `mapstructure:"clear_rank"`
		RankingEvent struct {
			BestRecord struct {
				ElapsedTimeMs  int  `mapstructure:"elapsed_time_ms"`
				Score          int  `mapstructure:"score"`
				IsAccomplished bool `mapstructure:"is_accomplished"`
			} `mapstructure:"best_record"`
		} `mapstructure:"ranking_event"`
	} `mapstructure:"quest_progress"`
	LastMainQuestId int `mapstructure:"last_main_quest_id"`
	GachaInfoList   []struct {
		GachaId            int  `mapstructure:"gacha_id"`
		IsDailyFirst       bool `mapstructure:"is_daily_first"`
		IsAccountFirst     bool `mapstructure:"is_account_first"`
		DailyOneCount      int  `mapstructure:"daily_one_count"`
		DailyTenCount      int  `mapstructure:"daily_ten_count"`
		CrazyDrawCount     int  `mapstructure:"crazy_draw_count"`
		GachaExchangePoint int  `mapstructure:"gacha_exchange_point,omitempty"`
	} `mapstructure:"gacha_info_list"`
	AvailableAssetVersion            string          `mapstructure:"available_asset_version"`
	ShouldPromptTakeoverRegistration bool            `mapstructure:"should_prompt_takeover_registration"`
	HasUnreadNewsItem                bool            `mapstructure:"has_unread_news_item"`
	FundReceiveList                  []int           `mapstructure:"fund_receive_list"`
	CrazyGachaResultList             [][]interface{} `mapstructure:"crazy_gacha_result_list"`
	LastCrazyGachaDrawResult         []interface{}   `mapstructure:"last_crazy_gacha_draw_result"`
	MonthlyChargeBonusInfo           struct {
		InitTime    int `mapstructure:"init_time"`
		BonusDays   int `mapstructure:"bonus_days"`
		ExpiredTime int `mapstructure:"expired_time"`
	} `mapstructure:"monthly_charge_bonus_info"`
	SimplePaymentItemList []struct {
		StoreProductId string  `mapstructure:"store_product_id"`
		StartTime      float64 `mapstructure:"start_time"`
		EndTime        float64 `mapstructure:"end_time"`
	} `mapstructure:"simple_payment_item_list"`
	MonthlyTip bool `mapstructure:"monthly_tip"`
	UserOption struct {
		GachaPlayNoRarityUpMovie bool `mapstructure:"gacha_play_no_rarity_up_movie"`
		AutoPlay                 bool `mapstructure:"auto_play"`
		NumberNotationSymbol     bool `mapstructure:"number_notation_symbol"`
		PaymentAlert             bool `mapstructure:"payment_alert"`
		RoomNumberHidden         bool `mapstructure:"room_number_hidden"`
		AttentionSoundEffect     bool `mapstructure:"attention_sound_effect"`
		AttentionVibration       bool `mapstructure:"attention_vibration"`
		AttentionEnableInBattle  bool `mapstructure:"attention_enable_in_battle"`
		SimpleAbilityDescription bool `mapstructure:"simple_ability_description"`
		Stamina                  bool `mapstructure:"stamina"`
		ServerPush               bool `mapstructure:"server_push"`
	} `mapstructure:"user_option"`
	DrawnQuestList []struct {
		CategoryId int `mapstructure:"category_id"`
		QuestId    int `mapstructure:"quest_id"`
		OddsId     int `mapstructure:"odds_id"`
	} `mapstructure:"drawn_quest_list"`
	MailArrived               bool `mapstructure:"mail_arrived"`
	MissionTips               bool `mapstructure:"mission_tips"`
	ClearedRegularMissionList []struct {
		MissionId int `mapstructure:"mission_id"`
		Stage     int `mapstructure:"stage"`
	} `mapstructure:"cleared_regular_mission_list"`
	AllActiveMissionList []struct {
		MissionId     int `mapstructure:"mission_id"`
		ProgressValue int `mapstructure:"progress_value"`
		Stages        []struct {
			Stage    int  `mapstructure:"stage"`
			Received bool `mapstructure:"received"`
		} `mapstructure:"stages"`
	} `mapstructure:"all_active_mission_list"`
	GachaCampaignList []struct {
		CampaignId int `mapstructure:"campaign_id"`
		GachaId    int `mapstructure:"gacha_id"`
		Count      int `mapstructure:"count"`
	} `mapstructure:"gacha_campaign_list"`
	PurchasedTimesList          map[string]int `mapstructure:"purchased_times_list"`
	SpecialExchangeCampaignList []struct {
		Null       int `mapstructure:"null"`
		CampaignId int `mapstructure:"campaign_id"`
		Status     int `mapstructure:"status"`
	} `mapstructure:"special_exchange_campaign_list"`
	StartDashExchangeCampaignList []interface{} `mapstructure:"start_dash_exchange_campaign_list"`
	Config                        struct {
		AttentionRecruitmentIntervalSeconds      int     `mapstructure:"attention_recruitment_interval_seconds"`
		AttentionRecruitmentRedeliverLimit       int     `mapstructure:"attention_recruitment_redeliver_limit"`
		AttentionPollingIntervalSecondsNormal    int     `mapstructure:"attention_polling_interval_seconds_normal"`
		AttentionPollingIntervalSecondsBattle    int     `mapstructure:"attention_polling_interval_seconds_battle"`
		MultiAttentionLifetimeSeconds            int     `mapstructure:"multi_attention_lifetime_seconds"`
		ContributionScoreRateToParasite          float64 `mapstructure:"contribution_score_rate_to_parasite"`
		AttentionLogIntervalSeconds              int     `mapstructure:"attention_log_interval_seconds"`
		DisableFinishDurationSeconds             int     `mapstructure:"disable_finish_duration_seconds"`
		DisableDeclineCountSeconds               int     `mapstructure:"disable_decline_count_seconds"`
		DisableDeclineCountLimit                 int     `mapstructure:"disable_decline_count_limit"`
		DisableDeclineDurationSeconds            int     `mapstructure:"disable_decline_duration_seconds"`
		DisableIntentDisconnectDurationSeconds   int     `mapstructure:"disable_intent_disconnect_duration_seconds"`
		DisableUnintentDisconnectDurationSeconds int     `mapstructure:"disable_unintent_disconnect_duration_seconds"`
		DisableRemoteErrorDurationSeconds        int     `mapstructure:"disable_remote_error_duration_seconds"`
		AttentionAnimationTimeSeconds            int     `mapstructure:"attention_animation_time_seconds"`
		DisableExpireCountLimit                  int     `mapstructure:"disable_expire_count_limit"`
		DisableExpireDurationSeconds             int     `mapstructure:"disable_expire_duration_seconds"`
		PollingDelayNormalSecondsRangeMin        int     `mapstructure:"polling_delay_normal_seconds_range_min"`
		PollingDelayNormalSecondsRangeMax        int     `mapstructure:"polling_delay_normal_seconds_range_max"`
		PollingDelayBattleSecondsRangeMin        int     `mapstructure:"polling_delay_battle_seconds_range_min"`
		PollingDelayBattleSecondsRangeMax        int     `mapstructure:"polling_delay_battle_seconds_range_max"`
	} `mapstructure:"config"`
	LoginInfo struct {
		Sign          string `mapstructure:"sign"`
		CreateDate    string `mapstructure:"createDate"`
		RoleName      string `mapstructure:"roleName"`
		RoleId        int    `mapstructure:"roleId"`
		ServerName    string `mapstructure:"serverName"`
		ServerId      string `mapstructure:"serverId"`
		TimeUsed      int    `mapstructure:"timeUsed"`
		AccountName   string `mapstructure:"accountName"`
		LoginMode     int    `mapstructure:"loginMode"`
		LoginType     int    `mapstructure:"loginType"`
		NewAccount    int    `mapstructure:"newAccount"`
		CreditAccount int    `mapstructure:"creditAccount"`
		PhysicalValue int    `mapstructure:"physicalValue"`
		RoleLevel     int    `mapstructure:"roleLevel"`
		Ip            string `mapstructure:"ip"`
	} `mapstructure:"login_info"`
	CnCrashUrl            string `mapstructure:"cn_crash_url"`
	SurveyUrl             string `mapstructure:"survey_url"`
	QqGroupUrl            string `mapstructure:"qq_group_url"`
	BugReportUrl          string `mapstructure:"bug_report_url"`
	EnableGift            bool   `mapstructure:"enable_gift"`
	EnableCustomerService bool   `mapstructure:"enable_customer_service"`
	EnableRename          bool   `mapstructure:"enable_rename"`
	ExecuteInBackground   bool   `mapstructure:"execute_in_background"`
	EnableDeleteFile      bool   `mapstructure:"enable_delete_file"`
}

type ItemClaimedData struct {
	ItemList      map[int]int `json:"item_list"`
	EquipmentList []struct {
		Null        int  `json:"null"`
		ViewerId    int  `json:"viewer_id"`
		EquipmentId int  `json:"equipment_id"`
		Protection  bool `json:"protection"`
		Level       int  `json:"level"`
		Stack       int  `json:"stack"`
	} `json:"equipment_list"`
	MailArrived bool `json:"mail_arrived"`
}
