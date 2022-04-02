package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const stateTitleID = "title"

type stateTitle struct {
	spriteFactory *spriteFactory

	hoverPlay bool
}

func (state *stateTitle) init() {}

func (state *stateTitle) tick(ms int) (next string) {
	return ""
}

func (state *stateTitle) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *stateTitle) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	if event.Type == interaction.MouseMove {
		state.hoverPlay = playButton.withinBounds(event.X, event.Y)
	}
	if event.Type == interaction.MouseUp && state.hoverPlay {
		return statePlayingStartID
	}
	return ""
}

func (state *stateTitle) renderable() canvas2drendering.Renderable {
	return canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
		state.playButton(),
	}
}

func (state *stateTitle) playButton() canvas2drendering.Renderable {
	id := "play_button"
	if state.hoverPlay {
		id = "play_button_hover"
	}
	return state.spriteFactory.create(id, playButtonX, playButtonY, 0)
}

var playButton = rectangle{
	x:      100,
	y:      100,
	width:  61,
	height: 19,
}

const (
	playButtonX      = 100
	playButtonY      = 100
	playButtonWidth  = 61
	playButtonHeight = 19
)
