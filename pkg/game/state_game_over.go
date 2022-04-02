package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const stateGameOverID = "game_over"

type stateGameOver struct{}

func (state *stateGameOver) init() {}

func (state *stateGameOver) tick(ms int) (next string) {
	return ""
}

func (state *stateGameOver) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *stateGameOver) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *stateGameOver) renderable() canvas2drendering.Renderable {
	return background
}
