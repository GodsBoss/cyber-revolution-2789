package game

import (
	"math"
	"math/rand"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

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

type personQueue struct {
	persons []person
}

func (queue *personQueue) init() {
	queue.persons = make([]person, 0)
}

func (queue *personQueue) calculateDesiredX() {
	l := len(queue.persons)
	for i := range queue.persons {
		queue.persons[i].desiredX = float64(personMostRightX - personHorizontalDistance*(l-i-1))
	}
}

func (queue *personQueue) Tick(ms int) {
	for i := range queue.persons {
		queue.persons[i].Tick(ms)
	}
}

func (queue *personQueue) addPerson(p person) {
	queue.persons = append([]person{p}, queue.persons...)
}

type person struct {
	Type string

	x        float64
	desiredX float64

	markerAnimation    animation.Frames
	selectionAnimation animation.Frames
}

func (p person) bounds() rectangle {
	return rectangle{
		x:      int(math.Floor(p.x)),
		y:      personRenderY,
		width:  32,
		height: 48,
	}
}

func (p *person) Tick(ms int) {
	p.selectionAnimation.Tick(ms)
	p.markerAnimation.Tick(ms)
	speed := personSpeed * (float64(ms) / 1000)
	if math.Abs(p.x-p.desiredX) <= speed {
		p.x = p.desiredX
		return
	}
	if p.desiredX < p.x {
		speed = -speed
	}
	p.x += speed
}

const (
	personTypePlayer     = "player"
	personTypeGreenAlien = "green_alien"

	personRenderY            = 80
	personHorizontalDistance = 40

	// personMostRightX is the x position the most right person desires.
	personMostRightX = 260

	// personSpeed is the speed of a person in pixel per second.
	personSpeed = 25
)

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

type personType struct {
	tags []string
}

var allPersonTypes = map[string]personType{
	personTypeGreenAlien: {},
	personTypePlayer:     {},
}
