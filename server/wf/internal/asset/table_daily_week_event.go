package asset

import (
	"wf_api/server/util"
)

type DailyWeekEvent = Event

type DailyWeekEventTable = map[int]DailyWeekEvent

func (a *Asset) GetDailyWeekEventTable() DailyWeekEventTable {
	if cache, ok := a.Cache["DailyWeekEvent"].(DailyWeekEventTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/event/daily_week_event")
	intMap := parseIntMap(reader)
	table := make(DailyWeekEventTable)
	for id, strings := range intMap {
		startTime := util.ParseIso(strings[13])
		playableEndTime := util.ParseIso(strings[13])
		exchangeableEndTime := util.ParseIso(strings[13])
		table[id] = DailyWeekEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
		}
	}
	a.Cache["DailyWeekEvent"] = table
	return table
}
