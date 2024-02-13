package storage

import "errors"

var (
	ErrBadArguments       = errors.New("введите правильныe аргументы в формате url описание теги")
	ErrNotFound           = errors.New("данный URL отсутствует в списке, не могу удалить")
	ErrTooShortSearchWord = errors.New("слишком короткое слово для поиска")
	ErrSearchNotFound     = errors.New("ни одного URL не найдено по Вашему запросу")
)
