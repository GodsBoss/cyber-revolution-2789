package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingID = "playing"

type statePlaying struct {
	spriteFactory *spriteFactory

	personQueue *personQueue
	cheats      []cheat

	// selectedCheat is the index of the chosen cheat. Contains -1 if no cheat is selected.
	selectedCheat int

	// selectedCheatTargets are the indexes of the chosen cheat targets.
	selectedCheatTargets []int
}

func (state *statePlaying) init() {
	state.personQueue = &personQueue{
		persons: []person{
			{
				Type: personTypeGreenAlien,
				x:    100,
			},
			{
				Type: personTypePlayer,
				x:    140,
			},
			{
				Type: personTypeGreenAlien,
				x:    180,
			},
		},
	}
	state.personQueue.calculateDesiredX()

	state.selectedCheat = noCheatSelected
	state.cheats = make([]cheat, 0)
	state.addRandomCheat()
	state.addRandomCheat()
}

func (state *statePlaying) tick(ms int) (next string) {
	state.personQueue.Tick(ms)
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
		if state.selectedCheat == noCheatSelected {
			state.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if len(state.selectedCheatTargets) == len(allCheats[state.cheats[state.selectedCheat].id].targets) {
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
