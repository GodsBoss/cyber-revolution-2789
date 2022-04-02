package game

import (
	"math/rand"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
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
			state.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if len(state.data.cheats.selectedCheatTargets) == len(allCheats[state.data.cheats.availableCheats[state.data.cheats.selectedCheat].id].targets) {
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
	renderables = append(renderables, state.renderedPersons()...)
	renderables = append(renderables, state.renderedCheats()...)

	return renderables
}

func (state *statePlaying) addRandomCheat() {
	cheatIDs := make([]string, 0)
	for id := range allCheats {
		cheatIDs = append(cheatIDs, id)
	}
	id := cheatIDs[rand.Intn(len(cheatIDs))]
	newCheat := cheat{
		id:              id,
		markerAnimation: animation.NewFrames(3, 80),
	}
	newCheat.markerAnimation.Randomize()
	state.data.cheats.availableCheats = append(state.data.cheats.availableCheats, newCheat)
}

func (state *statePlaying) trySelectCheat(x int, y int) {
	for i := range state.data.cheats.availableCheats {
		cheatX, cheatY := state.cheatCoords(i)

		cheatBounds := rectangle{
			x:      cheatX,
			y:      cheatY,
			width:  cheatWidth,
			height: cheatHeight,
		}

		if cheatBounds.withinBounds(x, y) {
			state.data.cheats.selectedCheat = i
			return
		}
	}
}

func (state *statePlaying) unselectCheat() {
	state.data.unselectCheat()
}

func (state *statePlaying) tryActivateCheat(x int, y int) {
	cheatX, cheatY := state.cheatCoords(state.data.cheats.selectedCheat)

	cheatBounds := rectangle{
		x:      cheatX,
		y:      cheatY,
		width:  cheatWidth,
		height: cheatHeight,
	}

	if !cheatBounds.withinBounds(x, y) {
		return
	}

	allCheats[state.data.cheats.availableCheats[state.data.cheats.selectedCheat].id].invoke(&state.data.personQueue, state.data.cheats.selectedCheatTargets)

	// Cheat has been used, remove it.
	state.data.cheats.availableCheats = append(
		state.data.cheats.availableCheats[0:state.data.cheats.selectedCheat],
		state.data.cheats.availableCheats[state.data.cheats.selectedCheat+1:]...,
	)
	state.unselectCheat()

	// Person queue probably changed, recalculate.
	state.data.personQueue.calculateDesiredX()

	state.addRandomCheat()
}

func (state *statePlaying) renderedCheats() canvas2drendering.Renderables {
	l := len(state.data.cheats.availableCheats)

	renderables := make(canvas2drendering.Renderables, l)

	for i, cheat := range state.data.cheats.availableCheats {
		x, y := state.cheatCoords(i)

		renderables[i] = state.spriteFactory.create(cheat.SpriteID(), x, y, 0)

		// If no cheat is selected, highlight all cheats as possible user interactions.
		if state.data.isNoCheatSelected() {
			renderables = append(renderables, state.spriteFactory.create("cheat_marker", x-3, y-3, cheat.markerAnimation.Frame()))
		}
	}

	if !state.data.isNoCheatSelected() && len(allCheats[state.data.cheats.availableCheats[state.data.cheats.selectedCheat].id].targets) == len(state.data.cheats.selectedCheatTargets) {
		x, y := state.cheatCoords(state.data.cheats.selectedCheat)
		renderables = append(
			renderables,
			state.spriteFactory.create("cheat_marker", x-3, y-3, state.data.cheats.availableCheats[state.data.cheats.selectedCheat].markerAnimation.Frame()),
		)
	}

	return renderables
}

func (state *statePlaying) cheatCoords(index int) (x int, y int) {
	return state.cheatCoords(index)
}

func (state *statePlaying) renderedPersons() canvas2drendering.Renderables {
	renderables := make(canvas2drendering.Renderables, len(state.data.personQueue.persons))
	for i, person := range state.data.personQueue.persons {
		renderables[i] = state.spriteFactory.create("person_"+person.Type, int(person.x), personRenderY, 0)
	}
	for _, index := range state.data.cheats.selectedCheatTargets {
		p := state.data.personQueue.persons[index]
		renderables = append(
			renderables,
			state.spriteFactory.create("person_selection", int(p.x), personRenderY, p.selectionAnimation.Frame()),
		)
	}
	if !state.data.isNoCheatSelected() {
		necessaryTargets := allCheats[state.data.cheats.availableCheats[state.data.cheats.selectedCheat].id].targets
		if len(necessaryTargets) > len(state.data.cheats.selectedCheatTargets) {
			nextTarget := necessaryTargets[len(state.data.cheats.selectedCheatTargets)]

			for i, p := range state.data.personQueue.persons {
				if nextTarget.isValidTarget(state.data.personQueue, i, state.data.cheats.selectedCheatTargets) {
					renderables = append(
						renderables,
						state.spriteFactory.create("person_marker", int(p.x), personRenderY, p.markerAnimation.Frame()),
					)
				}
			}
		}
	}
	return renderables
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
	p.markerAnimation = animation.NewFrames(3, 49)
	p.markerAnimation.Randomize()
	p.selectionAnimation = animation.NewFrames(3, 49)
	p.selectionAnimation.Randomize()

	state.data.personQueue.addPerson(p)
	state.data.personQueue.calculateDesiredX()
}
