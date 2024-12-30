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

	return fmt.Errorf("invalid value received: %s", value)
}

func (e *EnumValue[T]) Get() any {
	if e.selected == "" {
		return e.Default
	}
	return e.selected
}

func (e *EnumValue[T]) String() string {
	return string(e.Get().(T))
}

func (e *EnumValue[T]) Selected() T {
	return e.Get().(T)
}
