package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const stateTitleID = "title"

type stateTitle struct {
	spriteFactory *spriteFactory
}

func (state *stateTitle) init() {}

func (state *stateTitle) tick(ms int) (next string) {
	return ""
}

func (state *stateTitle) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *stateTitle) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *stateTitle) renderable() canvas2drendering.Renderable {
	return state.spriteFactory.create("background", 0, 0, 0)
}
