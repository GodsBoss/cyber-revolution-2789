package game

import (
	"math/rand"

	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingID = "playing"

type statePlaying struct {
	spriteFactory *spriteFactory

	data *playingData
}

func (state *statePlaying) init() {
	state.data = &playingData{}
	state.data.init()

	state.addRandomPerson(180)
	state.addPlayer(140)
	state.addRandomPerson(100)

	state.addRandomCheat()
	state.addRandomCheat()
}

func (state *statePlaying) tick(ms int) (next string) {
	state.data.tick(ms)

	return ""
}

func (state *statePlaying) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	if event.Type == interaction.KeyUp && event.Key == "Escape" {
		state.unselectCheat()
	}
	return ""
}

func (state *statePlaying) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	if event.Type == interaction.MouseUp {
		// No cheat selected yet, so try to select one.
		if state.data.isNoCheatSelected() {
			state.data.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if state.data.areAllTargetsSelected() {
			state.tryActivateCheat(event.X, event.Y)
			return ""
		}

		// Try to select target.
		state.trySelectTarget(event.X, event.Y)
	}

	return ""
}

func (state *statePlaying) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory)...)

	return renderables
}

func (state *statePlaying) addRandomCheat() {
	state.data.addRandomCheat()
}

func (state *statePlaying) unselectCheat() {
	state.data.unselectCheat()
}

func (state *statePlaying) tryActivateCheat(x int, y int) {
	state.data.tryActivateCheat(x, y)
}

func (state *statePlaying) addRandomPerson(x float64) {
	ids := make([]string, 0)
	for id := range allPersonTypes {
		if id != personTypePlayer {
			ids = append(ids, id)
		}
	}
	typ := ids[rand.Intn(len(ids))]

	state.addPerson(
		person{
			Type: typ,
			x:    x,
		},
	)
}

func (state *statePlaying) addPlayer(x float64) {
	state.addPerson(
		person{
			Type: personTypePlayer,
			x:    x,
		},
	)
}

func (state *statePlaying) addPerson(p person) {
	state.data.addPerson(p)
}
