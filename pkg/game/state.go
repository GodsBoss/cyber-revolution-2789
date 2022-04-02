package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/interaction"
	"github.com/GodsBoss/gggg/pkg/rendering/canvas2drendering"
)

type state interface {
	// init is called whenever the game switches to this state.
	init()

	tick(ms int) (next string)

	receiveKeyEvent(event interaction.KeyEvent) (next string)

	receiveMouseEvent(event interaction.MouseEvent) (next string)

	renderable() canvas2drendering.Renderable
}

type states struct {
	states         map[string]state
	currentStateID string
}

func (st *states) currentState() state {
	return st.states[st.currentStateID]
}

// setNextState sets the next state to be identified by id. Does nothing if passed an empty string, panics if id is unknown.
func (st *states) setNextState(id string) {
	if id == "" {
		return
	}
	_, ok := st.states[id]
	if !ok {
		panic("invalid state id " + id)
	}
	st.currentStateID = id
	st.currentState().init()
}

func (st *states) receiveKeyEvent(event interaction.KeyEvent) {
	st.setNextState(st.currentState().receiveKeyEvent(event))
}

func (st *states) receiveMouseEvent(event interaction.MouseEvent) {
	st.setNextState(st.currentState().receiveMouseEvent(event))
}

func (st *states) tick(ms int) {
	st.setNextState(st.currentState().tick(ms))
}

func (st *states) render(output *dom.Context2D) {
	st.currentState().renderable().Render(output)
}
