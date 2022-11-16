package asset

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"wf_api/util"
)

// errors
var (
	ErrInvalidEventType = errors.New("无效的事件类型")
	ErrItemRarity       = errors.New("无效的稀有度")
	ErrShopItemType     = errors.New("无效的商店物品类型")
	ErrShopCurrencyType = errors.New("无效的货币类型")
)

type Shop interface {
	GetShopType() int
}

type Event struct {
	Name                string
	StartTime           time.Time
	PlayableEndTime     time.Time
	ExchangeableEndTime time.Time
}

type EventType int

const (
	EventTypeAdvent EventType = iota
	EventTypeRanking
	EventTypeStory
	EventTypeDailyWeek
	EventTypeChallengeDungeon
	EventTypeDailyExpMana
	EventTypeWorldStory
	EventTypeTowerDungeon
	EventTypeExpertSingle
)

type Item struct {
	Type   ItemType
	Name   string
	Id     int
	Count  int
	Rarity int
}

type ShopItem struct {
	Items     []Item
	Name      string
	Rarity    int
	StartTime time.Time
	EndTime   *time.Time
	Costs     []Item
}

func (a *Asset) parseShopItem(strings []string, posItem, posRarity, posTime, costPos int) *ShopItem {
	itemTable := a.GetItemListTable()

	//item
	var items []Item
	for i := posItem; i < posItem+3*6; i += 3 {
		if itemType := strings[i]; itemType != "(None)" {
			num, err := strconv.Atoi(itemType)
			if err != nil {
				panic(fmt.Sprintf("%v\n%v", err, i))
			}
			_type := ItemType(num)
			var itemId = 0
			var count = 1
			switch _type {
			case ITEM, CHARACTER, EQUIPMENT:
				itemId, err = strconv.Atoi(strings[i+1])
				if err != nil {
					panic(err)
				}
				fallthrough
			case EXP, MANA:
				count, err = strconv.Atoi(strings[i+2])
				if err != nil {
					panic(err)
				}
			default:
				panic(fmt.Sprintf("%v\n%v", ErrShopItemType, _type))
			}
			items = append(items, Item{
				Id:    itemId,
				Type:  _type,
				Count: count,
			})
		}

	}

	//rarity
	rarity, err := strconv.Atoi(strings[posRarity])
	if err != nil {
		panic(err)
	}
	if rarity < 1 || rarity > 5 {
		panic(fmt.Sprintf("%v\n%v", ErrItemRarity, rarity))
	}

	//time
	startTime := util.ParseIso(strings[posTime])
	if err != nil {
		panic(err)
	}
	var endTime *time.Time
	if endTimeStr := strings[posTime+1]; endTimeStr != "(None)" {
		iso := util.ParseIso(endTimeStr)
		if err != nil {
			panic(err)
		}
		endTime = &iso
	}

	//costs
	var costs []Item
	if priceType := strings[costPos]; priceType != "(None)" {
		var _type ItemType
		switch priceType {
		case "0":
			_type = STONE
		case "1":
			_type = MANA
		case "2":
			_type = BondToken
		default:
			panic(fmt.Sprintf("%v\n%v", ErrShopCurrencyType, _type))
		}
		count, err := strconv.Atoi(strings[costPos+1])
		if err != nil {
			panic(err)
		}
		costs = append(costs, Item{
			Type:  _type,
			Count: count,
		})
		fmt.Println(costs)
	}
	if err != nil {
		panic(err)
	}
	for i := costPos + 2; i < costPos+2+2*4; i += 2 {
		if costId := strings[i]; costId != "(None)" {
			id, err := strconv.Atoi(costId)
			if err != nil {
				panic(err)
			}
			count, err := strconv.Atoi(strings[i+1])
			if err != nil {
				panic(err)
			}
			item := itemTable[id]
			costs = append(costs, Item{
				Type:   ITEM,
				Id:     id,
				Name:   item.Name,
				Rarity: item.Rarity,
				Count:  count,
			})
		}
	}
	return &ShopItem{
		Items:     items,
		Name:      strings[4],
		Rarity:    rarity,
		StartTime: startTime,
		EndTime:   endTime,
		Costs:     costs,
	}
}

type ItemType int

const (
	ITEM      ItemType = iota //材料
	EXP                       //经验
	MANA                      //玛纳
	CHARACTER                 //角色
	EQUIPMENT                 //装备
	STONE                     //星导石
	BondToken                 //兑换券
)

func (t ItemType) Name() string {
	switch t {
	case ITEM:
		return "材料"
	case EXP:
		return "经验"
	case MANA:
		return "玛纳"
	case CHARACTER:
		return "角色"
	case EQUIPMENT:
		return "装备"
	case STONE:
		return "星导石"
	case BondToken:
		return "兑换券"
	default:
		return ""
	}
}
