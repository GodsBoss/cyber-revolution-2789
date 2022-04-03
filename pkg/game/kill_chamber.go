package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

type killChamber struct {
	normalAnimation animation.Frames
	activeAnimation animation.Frames
}

func newKillChamber() *killChamber {
	return &killChamber{
		normalAnimation: animation.NewFrames(5, 300),
		activeAnimation: animation.NewFrames(5, 50),
	}
}

func (chamber *killChamber) tick(ms int) {
	chamber.normalAnimation.Tick(ms)
	chamber.activeAnimation.Tick(ms)
}

func (chamber *killChamber) render(sf *spriteFactory, active bool) canvas2drendering.Renderable {
	if active {
		return sf.create("kill_chamber_ground_active", killChamberBottomX, killChamberBottomY, chamber.activeAnimation.Frame())
	}

	return sf.create("kill_chamber_ground", killChamberBottomX, killChamberBottomY, chamber.normalAnimation.Frame())
}

const (
	killChamberBottomX = 286
	killChamberBottomY = 123
)
