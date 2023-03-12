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
	for id, params := range intMap {
		startTime := util.ParseIso(params[14])
		playableEndTime := util.ParseIso(params[15])
		exchangeableEndTime := util.ParseIso(params[16])
		table[id] = StoryEvent{
			Name:                params[2],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
			Type:                EventTypeStory,
		}
	}
	a.Cache["StoryEvent"] = table
	return table
}
