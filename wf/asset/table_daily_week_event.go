package asset

import (
	"wf_api/util"
)

type DailyWeekEvent = Event

type DailyWeekEventTable map[int]DailyWeekEvent

func (a *CacheAsset) GetDailyWeekEventTable() DailyWeekEventTable {
	if cache, ok := a.Cache["DailyWeekEvent"].(DailyWeekEventTable); ok {
		return cache
	}
	reader := a.getTableFile("/quest/event/daily_week_event")
	intMap := parseIntMap(reader)
	table := make(DailyWeekEventTable)
	for i, strings := range intMap {
		startTime := util.ParseIso(strings[13])
		playableEndTime := util.ParseIso(strings[13])
		exchangeableEndTime := util.ParseIso(strings[13])
		table[i] = DailyWeekEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
		}
	}
	a.Cache["DailyWeekEvent"] = table
	return table
}
