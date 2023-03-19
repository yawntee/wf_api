package wf

import (
	"wf_api/wf/channel"
	"wf_api/wf/internal/client"
)

type Client = client.Client

func NewClient(channel channel.Id) *Client {
	return client.NewClient(channel)
}
