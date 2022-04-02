package game

import (
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

const statePlayingID = "playing"

type statePlaying struct {
	spriteFactory *spriteFactory

	data *playingData
}

func (state *statePlaying) init() {}

func (state *statePlaying) tick(ms int) (next string) {
	state.data.tick(ms)

	return ""
}

func (state *statePlaying) receiveKeyEvent(event interaction.KeyEvent) (next string) {
	if event.Type == interaction.KeyUp && event.Key == "Escape" {
		state.unselectCheat()
	}
	return ""
}

func (state *statePlaying) receiveMouseEvent(event interaction.MouseEvent) (next string) {
	if event.Type == interaction.MouseUp {
		// No cheat selected yet, so try to select one.
		if state.data.isNoCheatSelected() {
			state.data.trySelectCheat(event.X, event.Y)
			return ""
		}

		// All cheat targets are selected, try to activate cheat.
		if state.data.areAllTargetsSelected() {
			state.tryActivateCheat(event.X, event.Y)
			return ""
		}

		// Try to select target.
		state.trySelectTarget(event.X, event.Y)
	}

	return ""
}

func (state *statePlaying) renderable() canvas2drendering.Renderable {
	renderables := canvas2drendering.Renderables{
		state.spriteFactory.create("background", 0, 0, 0),
	}
	renderables = append(renderables, state.data.rendered(state.spriteFactory)...)

	return renderables
}

func (state *statePlaying) unselectCheat() {
	state.data.unselectCheat()
}

func (state *statePlaying) tryActivateCheat(x int, y int) {
	state.data.tryActivateCheat(x, y)
}
