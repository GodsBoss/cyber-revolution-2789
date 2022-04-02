package game

type playingData struct {
	personQueue personQueue
	cheats      cheats
}

// init initializes the data with everything being empty.
func (data *playingData) init() {
	data.personQueue = personQueue{
		persons: make([]person, 0),
	}
	data.cheats = cheats{
		availableCheats:      make([]cheat, 0),
		selectedCheat:        noCheatSelected,
		selectedCheatTargets: nil,
	}
}

func (data *playingData) unselectCheat() {
	data.cheats.unselectCheat()
}
