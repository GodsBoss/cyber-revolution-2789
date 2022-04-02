package game

import "github.com/GodsBoss/gggg/pkg/dom"

func (g *game) SetOutput(ctx2d *dom.Context2D) {
	g.ctx2d = ctx2d
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	return 320, 200, 1.0, 1.0
}

func (g *game) Render() {}
