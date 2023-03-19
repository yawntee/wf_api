package asset

import (
	"bytes"
	"time"
	"wf_api/util"
)

/*
{
	1:{
		id:{
			chapter:{}
			}
	}
}
*/

type StoryEventQuest struct {
	Name      string
	StartTime time.Time
	EndTime   *time.Time
}

type StoryEventQuestTable map[int][]StoryEventQuest

func (a *Asset) GetStoryEventQuestTable() StoryEventQuestTable {
	if cache, ok := a.Cache["StoryEventQuest"].(StoryEventQuestTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/event/story_event_single_quest")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(StoryEventQuestTable)
	for id, chapter := range intMap {
		var quests []StoryEventQuest
		for _, params := range chapter {
			startTime := util.ParseIso(params[5])
			var endTime *time.Time
			if endTimeStr := params[6]; endTimeStr != "(None)" {
				iso := util.ParseIso(endTimeStr)
				endTime = &iso
			}
			quests = append(quests, StoryEventQuest{
				Name:      params[2],
				StartTime: startTime,
				EndTime:   endTime,
			})
		}
		table[id] = quests
	}
	a.Cache["StoryEventQuest"] = table
	return table
}
