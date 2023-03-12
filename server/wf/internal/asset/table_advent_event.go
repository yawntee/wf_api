package asset

import (
	"wf_api/server/util"
)

type AdventEvent = Event

type AdventEventTable = map[int]AdventEvent

func (a *Asset) GetAdventEventTable() AdventEventTable {
	if cache, ok := a.Cache["AdventEvent"].(AdventEventTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/event/advent_event")
	intMap := parseIntMap(reader)
	table := make(AdventEventTable)
	for i, params := range intMap {
		startTime := util.ParseIso(params[16])
		playableEndTime := util.ParseIso(params[17])
		exchangeableEndTime := util.ParseIso(params[18])
		table[i] = AdventEvent{
			Name:                params[2],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                EventTypeAdvent,
		}
	}
	a.Cache["AdventEvent"] = table
	return table
}
