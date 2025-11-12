package core

import "errors"

var ErrUnableToMap = errors.New("unable to map")

func Map[T any](in any) (*T, error) {
	var v T

	v, ok := in.(T)
	if !ok {
		return nil, ErrUnableToMap
	}

	return &v, nil
}
