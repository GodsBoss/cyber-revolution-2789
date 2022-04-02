package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingSpawnPersonID = "playing_spawn_person"

type statePlayingSpawnPerson struct {
	spriteFactory *spriteFactory

	data *playingData

	beamState     string
	nextBeamState int
	beamAnimation animation.Frames
}

func (state *statePlayingSpawnPerson) init() {
	state.beamState = beamStates[0]
	state.nextBeamState = beamStateSwitchInterval
	state.beamAnimation = animation.NewFrames(3, 75)
	state.beamAnimation.Randomize()
}

func (state *statePlayingSpawnPerson) tick(ms int) (next string) {
	state.data.tick(ms)
	state.beamAnimation.Tick(ms)

	if state.beamState == "" && !state.data.isAnyPersonMoving() {
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

	return ""
}

func (state *statePlayingSpawnPerson) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

func (state *statePlayingSpawnPerson) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingSpawnPerson) renderable() canvas2drendering.Renderable {
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
