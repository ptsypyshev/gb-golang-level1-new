package app

import "errors"

var (
	ErrAppExited = errors.New("exit app")
	ErrAppKilled = errors.New("killed app")
)
