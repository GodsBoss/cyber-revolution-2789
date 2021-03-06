package game

import (
	"math"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
)

type personQueue struct {
	persons []person

	personMostRightX float64
}

func (queue *personQueue) init() {
	queue.persons = make([]person, 0)
}

func (queue *personQueue) setMostRightX(x float64) {
	queue.personMostRightX = x
}

func (queue *personQueue) calculateDesiredX() {
	l := len(queue.persons)
	for i := range queue.persons {
		queue.persons[i].desiredX = float64(int(queue.personMostRightX) - personHorizontalDistance*(l-i-1))
	}
}

func (queue *personQueue) removeMostRightPerson() {
	queue.persons = queue.persons[:len(queue.persons)-1]
}

func (queue *personQueue) isPlayerAlive() bool {
	for _, person := range queue.persons {
		if person.Type == personTypePlayer {
			return true
		}
	}
	return false
}

func (queue *personQueue) Tick(ms int) {
	for i := range queue.persons {
		queue.persons[i].Tick(ms)
	}
}

func (queue *personQueue) addPerson(p person) {
	p.animation = animation.NewFrames(3, 200)
	p.animation.Randomize()
	p.markerAnimation = animation.NewFrames(3, 49)
	p.markerAnimation.Randomize()
	p.selectionAnimation = animation.NewFrames(3, 49)
	p.selectionAnimation.Randomize()

	queue.persons = append([]person{p}, queue.persons...)
	queue.calculateDesiredX()
}

func (queue personQueue) isAnyPersonMoving() bool {
	for _, p := range queue.persons {
		if p.isMoving() {
			return true
		}
	}
	return false
}

// swapPersons swaps the two persons at the indexes i and j. Will panic if these are not in range.
func (queue *personQueue) swapPersons(i, j int) {
	queue.persons[i], queue.persons[j] = queue.persons[j], queue.persons[i]
}

// Len returns the number of persons in the queue.
func (queue *personQueue) Len() int {
	return len(queue.persons)
}

type person struct {
	Type string

	x        float64
	desiredX float64

	animation          animation.Frames
	markerAnimation    animation.Frames
	selectionAnimation animation.Frames

	moving bool
}

func (p person) spriteID() string {
	return "person_" + p.Type
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
	p.animation.Tick(ms)
	p.selectionAnimation.Tick(ms)
	p.markerAnimation.Tick(ms)
	speed := personSpeed * (float64(ms) / 1000)
	if math.Abs(p.x-p.desiredX) <= speed {
		p.x = p.desiredX
		p.moving = false
		return
	}
	if p.desiredX < p.x {
		speed = -speed
	}
	p.x += speed
	p.moving = true
}

func (p person) isMoving() bool {
	return p.moving
}

const (
	personTypePlayer = "player"

	personTypeAlienGray    = "alien_gray"
	personTypeAlienBuddy   = "alien_buddy"
	personTypeAlienFerengi = "alien_ferengi"
	personTypeAlienRobot   = "alien_robot"

	personRenderY            = 80
	personHorizontalDistance = 30

	// personMostRightX is the x position the most right person desires.
	personMostRightX = 260

	// personSpeed is the speed of a person in pixel per second.
	personSpeed = 75
)

type personType struct {
	tags []string
}

var allPersonTypes = map[string]personType{
	personTypePlayer: {
		tags: []string{},
	},
	personTypeAlienGray: {
		tags: []string{
			tagAnosmic,
			tagWeak,
		},
	},
	personTypeAlienBuddy: {
		tags: []string{
			tagDumb,
		},
	},
	personTypeAlienFerengi: {
		tags: []string{
			tagGreedy,
		},
	},
	personTypeAlienRobot: {
		tags: []string{
			tagAnosmic,
			tagMechanical,
		},
	},
}

const (
	tagAnosmic    = "anosmic"
	tagDumb       = "dumb"
	tagGreedy     = "greedy"
	tagMechanical = "mechanical"
	tagWeak       = "weak"
)
