package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
)

type Option func(*State)

func WithPrompt(prompt string) Option {
	return func(state *State) {
		state.prompt = prompt
	}
}

func WithOnConfirm(onConfirm tea.Cmd) Option {
	return func(state *State) {
		state.onConfirm = onConfirm
	}
}

func WithOnCancel(onCancel tea.Cmd) Option {
	return func(state *State) {
		state.onCancel = onCancel
	}
}

func New(options ...Option) *State {
	state := &State{
		prompt: "Are you sure?",
		keyMap: DefaultKeyMap(),
		size: soda.Size{
			Width:  80,
			Height: 40,
		},
	}

	for _, option := range options {
		option(state)
	}

	return state
}
