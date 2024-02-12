package storage

import "errors"

var (
	ErrBadArguments = errors.New("введите правильный аргументы в формате url описание теги")
	ErrNotFound     = errors.New("данный URL отсутствует в списке, не могу удалить")
)
