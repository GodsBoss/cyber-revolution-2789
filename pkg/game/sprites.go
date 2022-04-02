package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

type spriteFactory struct {
	source *dom.Image
	scaler scaler
	infos  map[string]SpriteInfo
}

func (sf *spriteFactory) create(id string, x, y, frame int) canvas2drendering.Renderable {
	info, ok := sf.infos[id]
	if !ok {
		return canvas2drendering.NopRenderable()
	}
	return sprite{
		source: sf.source,
		scaler: sf.scaler,

		x:     x,
		y:     y,
		frame: frame,

		sx: info.X,
		sy: info.Y,
		sw: info.W,
		sh: info.H,
	}
}

type sprite struct {
	source *dom.Image
	scaler scaler

	x     int
	y     int
	frame int

	sx int
	sy int
	sw int
	sh int
}

func (s sprite) Render(output *dom.Context2D) {
	scale := s.scaler.Scale()
	x := s.x * scale
	y := s.y * scale
	w := s.sw * scale
	h := s.sh * scale

	sx := s.sx + s.frame*s.sw

	output.DrawImage(s.source, sx, s.sy, s.sw, s.sh, x, y, w, h)
}

type SpriteInfo struct {
	X int
	Y int
	W int
	H int
}
