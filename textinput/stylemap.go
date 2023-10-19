package textinput

import "github.com/charmbracelet/lipgloss"

type StyleMap struct {
	Error lipgloss.Style
}

func DefaultStyleMap() StyleMap {
	return StyleMap{
		Error: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{
			Light: "#9d0006",
			Dark:  "#fb4934",
		}),
	}
}
