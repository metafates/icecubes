package list

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

var _ soda.State = (*State)(nil)

type State struct {
	title  title.Title
	list   list.Model
	keyMap keyMap
}

func (s *State) Internal() *list.Model {
	return &s.list
}

func (s *State) SetItems(items []Item) tea.Cmd {
	listItems := make([]list.Item, len(items))

	for i, item := range items {
		listItems[i] = item
	}

	return s.list.SetItems(listItems)
}

func (s *State) Filtering() bool {
	return s.list.FilterState() == list.Filtering
}

func (s *State) Destroy() {
}

func (s *State) Backable() bool {
	return s.list.FilterState() == list.Unfiltered
}

func (s *State) Resize(size soda.Size) tea.Cmd {
	s.list.SetSize(size.Width, size.Height)

	if s.list.Paginator.TotalPages > 5 {
		s.list.Paginator.Type = paginator.Arabic
	} else {
		s.list.Paginator.Type = paginator.Dots
	}

	return nil
}

func (s *State) Title() title.Title {
	return s.title
}

func (s *State) Status() string {
	if s.list.FilterState() == list.Filtering {
		return s.list.FilterInput.View()
	}

	if s.list.Paginator.TotalPages > 1 {
		return s.list.Paginator.View()
	}

	return ""
}

func (s *State) Subtitle() string {
	singular, plural := s.list.StatusBarItemName()

	var subtitle strings.Builder

	subtitle.Grow(max(len(singular), len(plural)) * 2)

	itemsCount := len(s.list.VisibleItems())
	subtitle.WriteString(strconv.Itoa(itemsCount))
	subtitle.WriteString(" ")

	if itemsCount == 1 {
		subtitle.WriteString(singular)
	} else {
		subtitle.WriteString(plural)
	}

	if s.list.FilterState() == list.FilterApplied {
		subtitle.WriteString(" ")
		subtitle.WriteString(fmt.Sprintf("%q", s.list.FilterValue()))
	}

	return subtitle.String()
}

func (s *State) KeyMap() help.KeyMap {
	return s.keyMap
}

func (s *State) Init(mh soda.ModelHandler) tea.Cmd {
	return nil
}

func (s *State) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	return cmd
}

func (s *State) View(mh soda.ModelHandler) string {
	return s.list.View()
}
