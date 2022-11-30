package asset

import (
	"time"
	"wf_api/server/util"
)

type CollectItemEvent = Event

type CollectItemEventTable = map[int]CollectItemEvent

func (a *Asset) GetCollectItemEventTable() CollectItemEventTable {
	if cache, ok := a.Cache["CollectItemEvent"].(CollectItemEventTable); ok {
		return cache
	}
	reader := a.GetTableFile("/reward/event/collect_item_event")
	intMap := parseIntMap(reader)
	table := make(CollectItemEventTable)
	for id, strings := range intMap {
		startTime := util.ParseIso(strings[17])
		playableEndTime := util.ParseIso(strings[18])
		var exchangeableEndTime time.Time
		if t := strings[19]; t != "(None)" {
			exchangeableEndTime = util.ParseIso(strings[19])
		} else {
			exchangeableEndTime = playableEndTime
		}
		table[id] = CollectItemEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                9,
		}
	}
	a.Cache["CollectItemEvent"] = table
	return table
}
