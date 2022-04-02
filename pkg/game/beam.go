package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

func newBeam() *beam {
	return &beam{
		phase:         beamStates[0],
		nextBeamState: beamStateSwitchInterval,
		animation:     animation.NewFrames(3, 75),
	}
}

type beam struct {
	phase         string
	nextBeamState int

	animation animation.Frames
}

func (b *beam) tick(ms int) {
	b.animation.Tick(ms)

	b.nextBeamState -= ms
	if b.nextBeamState <= 0 {
		switch b.phase {
		case beamStates[0]:
			b.phase = beamStates[1]
			b.nextBeamState = beamStateSwitchInterval
		case beamStates[1]:
			b.phase = beamStates[2]
			b.nextBeamState = beamStateSwitchInterval
		case beamStates[2]:
			b.phase = ""
		}
	}
}

func (b *beam) isBeamed() bool {
	return b.phase == beamStates[2] || b.isOver()
}

func (b *beam) isOver() bool {
	return b.phase == ""
}

func (b *beam) rendered(sf *spriteFactory, x int, y int) canvas2drendering.Renderable {
	if b.phase == "" {
		return canvas2drendering.NopRenderable()
	}
	return sf.create("beam_"+b.phase, x, y, b.animation.Frame())
}

var beamStates = []string{
	"start",
	"middle",
	"end",
}

const beamStateSwitchInterval = 500
