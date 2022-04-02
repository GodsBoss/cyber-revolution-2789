package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingID = "playing"

type statePlaying struct{}

func (state *statePlaying) init() {}

func (state *statePlaying) tick(ms int) (next string) {
	return ""
}

func (state *statePlaying) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *statePlaying) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlaying) renderable() canvas2drendering.Renderable {
	return canvas2drendering.NopRenderable()
}
