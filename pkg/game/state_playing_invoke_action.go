package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingInvokeActionID = "playing_invoke_action"

type statePlayingInvokeAction struct {
	spriteFactory *spriteFactory
	kc            *killChamber

	data *playingData
}

func (state *statePlayingInvokeAction) init() {
	state.data.activateCheat()
}

func (state *statePlayingInvokeAction) tick(ms int) (next string) {
	state.data.tick(ms)
	state.kc.tick(ms)

	if !state.data.isAnyPersonMoving() {
		return statePlayingKillID
	}

	return ""
}

// receiveKeyEvent does nothing.
func (state *statePlayingInvokeAction) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

// receiveMouseEvent does nothing.
func (state *statePlayingInvokeAction) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingInvokeAction) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
		state.kc.render(state.spriteFactory, false),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)

	return renderables
}
