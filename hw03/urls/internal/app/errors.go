package app

import "errors"

var (
	ErrExitApp   = errors.New("exit app")
	ErrAppKilled = errors.New("killed app")
)
