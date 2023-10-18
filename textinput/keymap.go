package textinput

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

var _ help.KeyMap = (*KeyMap)(nil)

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Confirm: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "confirm")),
	}
}

type KeyMap struct {
	Confirm key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Confirm}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
