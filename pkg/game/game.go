package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/scale"
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
		states: &states{
			states: map[string]state{
				stateTitleID:    &stateTitle{},
				stateGameOverID: &stateGameOver{},
				statePlayingID:  &statePlaying{},
			},
			currentStateID: stateTitleID,
		},
	}
}

type game struct {
	img    *dom.Image
	ctx2d  *dom.Context2D
	scaler scaler
	states *states
}
