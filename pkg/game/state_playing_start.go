package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingStartID = "playing_start"

// statePlayingStart is entered only when the player runs a new playing session. The person queue is filled and the
// initial set of cheats is added. If everything is in place, the next state, player interaction, is started.
type statePlayingStart struct {
	spriteFactory *spriteFactory

	data *playingData

	remainingPersons []func(x float64)
	nextPerson       int

	remainingCheats int
	nextCheat       int
}

const (
	// timeBetweenPersonSpawns is the time between person spawns in ms.
	timeBetweenPersonSpawns = 1000
	timeBetweenCheatsSpawns = 2000
)

func (state *statePlayingStart) init() {
	state.data.init()

	state.remainingPersons = []func(x float64){
		state.data.addRandomPerson,
		state.data.addRandomPerson,
		state.data.addRandomPerson,
		state.data.addPlayer,
		state.data.addRandomPerson,
		state.data.addRandomPerson,
		state.data.addRandomPerson,
	}

	state.remainingCheats = 4
	state.nextCheat = 1500
}

func (state *statePlayingStart) tick(ms int) (next string) {
	state.data.tick(ms)

	state.nextPerson -= ms
	if state.nextPerson <= 0 && len(state.remainingPersons) > 0 {
		state.nextPerson += timeBetweenPersonSpawns
		state.remainingPersons[0](0)
		state.remainingPersons = state.remainingPersons[1:]
	}

	state.nextCheat -= ms
	if state.nextCheat <= 0 && state.remainingCheats > 0 {
		state.nextCheat += timeBetweenCheatsSpawns
		state.data.addRandomCheat()
		state.remainingCheats--
	}

	if len(state.remainingPersons) == 0 && state.remainingCheats == 0 && !state.data.isAnyPersonMoving() {
		return statePlayingInteractionID
	}

	return ""
}

// receiveKeyEvent does nothing.
func (state *statePlayingStart) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

// receiveMouseEvent does nothing.
func (state *statePlayingStart) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingStart) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory)...)

	return renderables
}