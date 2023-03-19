package internal

import (
	"context"
	"encoding/json"
	"wf_api/internal/types"
)

func ReportError(err error) (*types.Resp, error) {
	return ReportMsg(err.Error())
}

func ReportMsg(msg string) (*types.Resp, error) {
	return &types.Resp{
		Code: 1,
		Msg:  msg,
	}, nil
}

func Success(msg string, data any) (*types.Resp, error) {
	return &types.Resp{
		Msg:  msg,
		Data: data,
	}, nil
}

func GetUserId(ctx context.Context) int64 {
	userId, err := ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	return userId
}
