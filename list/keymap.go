package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

var _ help.KeyMap = (*keyMap)(nil)

type keyMap struct {
	list  list.KeyMap
	extra help.KeyMap
}

func (k keyMap) ShortHelp() []key.Binding {
	keys := []key.Binding{
		k.list.CursorUp,
		k.list.CursorDown,
	}

	if k.extra == nil {
		return keys
	}

	return append(keys, k.extra.ShortHelp()...)
}

func (k keyMap) FullHelp() [][]key.Binding {
	keys := [][]key.Binding{
		{
			k.list.Filter,
			k.list.NextPage,
			k.list.PrevPage,
			k.list.GoToStart,
			k.list.GoToEnd,
		},
	}

	if k.extra == nil {
		return keys
	}

	return append(keys, k.extra.FullHelp()...)
}
