package loading

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	title    title.Title
	subtitle string

	spinner spinner.Model

	backable func() bool

	message  string
	messages <-chan string

	size soda.Size

	styleMap StyleMap

	keyMap keyMap
}

func (s *State) Destroy() {
}

func (s *State) Backable() bool {
	return s.backable()
}

func (s *State) Resize(size soda.Size) tea.Cmd {
	s.size = size
	return nil
}

func (s *State) Title() title.Title {
	return s.title
}

func (s *State) Subtitle() string {
	return s.subtitle
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(mh soda.ModelHandler) tea.Cmd {
	return tea.Batch(s.receive, s.spinner.Tick)
}

func (s *State) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		s.spinner, cmd = s.spinner.Update(msg)

		return cmd
	case messageMsg:
		s.message = msg.Message
		return s.receive
	}

	return nil
}

func (s *State) viewMessage() string {
	return s.styleMap.Message.Render(s.message)
}

func (s *State) View(mh soda.ModelHandler) string {
	// make it more centered visually
	height := s.size.Height - 2
	if s.subtitle != "" {
		height -= 2
	}

	width := s.size.Width

	style := lipgloss.NewStyle().Width(width).Height(height / 2).AlignHorizontal(lipgloss.Center)
	spinnerStyle := style.Copy().AlignVertical(lipgloss.Bottom)
	messageStyle := s.styleMap.Message.Inherit(style.Copy().AlignVertical(lipgloss.Top))

	spinner := spinnerStyle.Render(s.spinner.View())
	message := messageStyle.Render(s.message)

	return lipgloss.JoinVertical(lipgloss.Center, spinner, message)
}
