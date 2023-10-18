package loading

import (
	"github.com/charmbracelet/bubbles/spinner"
)

type Option func(*State)

func WithMessages(messages <-chan string) Option {
	return func(state *State) {
		state.messages = messages
	}
}

func WithMessage(message string) Option {
	return func(state *State) {
		state.message = message
	}
}

func WithBackable(backable func() bool) Option {
	return func(state *State) {
		state.backable = backable
	}
}

func WithSpinner(s spinner.Spinner) Option {
	return func(state *State) {
		state.spinner.Spinner = s
	}
}

func WithStyleMap(styleMap StyleMap) Option {
	return func(state *State) {
		state.styleMap = styleMap
	}
}

func New(options ...Option) *State {
	state := &State{
		backable: func() bool {
			return true
		},
		message:  "Loading",
		messages: make(chan string),
		spinner:  spinner.New(spinner.WithSpinner(spinner.Ellipsis)),
		styleMap: DefaultStyleMap(),
		keyMap:   keyMap{},
	}

	for _, option := range options {
		option(state)
	}

	state.spinner.Style = state.styleMap.Spinner

	return state
}
