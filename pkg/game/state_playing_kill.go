package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingKillID = "playing_kill"

// statePlayingKill is the state responsible for killing the person from the first position of the queue.
type statePlayingKill struct {
	spriteFactory *spriteFactory
	kc            *killChamber

	data *playingData

	killState       string
	nextKillState   int
	killAnimation   animation.Frames
	killFadingFrame int
}

func (state *statePlayingKill) init() {
	state.data.setMostRightX(killChamberX)
	state.data.personQueue.calculateDesiredX()

	state.killState = killStates[0]
	state.nextKillState = 500
	state.killFadingFrame = 0
	state.killAnimation = animation.NewFrames(3, 75)
}

func (state *statePlayingKill) tick(ms int) (next string) {
	state.data.tick(ms)
	state.kc.tick(ms)
	state.killAnimation.Tick(ms)

	state.nextKillState -= ms
	if state.nextKillState <= 0 {
		switch state.killState {
		case killStates[0]:
			state.nextKillState = 500
			state.killState = killStates[1]
		case killStates[1]:
			state.nextKillState = 100
			state.killState = killStates[2]
		case killStates[2]:
			state.nextKillState = 100
			state.killFadingFrame++
			if state.killFadingFrame > 7 {
				state.killState = ""
			}
		}
	}

	if !state.data.isAnyPersonMoving() && state.killState == "" {
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
		state.kc.render(state.spriteFactory, true),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory, false)...)

	switch state.killState {
	case killStates[0]:
		renderables = append(
			renderables,
			state.spriteFactory.create("kill_"+killStates[0], killChamberX, personRenderY, state.killAnimation.Frame()),
		)
	case killStates[1]:
		renderables = append(
			renderables,
			state.spriteFactory.create("kill_"+killStates[1], killChamberX, personRenderY, state.killAnimation.Frame()),
		)
	case killStates[2]:
		renderables = append(
			renderables, state.spriteFactory.create("kill_"+killStates[2], killChamberX, personRenderY, state.killFadingFrame),
		)
	}
	return renderables
}

const killChamberX = 280.0

var killStates = []string{
	"prolog",
	"blazing",
	"fading",
}
