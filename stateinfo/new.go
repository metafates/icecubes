package stateinfo

import (
	"github.com/metafates/icecubes/stateinfo/infosupplier"
	"github.com/metafates/soda"
)

func New[T soda.State](supplier infosupplier.Supplier, state T) *State[T] {
	return &State[T]{
		state:    state,
		supplier: supplier,
	}
}
