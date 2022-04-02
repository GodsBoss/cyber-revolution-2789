package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingReplenishID = "playing_replenish"

type statePlayingReplenish struct {
	spriteFactory *spriteFactory

	data *playingData

	beamState     string
	nextBeamState int
	beamAnimation animation.Frames

	remainingCheats int
	nextCheat       int
}

func (state *statePlayingReplenish) init() {
	state.beamState = beamStates[0]
	state.nextBeamState = beamStateSwitchInterval
	state.beamAnimation = animation.NewFrames(3, 75)
	state.beamAnimation.Randomize()

	state.remainingCheats = 1
	state.nextCheat = 1000
}

func (state *statePlayingReplenish) tick(ms int) (next string) {
	state.data.tick(ms)
	state.beamAnimation.Tick(ms)

	if state.beamState == "" && !state.data.isAnyPersonMoving() && state.remainingCheats == 0 {
		return statePlayingInteractionID
	}

	state.nextBeamState -= ms
	if state.nextBeamState <= 0 {
		switch state.beamState {
		case beamStates[0]:
			state.beamState = beamStates[1]
			state.nextBeamState = beamStateSwitchInterval
		case beamStates[1]:
			state.beamState = beamStates[2]
			state.nextBeamState = beamStateSwitchInterval
			state.data.addRandomPerson(0)
		case beamStates[2]:
			state.beamState = ""
		}
	}

	state.nextCheat -= ms
	if state.nextCheat <= 0 && state.remainingCheats > 0 {
		state.nextCheat += 1000
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

	if state.beamState != "" {
		renderables = append(
			renderables,
			state.spriteFactory.create("beam_"+state.beamState, 0, personRenderY, state.beamAnimation.Frame()),
		)
	}

	return renderables
}

var beamStates = []string{
	"start",
	"middle",
	"end",
}

const beamStateSwitchInterval = 500
