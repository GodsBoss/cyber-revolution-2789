package game

import (
	"math/rand"

	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

type playingData struct {
	personQueue personQueue
	cheats      cheats
}

// init initializes the data with everything being empty.
func (data *playingData) init() {
	data.personQueue.init()
	data.cheats.init()
}

func (data *playingData) unselectCheat() {
	data.cheats.unselectCheat()
}

func (data *playingData) tick(ms int) {
	data.personQueue.Tick(ms)
	data.cheats.tick(ms)
}

func (data *playingData) isNoCheatSelected() bool {
	return data.cheats.isNoCheatSelected()
}

func (data *playingData) cheatCoords(index int) (x int, y int) {
	return data.cheats.cheatCoords(index)
}

func (data *playingData) areAllTargetsSelected() bool {
	return data.cheats.areAllTargetsSelected()
}

func (data *playingData) trySelectCheat(x int, y int) {
	data.cheats.trySelectCheat(x, y)
}

func (data *playingData) addRandomCheat() {
	data.cheats.addRandomCheat()
}

func (data *playingData) tryActivateCheat(x int, y int) {
	cheatX, cheatY := data.cheatCoords(data.cheats.selectedCheat)

	cheatBounds := rectangle{
		x:      cheatX,
		y:      cheatY,
		width:  cheatWidth,
		height: cheatHeight,
	}

	if !cheatBounds.withinBounds(x, y) {
		return
	}

	allCheats[data.cheats.availableCheats[data.cheats.selectedCheat].id].invoke(&data.personQueue, data.cheats.selectedCheatTargets)

	// Cheat has been used, remove it.
	data.cheats.availableCheats = append(
		data.cheats.availableCheats[0:data.cheats.selectedCheat],
		data.cheats.availableCheats[data.cheats.selectedCheat+1:]...,
	)
	data.unselectCheat()

	// Person queue probably changed, recalculate.
	data.personQueue.calculateDesiredX()

	data.addRandomCheat()
}

func (data *playingData) addRandomPerson(x float64) {
	ids := make([]string, 0)
	for id := range allPersonTypes {
		if id != personTypePlayer {
			ids = append(ids, id)
		}
	}
	typ := ids[rand.Intn(len(ids))]

	data.addPerson(
		person{
			Type: typ,
			x:    x,
		},
	)
}

func (data *playingData) addPlayer(x float64) {
	data.addPerson(
		person{
			Type: personTypePlayer,
			x:    x,
		},
	)
}

func (data *playingData) addPerson(p person) {
	data.personQueue.addPerson(p)
}

func (data *playingData) rendered(sf *spriteFactory) canvas2drendering.Renderables {
	renderables := data.renderedPersons(sf)
	renderables = append(renderables, data.renderedCheats(sf)...)
	return renderables
}

func (data *playingData) renderedPersons(sf *spriteFactory) canvas2drendering.Renderables {
	renderables := make(canvas2drendering.Renderables, len(data.personQueue.persons))
	for i, person := range data.personQueue.persons {
		renderables[i] = sf.create(person.spriteID(), int(person.x), personRenderY, 0)
	}
	for _, index := range data.cheats.selectedCheatTargets {
		p := data.personQueue.persons[index]
		renderables = append(
			renderables,
			sf.create("person_selection", int(p.x), personRenderY, p.selectionAnimation.Frame()),
		)
	}
	if !data.isNoCheatSelected() {
		necessaryTargets := allCheats[data.cheats.availableCheats[data.cheats.selectedCheat].id].targets
		if len(necessaryTargets) > len(data.cheats.selectedCheatTargets) {
			nextTarget := necessaryTargets[len(data.cheats.selectedCheatTargets)]

			for i, p := range data.personQueue.persons {
				if nextTarget.isValidTarget(data.personQueue, i, data.cheats.selectedCheatTargets) {
					renderables = append(
						renderables,
						sf.create("person_marker", int(p.x), personRenderY, p.markerAnimation.Frame()),
					)
				}
			}
		}
	}
	return renderables
}

func (data *playingData) renderedCheats(sf *spriteFactory) canvas2drendering.Renderables {
	l := len(data.cheats.availableCheats)

	renderables := make(canvas2drendering.Renderables, l)

	for i, cheat := range data.cheats.availableCheats {
		x, y := data.cheatCoords(i)

		renderables[i] = sf.create(cheat.SpriteID(), x, y, 0)

		// If no cheat is selected, highlight all cheats as possible user interactions.
		if data.isNoCheatSelected() {
			renderables = append(renderables, sf.create("cheat_marker", x-3, y-3, cheat.markerAnimation.Frame()))
		}
	}

	if data.areAllTargetsSelected() {
		x, y := data.cheatCoords(data.cheats.selectedCheat)
		renderables = append(
			renderables,
			sf.create("cheat_marker", x-3, y-3, data.cheats.availableCheats[data.cheats.selectedCheat].markerAnimation.Frame()),
		)
	}

	return renderables
}
