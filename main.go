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
}

func (m *Minecraft) XClient(ctxPtr *context.Context) interface{} {
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{}, ctxPtr)
}

func (c *Client) JoinServer(address string) error {
	cc := bot.NewClient()
	return cc.JoinServer(address)
}

