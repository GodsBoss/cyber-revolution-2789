package game

import (
	"fmt"

	"github.com/GodsBoss/gggg/pkg/dom"
)

func (g *game) SetOutput(ctx2d *dom.Context2D) {
	g.ctx2d = ctx2d
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	rw, rh, s := g.scaler.Scale(availableWidth, availableHeight)
	fmt.Println(rw, rh, s)
	return rw, rh, s, s
}

type scaler interface {
	Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scale float64)
}

func (g *game) Render() {}
