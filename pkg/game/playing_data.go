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
