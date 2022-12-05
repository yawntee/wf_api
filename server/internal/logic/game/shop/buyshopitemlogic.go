package shop

import (
	"context"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"runtime"
	"runtime/debug"
	"sync/atomic"
	"wf_api/server/internal"
	"wf_api/server/wf"
	"wf_api/server/wf/api"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuyShopItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuyShopItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BuyShopItemLogic {
	return &BuyShopItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuyShopItemLogic) BuyShopItem(req *types.BuyShopItemReq) (resp *types.Resp, err error) {
	if req.GameUserIds == nil || len(req.GameUserIds) == 0 {
		return internal.Success("", nil)
	}
	userId := internal.GetUserId(l.ctx)
	//任务互斥
	if _, ok := internal.TaskMutex[userId]; ok {
		return internal.ReportMsg("请等待上一个任务执行完毕")
	}
	batchStart(l, req, userId)
	return internal.Success("", nil)
}

func batchStart(l *BuyShopItemLogic, req *types.BuyShopItemReq, userId int64) {
	taskCount := int32(len(req.GameUserIds))
	internal.TaskMutex[userId] = &taskCount
	//创建通道
	taskChan := make(chan internal.TaskResult, len(req.GameUserIds))
	internal.TaskResults[userId] = taskChan
	//装配数据
	var selected []api.BuyingShop
	err := mapstructure.Decode(req.Selected, &selected)
	if err != nil {
		panic(err)
	}
	//执行的结果
	var results = make(map[int64]string)
	for _, gameUserId := range req.GameUserIds {
		results[gameUserId] = "成功"
		c, err := internal.ClientPool.GetClient(l.ctx, l.svcCtx.GameUserModel, gameUserId)
		if err != nil {
			results[gameUserId] = err.Error()
		} else {
			go startTask(gameUserId, taskChan, &taskCount, c, selected, results)
		}
	}
}

func startTask(gameUserId int64, taskChan chan internal.TaskResult, taskCount *int32, c *wf.Client, selected []api.BuyingShop, results map[int64]string) {
	defer func() {
		if err := recover(); err != nil {
			var msg string
			switch v := err.(type) {
			case string:
				msg = v
			case error:
				fmt.Println(v)
				msg = v.Error()
			}
			debug.PrintStack()
			taskChan <- internal.TaskResult{
				Id:  gameUserId,
				Msg: msg,
			}
		} else {
			taskChan <- internal.TaskResult{
				Id:  gameUserId,
				Msg: "成功",
			}
		}
		if atomic.AddInt32(taskCount, -1) == 0 {
			close(taskChan)
		}
		runtime.Goexit()
	}()
	//开始购买
	err := api.BulkBuying(c, selected)
	if err != nil {
		results[gameUserId] = err.Error()
		return
	}
}
