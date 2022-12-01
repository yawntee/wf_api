package asset

import (
	"wf_api/server/util"
)

type StoryEvent = Event

type StoryEventTable map[int]StoryEvent

func (a *Asset) GetStoryEventTable() StoryEventTable {
	if cache, ok := a.Cache["StoryEvent"].(StoryEventTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/event/story_event")
	intMap := parseIntMap(reader)
	table := make(StoryEventTable)
	for id, strings := range intMap {
		startTime := util.ParseIso(strings[13])
		playableEndTime := util.ParseIso(strings[13])
		exchangeableEndTime := util.ParseIso(strings[13])
		table[id] = StoryEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                2,
		}
	}
	a.Cache["StoryEvent"] = table
	return table
}
