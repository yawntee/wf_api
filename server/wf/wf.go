package wf

import (
	"wf_api/server/wf/channel"
	"wf_api/server/wf/internal/client"
)

type Client = client.Client

func NewClient(channel channel.Id) *Client {
	return client.NewClient(channel)
}
