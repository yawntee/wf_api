package asset

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"wf_api/server/util"
)

// errors
var (
	ErrInvalidEventType = errors.New("无效的事件类型")
	ErrItemRarity       = errors.New("无效的稀有度")
	ErrShopItemType     = errors.New("无效的商店物品类型")
	ErrShopCurrencyType = errors.New("无效的货币类型")
)

type Event struct {
	Name                string    //事件名称
	StartTime           time.Time //事件起始时间
	PlayableEndTime     time.Time //参与结束时间
	ExchangeableEndTime time.Time //兑换结束时间
	Type                int       //商店类型
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

const (
	ShopTypeBossCoin = 7
	ShopTypeEvent    = 4
)

// Item 物品
type Item struct {
	Type   ItemType `json:"type"`   //物品类型
	Name   string   `json:"name"`   //物品名称
	Id     int      `json:"id"`     //物品id
	Count  int      `json:"count"`  //物品数量
	Rarity uint8    `json:"rarity"` //物品稀有度
}

// ShopItem 商店兑换项
type ShopItem struct {
	Id        int        `json:"id"`                  //兑换项id
	Items     []Item     `json:"items"`               //获得的物品
	Name      string     `json:"name"`                //列表项名称
	Rarity    uint8      `json:"rarity"`              //稀有度
	StartTime time.Time  `json:"start_time"`          //起始时间
	EndTime   *time.Time `json:"end_time"`            //结束时间
	Costs     []Item     `json:"costs"`               //花费
	Stock     int        `json:"stock,omitempty"`     //库存
	MaxCount  int        `json:"max_count,omitempty"` //最大可兑换数量
}

func (a *Asset) parseShopItem(params []string, namePos, itemPos, rarityPos, timePos, maxCountPos, costPos int) *ShopItem {
	itemTable := a.GetItemListTable()

	//item
	var items []Item
	for i := itemPos; i < itemPos+3*6; i += 3 {
		if itemType := params[i]; itemType != "(None)" {
			num, err := strconv.Atoi(itemType)
			if err != nil {
				panic(fmt.Sprintf("%v\n%v", err, i))
			}
			_type := ItemType(num)
			var itemId = 0
			var count = 1
			switch _type {
			case ITEM, CHARACTER, EQUIPMENT:
				itemId, err = strconv.Atoi(params[i+1])
				if err != nil {
					panic(err)
				}
				fallthrough
			case EXP, MANA:
				count, err = strconv.Atoi(params[i+2])
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
	rarity, err := strconv.ParseUint(params[rarityPos], 10, 8)
	if err != nil {
		panic(err)
	}
	if rarity < 1 || rarity > 5 {
		panic(fmt.Sprintf("%v\n%v", ErrItemRarity, rarity))
	}

	//time
	startTime := util.ParseIso(params[timePos])
	if err != nil {
		panic(err)
	}
	var endTime *time.Time
	if endTimeStr := params[timePos+1]; endTimeStr != "(None)" {
		iso := util.ParseIso(endTimeStr)
		if err != nil {
			panic(err)
		}
		endTime = &iso
	}

	//maxCount
	maxCount, err := strconv.Atoi(params[maxCountPos])
	if err != nil {
		panic(err)
	}

	//costs
	var costs []Item
	if priceType := params[costPos]; priceType != "(None)" {
		var _type ItemType
		switch priceType {
		case "0":
			_type = STONE
		case "1":
			_type = MANA
		case "2":
			_type = BondToken
		default:
			panic(fmt.Sprintf("%v\n%v", ErrShopCurrencyType, params))
		}
		count, err := strconv.Atoi(params[costPos+1])
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
		if costId := params[i]; costId != "(None)" {
			id, err := strconv.Atoi(costId)
			if err != nil {
				panic(err)
			}
			count, err := strconv.Atoi(params[i+1])
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
		Name:      params[namePos],
		Rarity:    uint8(rarity),
		StartTime: startTime,
		EndTime:   endTime,
		Costs:     costs,
		MaxCount:  maxCount,
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
