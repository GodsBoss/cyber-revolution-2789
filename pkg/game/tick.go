package game

func (g *game) TicksPerSecond() int {
	return ticksPerSecond
}

const ticksPerSecond = 50

func (g *game) Tick(ms int) {}
