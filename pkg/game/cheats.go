package game

import (
	"math/rand"

	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

func (state *statePlaying) addRandomCheat() {
	cheatIDs := make([]string, 0)
	for id := range allCheats {
		cheatIDs = append(cheatIDs, id)
	}
	id := cheatIDs[rand.Intn(len(cheatIDs))]
	newCheat := cheat{
		id: id,
	}
	state.cheats = append(state.cheats, newCheat)
}

func (state *statePlaying) unselectCheat() {
	state.selectedCheat = noCheatSelected
	state.selectedCheatTargets = nil
}

func (state *statePlaying) trySelectCheat(x int, y int) {
	for i := range state.cheats {
		cheatX, cheatY := state.cheatCoords(i)

		cheatBounds := rectangle{
			x:      cheatX,
			y:      cheatY,
			width:  cheatWidth,
			height: cheatHeight,
		}

		if cheatBounds.withinBounds(x, y) {
			state.selectedCheat = i
			return
		}
	}
}

func (state *statePlaying) tryActivateCheat(x int, y int) {
	cheatX, cheatY := state.cheatCoords(state.selectedCheat)

	cheatBounds := rectangle{
		x:      cheatX,
		y:      cheatY,
		width:  cheatWidth,
		height: cheatHeight,
	}

	if !cheatBounds.withinBounds(x, y) {
		return
	}

	allCheats[state.cheats[state.selectedCheat].id].invoke(state.personQueue, state.selectedCheatTargets)

	// Cheat has been used, remove it.
	state.cheats = append(state.cheats[0:state.selectedCheat], state.cheats[state.selectedCheat+1:]...)
	state.unselectCheat()

	// Person queue probably changed, recalculate.
	state.personQueue.calculateDesiredX()

	state.addRandomCheat()
}

const (
	noCheatSelected = -1
)

type cheat struct {
	id string
}

func (state *statePlaying) renderedCheats() canvas2drendering.Renderables {
	l := len(state.cheats)

	renderables := make(canvas2drendering.Renderables, l)

	for i, cheat := range state.cheats {
		x, y := state.cheatCoords(i)

		renderables[i] = state.spriteFactory.create(cheat.SpriteID(), x, y, 0)
	}

	return renderables
}

func (state *statePlaying) cheatCoords(index int) (x int, y int) {
	l := len(state.cheats)

	x = cheatCenterX + cheatWidth*index - (cheatWidth*l)/2
	y = cheatRenderY

	if index == state.selectedCheat {
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

func (state *statePlaying) trySelectTarget(x int, y int) {
	ch := allCheats[state.cheats[state.selectedCheat].id]
	targetCandidates := make([]int, 0)

	for index := range state.personQueue.persons {
		// Person wasn't clicked.
		if !state.personQueue.persons[index].bounds().withinBounds(x, y) {
			continue
		}

		// Person isn't a valid target.
		if !ch.targets[len(state.selectedCheatTargets)].isValidTarget(state.personQueue, index, state.selectedCheatTargets) {
			continue
		}

		targetCandidates = append(targetCandidates, index)
	}

	// No candidate or too many.
	if len(targetCandidates) != 1 {
		return
	}

	state.selectedCheatTargets = append(state.selectedCheatTargets, targetCandidates[0])
}

type cheatTarget interface {
	isValidTarget(queue *personQueue, index int, currentTargets []int) bool
}

// cheatTargetAny accepts any person as a target.
type cheatTargetAny struct{}

func (target cheatTargetAny) isValidTarget(_ *personQueue, _ int, _ []int) bool {
	return true
}

// cheatTargetNotTargeted accepts any person as a target that has not been targeted as a target so far.
type cheatTargetNotTargeted struct{}

func (target cheatTargetNotTargeted) isValidTarget(_ *personQueue, index int, currentTargets []int) bool {
	for _, currentTargetIndex := range currentTargets {
		if currentTargetIndex == index {
			return false
		}
	}
	return true
}

var allCheats = map[string]cheatAction{
	cheatIDLeftMost: {
		invoke: func(queue *personQueue, _ []int) {
			playerIndex := 0
			for i := range queue.persons {
				if queue.persons[i].Type == personTypePlayer {
					playerIndex = i
				}
			}
			for playerIndex > 0 {
				queue.persons[playerIndex], queue.persons[playerIndex-1] = queue.persons[playerIndex-1], queue.persons[playerIndex]
				playerIndex--
			}
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
			queue.persons[targets[0]], queue.persons[targets[1]] = queue.persons[targets[1]], queue.persons[targets[0]]
		},
	},
}

const (
	cheatIDLeftMost = "leftmost"
	cheatIDSwap     = "swap"
)
