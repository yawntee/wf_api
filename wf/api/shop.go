package api

import (
	"github.com/zeromicro/go-zero/core/mathx"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"math"
	"os"
	"time"
	"wf_api/wf"
	"wf_api/wf/internal/asset"
)

type ShopEvent struct {
	EventIds  []int           `json:"event_ids"`
	EventType asset.EventType `json:"event_type"`
}

type Shop struct {
	Types  []int            `json:"types"`
	Events []ShopEvent      `json:"events"`
	Ids    []int            `json:"ids"`
	Name   string           `json:"name"`
	Items  []asset.ShopItem `json:"items"`
}

func IsInited() bool {
	_, err := os.Stat(asset.AssetDownloadDir)
	if err != nil {
		return false
	}
	return true
}

func BossShops() []*Shop {
	if cache, ok := asset.GlobalAsset.Cache["BossShops"].([]*Shop); ok {
		return cache
	}
	var entries = make(map[int]*Shop)
	bossTable := asset.GlobalAsset.GetBossTable()
	bossQuestTable := asset.GlobalAsset.GetBossQuestTable()
	itemTable := asset.GlobalAsset.GetBossCoinShopTable()
	now := time.Now()
	for id, boss := range bossTable {
		if quest, ok := bossQuestTable[id]; ok && quest[0].StartTime.Before(now) && (quest[0].EndTime == nil || quest[0].EndTime.After(now)) {
			entries[boss.BossShopId] = &Shop{
				Types:  []int{asset.ShopTypeBossCoin},
				Events: make([]ShopEvent, 0),
				Ids:    []int{boss.BossShopId},
				Name:   boss.Name,
			}
		}
	}
	for _, item := range itemTable {
		if item.StartTime.Before(now) && (item.EndTime == nil || item.EndTime.After(now)) {
			if shop, ok := entries[item.BossShopId]; ok {
				shop.Items = append(shop.Items, item.ShopItem)
			}
		}
	}
	rs := maps.Values(entries)
	asset.GlobalAsset.Cache["BossShops"] = rs
	return rs
}

func EventShops() []*Shop {
	if cache, ok := asset.GlobalAsset.Cache["EventShops"].([]*Shop); ok {
		return cache
	}
	type Key struct {
		id  int
		typ asset.EventType
	}
	var entries = make(map[Key]*Shop)
	//获取兑换项信息
	itemTable := asset.GlobalAsset.GetEventItemShopTable()
	//筛选有效活动
	now := time.Now()
	adventQuestTable := asset.GlobalAsset.GetAdventEventQuestTable()
	for id, event := range asset.GlobalAsset.GetAdventEventTable() {
		if _, ok := adventQuestTable[id]; ok && event.StartTime.Before(now) && event.ExchangeableEndTime.After(now) {
			entries[Key{
				id:  id,
				typ: event.Type,
			}] = &Shop{
				Types: []int{asset.ShopTypeEvent},
				Events: []ShopEvent{
					{EventIds: []int{id}, EventType: event.Type},
				},
				Ids:  make([]int, 0),
				Name: event.Name,
			}
		}
	}
	storyQuestTable := asset.GlobalAsset.GetStoryEventQuestTable()
	for id, event := range asset.GlobalAsset.GetStoryEventTable() {
		if _, ok := storyQuestTable[id]; ok && event.StartTime.Before(now) && event.ExchangeableEndTime.After(now) {
			entries[Key{
				id:  id,
				typ: event.Type,
			}] = &Shop{
				Types: []int{asset.ShopTypeEvent},
				Events: []ShopEvent{
					{EventIds: []int{id}, EventType: event.Type},
				},
				Ids:  make([]int, 0),
				Name: event.Name,
			}
		}
	}
	for id, event := range asset.GlobalAsset.GetCollectItemEventTable() {
		if event.StartTime.Before(now) && event.ExchangeableEndTime.After(now) {
			entries[Key{
				id:  id,
				typ: event.Type,
			}] = &Shop{
				Types: []int{asset.ShopTypeEvent},
				Events: []ShopEvent{
					{EventIds: []int{id}, EventType: event.Type},
				},
				Ids:  make([]int, 0),
				Name: event.Name,
			}
		}
	}
	//将物品放入对应活动商店
	for _, item := range itemTable {
		if entry, ok := entries[Key{
			id:  item.EventId,
			typ: item.EventType,
		}]; ok {
			entry.Items = append(entry.Items, item.ShopItem)
		}
	}
	//清除没有商店的活动
	for k, shop := range entries {
		if shop.Items == nil {
			delete(entries, k)
		}
	}
	//获取物品库存
	rs := maps.Values(entries)
	asset.GlobalAsset.Cache["EventShops"] = rs
	return rs
}

type BuyingShop struct {
	Types  []int            `mapstructure:"types"`
	Events []map[string]any `mapstructure:"events"`
	Ids    []int            `mapstructure:"ids"`
	Items  []int            `mapstructure:"items"`
}

func BulkBuying(c *wf.Client, shops []BuyingShop) error {
	itemTable := asset.GlobalAsset.GetItemListTable()
	eventShopTable := asset.GlobalAsset.GetEventItemShopTable()
	bossShopTable := asset.GlobalAsset.GetBossCoinShopTable()
	gameData, err := c.LoadGameData()
	if err != nil {
		return err
	}
	retained := gameData.ItemList
	for _, shop := range shops {
		var bought = true
		for bought {
			bought = false
			sales := c.GetSaleList(shop.Types, shop.Ids, shop.Events).SalesList
			for _, sale := range sales {
				if slices.Contains(shop.Items, sale.ShopItemId) {
					switch sale.StockQuantity {
					//无货
					case 0:
						continue
						//无限
					case -1:
						sale.StockQuantity = math.MaxInt
					}
					var item asset.ShopItem
					if len(shop.Ids) != 0 {
						item = bossShopTable[sale.ShopItemId].ShopItem
					} else {
						item = eventShopTable[sale.ShopItemId].ShopItem
					}
					//货币不足
					cost := item.Costs[0]
					if retained[cost.Id] < cost.Count {
						continue
					}
					//已达上限
					claim := item.Items[0]
					var minus int
					if claim.Type == asset.ITEM {
						minus = itemTable[claim.Id].MaxCount - retained[claim.Id]
						if minus < claim.Count {
							continue
						}
					}
					//库存量
					buyingCount := mathx.MinInt(sale.StockQuantity, item.MaxCount)
					//可购买量
					buyingCount = mathx.MinInt(buyingCount, retained[cost.Id]/cost.Count)
					if claim.Type == asset.ITEM {
						buyingCount = mathx.MinInt(buyingCount, minus/claim.Count)
					}
					//最大一次99
					buyingCount = mathx.MinInt(buyingCount, 99)
					c.Buy(shop.Types, sale.ShopItemId, buyingCount)
					retained[cost.Id] -= buyingCount * cost.Count
					if claim.Type == asset.ITEM {
						retained[claim.Id] += buyingCount * claim.Count
					}
					bought = true
				}
			}
		}
	}
	return nil
}
