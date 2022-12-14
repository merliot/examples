package nano33

import "github.com/merliot/merle"

//tinyjson:json
type nano33 struct {
	Msg     string
	TempC   int32
	macAddr string
}

func NewNano33() *nano33 {
	return &nano33{Msg: merle.ReplyState}
}

func (n *nano33) Subscribers() merle.Subscribers {
	return merle.Subscribers{
		merle.CmdInit: merle.NoInit,
		merle.CmdRun:  n.run,
	}
}

func (n *nano33) Assets() merle.ThingAssets {
	return merle.ThingAssets{}
}
