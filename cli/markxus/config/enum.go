package config

import (
	"fmt"
)

type EnumValue[T ~string] struct {
	Enum     []T
	Default  T
	selected T
}

func (e *EnumValue[T]) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == T(value) {
			e.selected = T(value)
			return nil
		}
	}

	return fmt.Errorf("only accept %s", e.Enum)
}

func (e *EnumValue[T]) Get() any {
	return e
}

func (e *EnumValue[T]) String() string {
	return string(e.Selected())
}

func (e *EnumValue[T]) Selected() T {
	return e.selected
}
