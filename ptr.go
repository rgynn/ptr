package ptr

import (
	"fmt"
)

func Convert[T any](input T) *T {
	return &input
}

func String[T any](input *T) string {
	if input == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", *input)
}
