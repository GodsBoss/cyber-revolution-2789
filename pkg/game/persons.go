package game

import (
	"math"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
)

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
	p.markerAnimation = animation.NewFrames(3, 49)
	p.markerAnimation.Randomize()
	p.selectionAnimation = animation.NewFrames(3, 49)
	p.selectionAnimation.Randomize()

	queue.persons = append([]person{p}, queue.persons...)
	queue.calculateDesiredX()
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

type personType struct {
	tags []string
}

var allPersonTypes = map[string]personType{
	personTypeGreenAlien: {},
	personTypePlayer:     {},
}
