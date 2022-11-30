package game

import (
	"context"
	"strconv"
	"sync/atomic"
	"wf_api/server/internal"

	"wf_api/server/internal/svc"
	"wf_api/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskStatusLogic {
	return &GetTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskStatusLogic) GetTaskStatus() (resp *types.Resp, err error) {
	userId := internal.GetUserId(l.ctx)
	if p, ok := internal.TaskMutex[userId]; ok {
		//任务未结束
		if atomic.LoadInt32(p) > 0 {
			return internal.Success("", struct{}{})
		}
	} else {
		//暂无任务
		return internal.Success("", nil)
	}
	data := make(map[string]string)
	for res := range internal.TaskResults[userId] {
		data[strconv.FormatInt(res.Id, 10)] = res.Msg
	}
	//清理
	delete(internal.TaskMutex, userId)
	delete(internal.TaskResults, userId)
	return internal.Success("", data)
}
