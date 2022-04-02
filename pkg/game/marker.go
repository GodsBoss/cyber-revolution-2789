package game

import (
	"math/rand"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
)

// marker highlights possible user interactions. It is basically a bunch of particles around a rectangle.
//
// Unfinished and unused.
type marker struct {
	// object is the marked object. Marking particles fly around it.
	object markedObject

	// width is the marker border width.
	width float64

	particles []markerParticle
}

func (m *marker) Tick(ms int) {
	for i := range m.particles {
		m.particles[i].move(m, ms)
	}
}

func newMarker(object markedObject, width float64) *marker {
	particlesCount := int((object.width() + object.height()) * 2 * markerDensity)

	marker := &marker{
		object:    object,
		width:     width,
		particles: make([]markerParticle, particlesCount),
	}

	for i := range marker.particles {
		marker.particles[i].animation = animation.NewFrames(3, markerParticleMsPerFrame)
		marker.particles[i].animation.Randomize()

		if rand.Intn(2) == 0 {
			marker.particles[i].turnDirection = counterClockwise
		}

		linePosition := rand.Float64() * (object.width() + object.height()) * 2

		switch {
		// Left border
		case linePosition > object.width()*2+object.height():
			linePosition -= object.width()*2 + object.height()
			marker.particles[i].x = object.x() - marker.width
			marker.particles[i].y = object.y() + linePosition
			if marker.particles[i].turnDirection == clockwise {
				marker.particles[i].dy = -markerParticleSpeed
			} else {
				marker.particles[i].dy = markerParticleSpeed
			}

		// Bottom border
		case linePosition <= object.width()*2+object.height() && linePosition > object.width()+object.height():
			linePosition -= object.width() + object.height()
			marker.particles[i].x = object.x() + linePosition
			marker.particles[i].y = object.y() + object.height() + marker.width
			if marker.particles[i].turnDirection == clockwise {
				marker.particles[i].dx = -markerParticleSpeed
			} else {
				marker.particles[i].dx = markerParticleSpeed
			}

		// Right border
		case linePosition <= object.width()+object.height() && linePosition > object.width():
			linePosition -= object.width()
			marker.particles[i].x = object.x() + marker.width
			marker.particles[i].y = object.y() + linePosition
			if marker.particles[i].turnDirection == clockwise {
				marker.particles[i].dy = markerParticleSpeed
			} else {
				marker.particles[i].dy = -markerParticleSpeed
			}

		// Top border
		case linePosition <= object.width():
			marker.particles[i].x = object.x() + linePosition
			marker.particles[i].y = object.y() - marker.width
			if marker.particles[i].turnDirection == clockwise {
				marker.particles[i].dx = markerParticleSpeed
			} else {
				marker.particles[i].dx = -markerParticleSpeed
			}
		}
	}

	return marker
}

// markerDensity is the marker density per pixel length of the object.
const markerDensity = 0.1

type markedObject interface {
	x() float64
	y() float64
	width() float64
	height() float64
}

type markerParticle struct {
	turnDirection turnDirection

	x float64
	y float64

	// dx is the horizontal speed in pixels per second.
	dx float64

	// dy is the vertical speed in pixels per second.
	dy float64

	animation animation.Frames
}

func (particle *markerParticle) move(m *marker, ms int) {
	particle.animation.Tick(ms)

	particle.x += particle.dx
	particle.y += particle.dy
}

type turnDirection int

const (
	clockwise turnDirection = iota
	counterClockwise
)

type direction int

const (
	up direction = iota
	left
	down
	right
)

// markerParticleSpeed is the speed of marker particles in pixels per second.
const markerParticleSpeed = 5.0

// markerParticleMsPerFrame is the animation slowness of particles.
const markerParticleMsPerFrame = 50
