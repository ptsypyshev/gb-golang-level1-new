package app

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
)

const (
	AddCmd    = "a"
	ListCmd   = "l"
	RemoveCmd = "r"
	QuitCmd   = "q"
)

// Storage is an interface to store URL items
type Storage interface {
	Add(args []string) error
	List() ([]models.URL, error)
	Remove(url string) error
	Close(ctx context.Context) error
}

// App is a main structure our application
type App struct {
	storage Storage
	reader  *bufio.Reader
}

// New is a constructor for App
func New(s Storage, r *bufio.Reader) *App {
	return &App{
		storage: s,
		reader:  r,
	}
}

// Run is a method which used to start our application
func (a *App) Run(ctx context.Context) error {
	defer a.storage.Close(ctx)
	var (
		cmd  string
		err  error
		exit = make(chan error)
	)

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите 'q'")

	go func() {
		for {
			fmt.Println("Введите 'a', 'l', 'r' или 'q' для выбора команды")
			cmd, err = a.reader.ReadString('\n')
			if err != nil {
				exit <- ErrExitApp
				return
			}

			cmd = strings.TrimSpace(cmd)
			switch cmd {
			case AddCmd:
				err = a.AddURL()
			case ListCmd:
				err = a.ListURLs()
			case RemoveCmd:
				err = a.RemoveURL()
			case QuitCmd:
				exit <- ErrExitApp
				return
			default:
				fmt.Printf("Команда '%s' не верная, повторите\n\n", cmd)
			}

			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
			}
		}
	}()

	select {
	case <-ctx.Done():
		return err
	case err = <-exit:
		return err
	}
}

// Shutdown is a method which used to stop our application
func (a *App) Shutdown(cancel context.CancelFunc) {
	cancel()
}
