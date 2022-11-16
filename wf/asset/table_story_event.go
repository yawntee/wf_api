package asset

import (
	"fmt"
	"wf_api/util"
)

type StoryEvent = Event

type StoryEventTable map[int]StoryEvent

func (a *Asset) GetStoryEventTable() StoryEventTable {
	if cache, ok := a.Cache["StoryEvent"].(StoryEventTable); ok {
		return cache
	}
	reader := a.getTableFile("/quest/event/story_event")
	intMap := parseIntMap(reader)
	fmt.Println(intMap)
	table := make(StoryEventTable)
	for i, strings := range intMap {
		startTime := util.ParseIso(strings[13])
		playableEndTime := util.ParseIso(strings[13])
		exchangeableEndTime := util.ParseIso(strings[13])
		table[i] = StoryEvent{
			Name:                strings[1],
			StartTime:           startTime,
			PlayableEndTime:     playableEndTime,
			ExchangeableEndTime: exchangeableEndTime,
		}
	}
	a.Cache["StoryEvent"] = table
	return table
}
