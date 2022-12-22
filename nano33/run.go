//go:build !tinygo
// +build !tinygo

package nano33

import "github.com/merliot/merle"

func (n *nano33) ConnectAP(ssid, pass string) string {
	return ""
}

func (n *nano33) run(p *merle.Packet) {
	merle.RunForever(p)
}
