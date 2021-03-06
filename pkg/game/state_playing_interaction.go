package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingInteractionID = "playing"

type statePlayingInteraction struct {
	spriteFactory *spriteFactory
	kc            *killChamber

	data *playingData

	buttonDiscardMarkerAnimation animation.Frames
	buttonPassMarkerAnimation    animation.Frames

	cheatDiscarded bool
}

func (state *statePlayingInteraction) init() {
	state.data.unselectCheat()

	state.buttonDiscardMarkerAnimation = animation.NewFrames(3, 80)
	state.buttonPassMarkerAnimation = animation.NewFrames(3, 80)

	state.cheatDiscarded = false
}

func (state *statePlayingInteraction) tick(ms int) (next string) {
	state.data.tick(ms)
	state.kc.tick(ms)

	state.buttonDiscardMarkerAnimation.Tick(ms)
	state.buttonPassMarkerAnimation.Tick(ms)

	return ""
}

func (state *statePlayingInteraction) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *statePlayingInteraction) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	if event.Type == interaction.MouseUp {
		// No cheat selected and pass button pressed, directly enter kill state.
		if buttonPassRectangle.withinBounds(event.X, event.Y) && state.data.isNoCheatSelected() {
			return statePlayingKillID
		}

		// Cancel cheat button pressed.
		if state.data.isCancelCheat(event.X, event.Y) {
			state.data.unselectCheat()
			return ""
		}

		// Cheat selected and discard button pressed, remove cheat.
		if buttonDiscardRectangle.withinBounds(event.X, event.Y) && !state.data.isNoCheatSelected() {
			state.data.removeSelectedCheat()
			state.data.cheats.unselectCheat()
			state.cheatDiscarded = true
			return ""
		}

		// No cheat selected yet, so try to select one.
		if state.data.isNoCheatSelected() {
			state.data.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if state.data.isCheatActivationClick(event.X, event.Y) {
			return statePlayingInvokeActionID
		}

		// Try to select target.
		if !state.data.areAllTargetsSelected() {
			state.trySelectTarget(event.X, event.Y)
		}
	}

	return ""
}

func (state *statePlayingInteraction) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
		state.kc.render(state.spriteFactory, false),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, true)...)

	renderables = append(
		renderables,
		state.spriteFactory.create("button_discard", ButtonDiscardRenderX, cheatRenderY, 0),
		state.spriteFactory.create("button_pass", ButtonPassRenderX, cheatRenderY, 0),
	)

	if state.data.isNoCheatSelected() {
		renderables = append(
			renderables,
			state.spriteFactory.create("cheat_marker", ButtonPassRenderX-3, cheatRenderY-3, state.buttonPassMarkerAnimation.Frame()),
		)
	}

	if !state.data.isNoCheatSelected() && !state.cheatDiscarded {
		renderables = append(
			renderables,
			state.spriteFactory.create("cheat_marker", ButtonDiscardRenderX-3, cheatRenderY-3, state.buttonDiscardMarkerAnimation.Frame()),
		)
	}

	return renderables
}

func (state *statePlayingInteraction) unselectCheat() {
	state.data.unselectCheat()
}

const (
	ButtonPassRenderX    = 250
	ButtonDiscardRenderX = 280
)

var buttonPassRectangle = rectangle{
	x:      ButtonPassRenderX,
	y:      cheatRenderY,
	width:  24,
	height: 24,
}

var buttonDiscardRectangle = rectangle{
	x:      ButtonDiscardRenderX,
	y:      cheatRenderY,
	width:  24,
	height: 24,
}
