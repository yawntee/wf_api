package asset

import (
	"fmt"
	"strconv"
)

type EventItemShop struct {
	EventId   int
	EventType EventType
	ShopItem
}

type EventItemShopTable map[int]EventItemShop

func (a *Asset) GetEventItemShopTable() EventItemShopTable {
	if cache, ok := a.Cache["EventItemShop"].(EventItemShopTable); ok {
		return cache
	}
	reader := a.GetTableFile("/shop/event_item_shop")
	intMap := parseIntMap(reader)
	table := make(EventItemShopTable)
	for id, strings := range intMap {
		//event
		eventType, err := strconv.Atoi(strings[0])
		if err != nil {
			panic(err)
		}
		if EventType(eventType) > EventTypeExpertSingle {
			panic(fmt.Sprintf("%v\n%v", ErrInvalidEventType, eventType))
		}
		eventId, err := strconv.Atoi(strings[1])
		if err != nil {
			panic(err)
		}
		item := a.parseShopItem(strings, 6, 30, 13, 24, 27, 14)
		item.Id = id
		if err != nil {
			panic(err)
		}
		table[id] = EventItemShop{
			EventId:   eventId,
			EventType: EventType(eventType),
			ShopItem:  *item,
		}
	}
	a.Cache["EventItemShop"] = table
	return table
}
