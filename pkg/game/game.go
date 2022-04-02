package game

import (
	"github.com/GodsBoss/delay-the-inevitable/pkg/scale"
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
)

// New creates and initializes a new game.
func New(img *dom.Image) dominit.Game {
	return &game{
		img: img,
		scaler: scale.ByInteger{
			UnscaledWidth:    320,
			UnscaledHeight:   200,
			HorizontalMargin: 20,
			VerticalMargin:   20,
		},
	}
}

type game struct {
	img    *dom.Image
	ctx2d  *dom.Context2D
	scaler scaler
}
