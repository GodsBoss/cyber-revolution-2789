package game

import (
	"math/rand"

	"github.com/GodsBoss/cyber-revolution-2789/pkg/animation"
)

type cheats struct {
	// availableCheats are the cheats the player has currently available.
	availableCheats []cheat

	// selectedCheat is the index of the chosen cheat. Contains -1 if no cheat is selected.
	selectedCheat int

	// selectedCheatTargets are the indexes of the chosen cheat targets (persons).
	selectedCheatTargets []int
}

func (chs *cheats) init() {
	chs.availableCheats = make([]cheat, 0)
	chs.selectedCheat = noCheatSelected
	chs.selectedCheatTargets = nil
}

func (chs *cheats) unselectCheat() {
	chs.selectedCheat = noCheatSelected
	chs.selectedCheatTargets = nil
}

func (chs *cheats) tick(ms int) {
	for i := range chs.availableCheats {
		chs.availableCheats[i].markerAnimation.Tick(ms)
	}
}

func (chs *cheats) isNoCheatSelected() bool {
	return chs.selectedCheat == noCheatSelected
}

// areAllTargetsSelected returns true if a cheat is selected and all its targets, too.
func (chs *cheats) areAllTargetsSelected() bool {
	if chs.isNoCheatSelected() {
		return false
	}

	return len(allCheats[chs.availableCheats[chs.selectedCheat].id].targets) == len(chs.selectedCheatTargets)
}

func (chs *cheats) trySelectCheat(x int, y int) {
	for i := range chs.availableCheats {
		cheatX, cheatY := chs.cheatCoords(i)

		cheatBounds := rectangle{
			x:      cheatX,
			y:      cheatY,
			width:  cheatWidth,
			height: cheatHeight,
		}

		if cheatBounds.withinBounds(x, y) {
			chs.selectedCheat = i
			return
		}
	}
}

func (chs *cheats) addRandomCheat() {
	cheatIDs := make([]string, 0)
	for id := range allCheats {
		cheatIDs = append(cheatIDs, id)
	}
	id := cheatIDs[rand.Intn(len(cheatIDs))]
	newCheat := cheat{
		id:              id,
		markerAnimation: animation.NewFrames(3, 80),
	}
	newCheat.markerAnimation.Randomize()
	chs.availableCheats = append(chs.availableCheats, newCheat)
}

const (
	noCheatSelected = -1
)

type cheat struct {
	id string

	markerAnimation animation.Frames
}

func (chs *cheats) cheatCoords(index int) (x int, y int) {
	l := len(chs.availableCheats)

	x = cheatCenterX + cheatWidth*index - (cheatWidth*l)/2
	y = cheatRenderY

	if index == chs.selectedCheat {
		y += cheatRenderYOffset
	}

	return x, y
}

func (ch cheat) SpriteID() string {
	return "cheat_" + ch.id
}

const (
	// cheatRenderY is the y position of unselected cheats.
	cheatRenderY = 160

	// cheatRenderYOffset is the vertical offset of the selected cheat.
	cheatRenderYOffset = -10

	cheatWidth  = 24
	cheatHeight = 24

	cheatMargin = 8

	cheatCenterX = 160
)

type cheatAction struct {
	targets []cheatTarget

	// invoke is called after target selection.
	invoke func(queue *personQueue, targets []int)
}

func (state *statePlayingInteraction) trySelectTarget(x int, y int) {
	ch := allCheats[state.data.cheats.availableCheats[state.data.cheats.selectedCheat].id]
	targetCandidates := make([]int, 0)

	for index := range state.data.personQueue.persons {
		// Person wasn't clicked.
		if !state.data.personQueue.persons[index].bounds().withinBounds(x, y) {
			continue
		}

		// Person isn't a valid target.
		if !ch.targets[len(state.data.cheats.selectedCheatTargets)].isValidTarget(state.data.personQueue, index, state.data.cheats.selectedCheatTargets) {
			continue
		}

		targetCandidates = append(targetCandidates, index)
	}

	// No candidate or too many.
	if len(targetCandidates) != 1 {
		return
	}

	state.data.cheats.selectedCheatTargets = append(state.data.cheats.selectedCheatTargets, targetCandidates[0])
}

type cheatTarget interface {
	isValidTarget(queue personQueue, index int, currentTargets []int) bool
}

// cheatTargetAny accepts any person as a target.
type cheatTargetAny struct{}

func (target cheatTargetAny) isValidTarget(_ personQueue, _ int, _ []int) bool {
	return true
}

// cheatTargetNotTargeted accepts any person as a target that has not been targeted as a target so far.
type cheatTargetNotTargeted struct{}

func (target cheatTargetNotTargeted) isValidTarget(_ personQueue, index int, currentTargets []int) bool {
	for _, currentTargetIndex := range currentTargets {
		if currentTargetIndex == index {
			return false
		}
	}
	return true
}

type cheatTargetHasTag string

func (target cheatTargetHasTag) isValidTarget(queue personQueue, index int, currentTargets []int) bool {
	targetTags := allPersonTypes[queue.persons[index].Type].tags
	for _, tag := range targetTags {
		if tag == string(target) {
			return true
		}
	}
	return false
}

type cheatTargetNot struct {
	target cheatTarget
}

func (target cheatTargetNot) isValidTarget(queue personQueue, index int, currentTargets []int) bool {
	return !target.target.isValidTarget(queue, index, currentTargets)
}

type cheatTargetAnd []cheatTarget

func (target cheatTargetAnd) isValidTarget(queue personQueue, index int, currentTargets []int) bool {
	for _, t := range target {
		if !t.isValidTarget(queue, index, currentTargets) {
			return false
		}
	}
	return true
}

type cheatTargetOr []cheatTarget

func (target cheatTargetOr) isValidTarget(queue personQueue, index int, currentTargets []int) bool {
	for _, t := range target {
		if t.isValidTarget(queue, index, currentTargets) {
			return true
		}
	}
	return false
}

// cheatTargetNotFirstInQueue is a cheat target that applies if a person is not first one of the queue, i.e. the next victim.
type cheatTargetNotFirstInQueue struct{}

func (target cheatTargetNotFirstInQueue) isValidTarget(queue personQueue, index int, _ []int) bool {
	return queue.Len()-1 > index
}

// cheatTargetOther is a cheat target that applies a cheat target to a different target than the one currently tested, defined by
// the offset. If no such target exists, this target is invalid.
type cheatTargetOther struct {
	offset int
	target cheatTarget
}

func (target cheatTargetOther) isValidTarget(queue personQueue, index int, currentTargets []int) bool {
	otherIndex := index + target.offset

	// If that other target does not exist, fail early.
	if otherIndex < 0 || otherIndex >= queue.Len() {
		return false
	}

	return target.target.isValidTarget(queue, otherIndex, currentTargets)
}

var allCheats = map[string]cheatAction{
	cheatIDBombThread: {
		invoke: func(queue *personQueue, _ []int) {
			rand.Shuffle(len(queue.persons), queue.swapPersons)

			// We don't want the player to die here, so move them to another position.
			lastIndex := queue.Len() - 1
			if queue.persons[lastIndex].Type == personTypePlayer {
				queue.swapPersons(lastIndex, rand.Intn(lastIndex))
			}
		},
	},
	cheatIDBribe: {
		targets: []cheatTarget{
			cheatTargetAnd{
				cheatTargetHasTag(tagGreedy),
				cheatTargetNotFirstInQueue{},
			},
		},
		invoke: func(queue *personQueue, targets []int) {
			queue.swapPersons(targets[0], targets[0]+1)
		},
	},
	cheatIDFart: {
		targets: []cheatTarget{
			cheatTargetAnd{
				cheatTargetNot{
					target: cheatTargetHasTag(tagMechanical),
				},
				cheatTargetOther{
					target: cheatTargetNot{
						target: cheatTargetHasTag(tagAnosmic),
					},
					offset: -1,
				},
			},
		},
		invoke: func(queue *personQueue, targets []int) {
			queue.swapPersons(targets[0], targets[0]-1)
		},
	},
	cheatIDLeftMost: {
		invoke: func(queue *personQueue, _ []int) {
			playerIndex := 0
			for i := range queue.persons {
				if queue.persons[i].Type == personTypePlayer {
					playerIndex = i
				}
			}
			for playerIndex > 0 {
				queue.swapPersons(playerIndex, playerIndex-1)
				playerIndex--
			}
		},
	},
	cheatIDPersuade: {
		targets: []cheatTarget{
			cheatTargetAnd{
				cheatTargetNot{
					cheatTargetHasTag(tagDumb),
				},
				cheatTargetOther{
					target: cheatTargetHasTag(tagDumb),
					offset: -1,
				},
			},
		},
		invoke: func(queue *personQueue, targets []int) {
			queue.swapPersons(targets[0], targets[0]-1)
		},
	},
	cheatIDShove: {
		targets: []cheatTarget{
			cheatTargetAnd{
				cheatTargetNot{
					target: cheatTargetHasTag(tagWeak),
				},
				cheatTargetOther{
					target: cheatTargetHasTag(tagWeak),
					offset: -1,
				},
			},
		},
		invoke: func(queue *personQueue, targets []int) {
			queue.swapPersons(targets[0], targets[0]-1)
		},
	},
	cheatIDSwap: {
		targets: []cheatTarget{
			cheatTargetAny{},
			cheatTargetNotTargeted{},
		},
		invoke: func(queue *personQueue, targets []int) {
			if len(targets) < 2 {
				return
			}
			queue.swapPersons(targets[0], targets[1])
		},
	},
}

const (
	cheatIDBombThread = "bomb_threat"
	cheatIDBribe      = "bribe"
	cheatIDFart       = "fart"
	cheatIDLeftMost   = "leftmost"
	cheatIDPersuade   = "persuade"
	cheatIDShove      = "shove"
	cheatIDSwap       = "swap"
)
