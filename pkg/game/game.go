package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
)

// New creates and initializes a new game.
func New(img *dom.Image) dominit.Game {
	return &game{
		img: img,
	}
}

type game struct {
	img   *dom.Image
	ctx2d *dom.Context2D
}
