package textinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/soda/title"
)

type Option func(*State)

func WithTitle(t title.Title) Option {
	return func(state *State) {
		state.title = t
	}
}

func WithKeyMap(keyMap KeyMap) Option {
	return func(state *State) {
		state.keyMap = keyMap
	}
}

func WithSubtitle(subtitle string) Option {
	return func(state *State) {
		state.subtitle = subtitle
	}
}

func WithOnConfirm(onConfirm func(string) tea.Cmd) Option {
	return func(state *State) {
		state.onConfirm = onConfirm
	}
}

func WithValidate(validate textinput.ValidateFunc) Option {
	return func(state *State) {
		state.textInput.Validate = validate
	}
}

func WithValue(value string) Option {
	return func(state *State) {
		state.textInput.SetValue(value)
	}
}

func WithPlaceholder(placeholder string) Option {
	return func(state *State) {
		state.textInput.Placeholder = placeholder
	}
}

func New(options ...Option) *State {
	state := &State{
		title:     title.New("Input"),
		textInput: textinput.New(),
		keyMap:    DefaultKeyMap(),
		onConfirm: func(s string) tea.Cmd {
			return nil
		},
	}

	state.textInput.PromptStyle = lipgloss.NewStyle().Bold(true)

	for _, option := range options {
		option(state)
	}

	return state
}
