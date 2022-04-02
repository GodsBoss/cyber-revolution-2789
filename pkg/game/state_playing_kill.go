package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingKillID = "playing_kill"

// statePlayingKill is the state responsible for killing the person from the first position of the queue.
type statePlayingKill struct {
	spriteFactory *spriteFactory

	data *playingData
}

func (state *statePlayingKill) init() {
	state.data.setMostRightX(killChamberX)
	state.data.personQueue.calculateDesiredX()
}

func (state *statePlayingKill) tick(ms int) (next string) {
	state.data.tick(ms)

	if !state.data.isAnyPersonMoving() {
		state.data.setMostRightX(killChamberX - personHorizontalDistance)
		state.data.removeMostRightPerson()

		if state.data.isPlayerAlive() {
			return statePlayingReplenishID
		}

		return statePlayingDeadID
	}

	return ""
}

// receiveKeyEvent does nothing.
func (state *statePlayingKill) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	return ""
}

// receiveMouseEvent does nothing.
func (state *statePlayingKill) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	return ""
}

func (state *statePlayingKill) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)

	return renderables
}

const killChamberX = 280.0
