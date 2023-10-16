package loading

import "github.com/charmbracelet/lipgloss"

type StyleMap struct {
	Spinner lipgloss.Style
	Message lipgloss.Style
}

func DefaultStyleMap() StyleMap {
	return StyleMap{
		Spinner: lipgloss.NewStyle().Bold(true),
		Message: lipgloss.NewStyle(),
	}
}
