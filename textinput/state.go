package textinput

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	title     title.Title
	subtitle  string
	textInput textinput.Model

	onConfirm func(string) tea.Cmd

	keyMap KeyMap
}

func (s *State) Destroy() {
}

func (s *State) Backable() bool {
	return true
}

func (s *State) Resize(size soda.Size) tea.Cmd {
	s.textInput.Width = size.Width
	return nil
}

func (s *State) Title() title.Title {
	return s.title
}

func (s *State) Subtitle() string {
	return s.subtitle
}

func (s *State) Status() string {
	if err := s.textInput.Err; err != nil {
		// TODO: red
		return err.Error()
	}

	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(mh soda.ModelHandler) tea.Cmd {
	return tea.Batch(textinput.Blink, s.textInput.Focus())
}

func (s *State) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Confirm) && s.textInput.Err == nil:
			return s.onConfirm(s.textInput.Value())
		}
	}

	var cmd tea.Cmd
	s.textInput, cmd = s.textInput.Update(msg)
	return cmd
}

func (s *State) View(mh soda.ModelHandler) string {
	return s.textInput.View()
}
