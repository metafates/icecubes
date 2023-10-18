package stateinfo

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/metafates/icecubes/stateinfo/infosupplier"
	"github.com/metafates/soda"
	"github.com/metafates/soda/title"
)

type State[T soda.State] struct {
	state    T
	supplier infosupplier.Supplier
}

func (s *State[T]) State() T {
	return s.state
}

func (s *State[T]) Destroy() {
	s.state.Destroy()
}

func (s *State[T]) Backable() bool {
	return s.state.Backable()
}

func (s *State[T]) Resize(size soda.Size) tea.Cmd {
	return s.state.Resize(size)
}

func (s *State[T]) Title() title.Title {
	supplierTitle := s.supplier.Title()
	if supplierTitle.String() == "" {
		return s.state.Title()
	}

	return supplierTitle
}

func (s *State[T]) Subtitle() string {
	wrapped := s.state.Subtitle()

	if wrapped == "" {
		return s.supplier.Subtitle()
	}

	return wrapped + " " + s.supplier.Subtitle()
}

func (s *State[T]) Status() string {
	wrapped := s.state.Status()

	if wrapped == "" {
		return s.supplier.Status()
	}

	return wrapped + " " + s.supplier.Status()
}

func (s *State[T]) KeyMap() help.KeyMap {
	return s.state.KeyMap()
}

func (s *State[T]) Init(mh soda.ModelHandler) tea.Cmd {
	return s.state.Init(mh)
}

func (s *State[T]) Update(mh soda.ModelHandler, msg tea.Msg) tea.Cmd {
	return s.state.Update(mh, msg)
}

func (s *State[T]) View(mh soda.ModelHandler) string {
	return s.state.View(mh)
}
