package asset

import (
	"fmt"
	"strconv"
)

type EventList struct {
	Id   int
	Type EventType
}

type EventListTable map[int]EventList

func (a *CacheAsset) GetEventListTable() EventListTable {
	if cache, ok := a.Cache["EventList"].(EventListTable); ok {
		return cache
	}
	reader := a.getTableFile("/quest/event/event_list")
	intMap := parseIntMap(reader)
	table := make(EventListTable)
	for i, strings := range intMap {
		id, err := strconv.Atoi(strings[1])
		if err != nil {
			panic(err)
		}
		_type, err := strconv.Atoi(strings[0])
		if err != nil {
			panic(err)
		}
		if EventType(_type) > EventTypeExpertSingle {
			panic(fmt.Sprintf("%v\n%v", ErrInvalidEventType, _type))
		}
		table[i] = EventList{
			Id:   id,
			Type: EventType(_type),
		}
	}
	a.Cache["EventList"] = table
	return table
}
