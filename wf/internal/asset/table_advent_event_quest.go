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

type AdventEventQuest struct {
	Name      string
	StartTime time.Time
	EndTime   *time.Time
}

type AdventEventQuestTable map[int][]AdventEventQuest

func (a *Asset) GetAdventEventQuestTable() AdventEventQuestTable {
	if cache, ok := a.Cache["AdventEventQuest"].(AdventEventQuestTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/event/advent_event_quest")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(AdventEventQuestTable)
	for id, chapter := range intMap {
		var quests []AdventEventQuest
		for _, params := range chapter {
			startTime := util.ParseIso(params[5])
			var endTime *time.Time
			if endTimeStr := params[6]; endTimeStr != "(None)" {
				iso := util.ParseIso(endTimeStr)
				endTime = &iso
			}
			quests = append(quests, AdventEventQuest{
				Name:      params[2],
				StartTime: startTime,
				EndTime:   endTime,
			})
		}
		table[id] = quests
	}
	a.Cache["AdventEventQuest"] = table
	return table
}
