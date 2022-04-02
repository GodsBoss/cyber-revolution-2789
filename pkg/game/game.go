package game

import (
	"fmt"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/scale"
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
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

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {
	g.states.receiveKeyEvent(event)
}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {
	g.states.receiveMouseEvent(event)
}

func (g *game) TicksPerSecond() int {
	return ticksPerSecond
}

const ticksPerSecond = 50

func (g *game) Tick(ms int) {
	g.states.tick(ms)
}

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

func (g *game) Render() {
	g.states.render(g.ctx2d)
}
