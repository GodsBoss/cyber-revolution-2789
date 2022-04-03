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
}

func (state *statePlayingDead) init() {}

func (state *statePlayingDead) tick(ms int) (next string) {
	state.data.tick(ms)
	state.kc.tick(ms)

	return ""
}

func (state *statePlayingDead) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	if event.Type == interaction.KeyUp && event.Key == "t" {
		return stateTitleID
	}

	return ""
}

func (state *statePlayingDead) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingDead) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
		state.kc.render(state.spriteFactory, false),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)

	return renderables
}
