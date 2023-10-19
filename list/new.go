package list

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
)

type Option func(*State)

func WithKeyMap(keyMap help.KeyMap) Option {
	return func(state *State) {
		state.keyMap.extra = keyMap
	}
}

func NewWithList(list list.Model, options ...Option) *State {
	state := &State{
		list:   list,
		keyMap: keyMap{list: list.KeyMap},
	}

	for _, option := range options {
		option(state)
	}

	return state
}

type Item = list.DefaultItem

func New[T any](
	delegateHeight int,
	singular, plural string,
	items []T,
	convert func(T) Item,
	options ...Option,
) *State {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = convert(item)
	}

	delegate := list.NewDefaultDelegate()

	delegate.Styles.FilterMatch.Underline(false)
	delegate.Styles.NormalTitle.Bold(true)
	delegate.Styles.SelectedTitle.Bold(true)

	if delegateHeight == 1 {
		delegate.ShowDescription = false
	}

	delegate.SetHeight(delegateHeight)

	l := list.New(listItems, delegate, 0, 0)
	l.SetShowHelp(false)
	l.SetShowFilter(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.SetShowPagination(false)
	l.DisableQuitKeybindings()
	l.InfiniteScrolling = false
	l.KeyMap.CancelWhileFiltering = key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel"))

	l.Paginator.Type = paginator.Arabic

	l.SetStatusBarItemName(singular, plural)

	return NewWithList(l, options...)
}
