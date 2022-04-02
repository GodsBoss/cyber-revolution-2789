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
