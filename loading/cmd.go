package loading

import tea "github.com/charmbracelet/bubbletea"

func (s *State) receive() tea.Msg {
	return messageMsg{Message: <-s.messages}
}
