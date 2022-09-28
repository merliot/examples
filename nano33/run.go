//go:build !tinygo
// +build !tinygo

package nano33

import "github.com/merliot/merle"

func (n *nano33) init(p *merle.Packet) {
}

func (n *nano33) run(p *merle.Packet) {
	merle.RunForever(p)
}
