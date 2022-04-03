package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingDeadID = "playing_dead"

// statePlayingDead occurs after the player has been killed.
type statePlayingDead struct {
	spriteFactory *spriteFactory
	kc            *killChamber

	data *playingData

	hoverBack bool
}

func (state *statePlayingDead) init() {
	state.hoverBack = false
}

func (state *statePlayingDead) tick(ms int) (next string) {
	state.data.tick(ms)
	state.kc.tick(ms)

	return ""
}

func (state *statePlayingDead) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *statePlayingDead) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	if event.Type == interaction.MouseMove {
		state.hoverBack = playButton.withinBounds(event.X, event.Y)
	}
	if event.Type == interaction.MouseUp && state.hoverBack {
		return stateTitleID
	}
	return ""
}

func (state *statePlayingDead) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
		state.kc.render(state.spriteFactory, false),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)
	renderables = append(renderables, state.backButton())

	return renderables
}

func (state *statePlayingDead) backButton() canvas2drendering.Renderable {
	id := "back_button"
	if state.hoverBack {
		id = "back_button_hover"
	}
	return state.spriteFactory.create(id, playButton.x, playButton.y, 0)
}
