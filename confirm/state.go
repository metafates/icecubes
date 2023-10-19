package confirm

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	prompt    string
	onConfirm tea.Cmd
	onCancel  tea.Cmd
	size      soda.Size

	keyMap KeyMap
}

func (s *State) Destroy() {
}

func (s *State) Backable() bool {
	return true
}

func (s *State) Resize(size soda.Size) tea.Cmd {
	s.size = size
	return nil
}

func (s *State) Title() title.Title {
	return title.New("Confirm")
}

func (s *State) Subtitle() string {
	return ""
}

func (s *State) Status() string {
	return ""
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(mh soda.ModelHandler) tea.Cmd {
	return nil
}

func (s *State) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Confirm):
			return s.onConfirm
		case key.Matches(msg, s.keyMap.Cancel):
			return s.onCancel
		}
	}

	return nil
}

func (s *State) View(mh soda.ModelHandler) string {
	// make it more centered visually
	height := s.size.Height

	width := s.size.Width

	style := lipgloss.NewStyle().Width(width).Height(height).Align(lipgloss.Center, lipgloss.Center)

	return style.Render(s.prompt)
}
