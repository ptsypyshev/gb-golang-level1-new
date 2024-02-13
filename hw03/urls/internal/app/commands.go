package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
)

// AddURL adds a new url to the storage.
func (a *App) AddURL() error {
	fmt.Println(AddCmdHelpMsq)

	text, _ := a.reader.ReadString('\n')
	args := strings.Fields(strings.TrimSpace(text))

	return a.storage.Add(args)
}

// ListURLs returns a list of added urls.
//
// # Output in format
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

	drawURLs(urls)

	return nil
}

// RemoveURL removes a url from the storage.
func (a *App) RemoveURL() error {
	fmt.Println(RemoveCmdHelpMsq)

	url, _ := a.reader.ReadString('\n')
	url = strings.TrimSpace(url)

	return a.storage.Remove(url)
}

// Search searches the url in the storage list.
func (a *App) Search() error {
	fmt.Print(SearchCmdHelpMsq)

	text, _ := a.reader.ReadString('\n')
	text = strings.TrimSpace(text)

	urls, err := a.storage.Search(text)
	if err != nil {
		return err
	}

	drawURLs(urls)

	return nil
}

// drawURLs draws the list of urls.
func drawURLs(urls []models.URL) {
	for k := range urls {
		fmt.Printf(
			ListCmdTemplate,
			urls[k].Description,
			urls[k].Link,
			urls[k].Tags,
			urls[k].Date.Format(time.DateTime),
		)
	}
}
