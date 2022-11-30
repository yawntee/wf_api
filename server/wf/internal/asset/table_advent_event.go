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
	for i, strings := range intMap {
		startTime := util.ParseIso(strings[15])
		playableEndTime := util.ParseIso(strings[16])
		exchangeableEndTime := util.ParseIso(strings[17])
		table[i] = AdventEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                0,
		}
	}
	a.Cache["AdventEvent"] = table
	return table
}
