package app

import (
	"fmt"
	"strings"
	"time"
)

// AddURL добавляет новый url в список хранения.
func (a *App) AddURL() error {
	fmt.Println("Введите новую запись в формате <url описание теги>")

	text, _ := a.reader.ReadString('\n')
	args := strings.Fields(strings.TrimSpace(text))

	return a.storage.Add(args)
}

// ListURLs возвращает список добавленных url.
//
// # Вывод в формате
//
// - Имя: <Описание>
//
// - URL: <url>
//
// - Теги: <Теги>
//
// - Дата: <дата>
func (a *App) ListURLs() error {
	urls, err := a.storage.List()
	if err != nil {
		return err
	}

	for k := range urls {
		fmt.Printf(
			"Имя: %s\nURL: %s\nТеги: %v\nДата: %s\n",
			urls[k].Description,
			urls[k].Link,
			urls[k].Tags,
			urls[k].Date.Format(time.DateTime),
		)
	}

	return nil
}

// RemoveURL удаляет url из списка хранения.
func (a *App) RemoveURL() error {
	fmt.Println("Введите имя ссылки, которое нужно удалить")

	url, _ := a.reader.ReadString('\n')
	url = strings.TrimSpace(url)

	return a.storage.Remove(url)
}
