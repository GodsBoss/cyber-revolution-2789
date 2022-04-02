package animation

import "math/rand"

type Frames interface {
	Tick(ms int)
	Frame() int
	Randomize()
}

func NewFrames(maxFrame int, msPerFrame int) Frames {
	return &animation{
		maxFrame:   maxFrame,
		msPerFrame: msPerFrame,
	}
}

type animation struct {
	maxFrame   int
	msPerFrame int

	current int
}

func (anim *animation) Tick(ms int) {
	if anim.maxFrame == 0 {
		return
	}
	anim.current += ms
	if anim.Frame() > anim.maxFrame {
		anim.current -= anim.Frame() * anim.msPerFrame
	}
}

func (anim *animation) Frame() int {
	if anim.maxFrame == 0 {
		return 0
	}
	return anim.current / anim.msPerFrame
}

func (anim *animation) Randomize() {
	if anim.maxFrame == 0 {
		return
	}
	anim.current = rand.Intn(anim.maxFrame * anim.msPerFrame)
}
