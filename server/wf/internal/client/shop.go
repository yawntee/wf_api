package client

import (
	"wf_api/server/wf/internal"
)

type ShopStock struct {
	SalesList []struct {
		ShopItemId           int `mapstructure:"shop_item_id"`
		StockQuantity        int `mapstructure:"stock_quantity"`
		TodayPurchaseNum     int `mapstructure:"today_purchase_num"`
		ThisMonthPurchaseNum int `mapstructure:"this_month_purchase_num"`
		TotalPurchaseNum     int `mapstructure:"total_purchase_num"`
		GroupInfo            struct {
			GroupTotalStockQuantity int  `mapstructure:"group_total_stock_quantity"`
			GroupTotalPurchaseNum   int  `mapstructure:"group_total_purchase_num"`
			MultiStage              bool `mapstructure:"multi_stage"`
		} `mapstructure:"group_info"`
		ShopType int `mapstructure:"shop_type"`
	} `mapstructure:"sales_list"`
}

func (c *Client) GetSaleList(types, ids []int, events []map[string]any) *ShopStock {
	err := c.SignUp()
	if err != nil {
		return nil
	}
	body := map[string]any{
		"viewer_id":                   c.viewerId,
		"shop_types":                  types,
		"browse_treasure_flag":        true,
		"boss_coin_shop_category_ids": ids,
		"event_list":                  events,
	}
	var resp GameResp[ShopStock]
	PostMsgpack(c, "https://shijtswygamegf.leiting.com//api/index.php/shop/get_sales_list", body, &resp, c.SignReqWithViewerId)
	return &resp.Data
}

func (c *Client) Buy(shopType []int, itemId, count int) *internal.ItemClaimedData {
	err := c.SignUp()
	if err != nil {
		return nil
	}
	c.apiCount++
	body := map[string]any{
		"shop_item_id": itemId,
		"api_count":    c.apiCount,
		"number":       count,
		"shop_type":    shopType[0],
		"viewer_id":    c.viewerId,
	}
	var resp GameResp[internal.ItemClaimedData]
	PostMsgpack(c, "https://shijtswygamegf.leiting.com//api/index.php/shop/buy", body, &resp, c.SignReqWithViewerId)
	return &resp.Data
}
