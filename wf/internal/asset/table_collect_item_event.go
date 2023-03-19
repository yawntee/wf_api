package asset

import (
	"time"
	"wf_api/util"
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
	for id, params := range intMap {
		startTime := util.ParseIso(params[19])
		playableEndTime := util.ParseIso(params[20])
		var exchangeableEndTime time.Time
		if t := params[21]; t != "(None)" {
			exchangeableEndTime = util.ParseIso(t)
		} else {
			exchangeableEndTime = playableEndTime
		}
		table[id] = CollectItemEvent{
			Name:                params[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                EventTypeCarnivalEvent,
		}
	}
	a.Cache["CollectItemEvent"] = table
	return table
}
