package game

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
