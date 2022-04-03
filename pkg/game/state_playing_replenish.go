package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingReplenishID = "playing_replenish"

type statePlayingReplenish struct {
	spriteFactory *spriteFactory

	data *playingData

	beam *beam

	addedPerson bool

	remainingCheats int
	nextCheat       int
}

func (state *statePlayingReplenish) init() {
	state.beam = newBeam(350)
	state.addedPerson = false

	state.remainingCheats = 0
	if len(state.data.cheats.availableCheats) < maxCheats {
		state.remainingCheats = 1
	}

	state.nextCheat = nextCheatInterval
}

func (state *statePlayingReplenish) tick(ms int) (next string) {
	state.data.tick(ms)

	if state.beam.isOver() && !state.data.isAnyPersonMoving() && state.remainingCheats == 0 {
		return statePlayingInteractionID
	}

	state.beam.tick(ms)

	if state.beam.isBeamed() && !state.addedPerson {
		state.data.addRandomPerson(0)
		state.addedPerson = true
	}

	state.nextCheat -= ms
	if state.nextCheat <= 0 && state.remainingCheats > 0 {
		state.nextCheat += nextCheatInterval
		state.remainingCheats--
		state.data.addRandomCheat()
	}

	return ""
}

func (state *statePlayingReplenish) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *statePlayingReplenish) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingReplenish) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)
	renderables = append(renderables, state.beam.rendered(state.spriteFactory, 0, personRenderY))

	return renderables
}

const nextCheatInterval = 250
