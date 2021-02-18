package minecraft

import (
	"context"

	"github.com/Tnze/go-mc/bot"
	"github.com/loadimpact/k6/js/common"
	"github.com/loadimpact/k6/js/modules"
)

func init() {
	modules.Register("k6/x/minecraft", new(Minecraft))
}

type Minecraft struct{}

type Client struct {
	cc *bot.Client
}

func (m *Minecraft) XClient(ctxPtr *context.Context) interface{} {
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{cc: bot.NewClient()}, ctxPtr)
}

func (c *Client) JoinServer(address string) error {
	return c.cc.JoinServer(address)
}

