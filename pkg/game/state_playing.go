package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingID = "playing"

type playingData struct {
	personQueue *personQueue
	cheats      []cheat

	// selectedCheat is the index of the chosen cheat. Contains -1 if no cheat is selected.
	selectedCheat int

	// selectedCheatTargets are the indexes of the chosen cheat targets.
	selectedCheatTargets []int
}

type statePlaying struct {
	spriteFactory *spriteFactory

	data *playingData
}

func (state *statePlaying) init() {
	state.data = &playingData{
		personQueue: &personQueue{
			persons: make([]person, 0),
		},
	}
	state.addRandomPerson(180)
	state.addPlayer(140)
	state.addRandomPerson(100)

	state.data.selectedCheat = noCheatSelected
	state.data.cheats = make([]cheat, 0)
	state.addRandomCheat()
	state.addRandomCheat()
}

func (state *statePlaying) tick(ms int) (next string) {
	state.data.personQueue.Tick(ms)

	for i := range state.data.cheats {
		state.data.cheats[i].markerAnimation.Tick(ms)
	}

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
		if state.data.selectedCheat == noCheatSelected {
			state.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if len(state.data.selectedCheatTargets) == len(allCheats[state.data.cheats[state.data.selectedCheat].id].targets) {
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
	renderables = append(renderables, state.persons()...)
	renderables = append(renderables, state.renderedCheats()...)

	return renderables
}
