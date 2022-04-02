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
