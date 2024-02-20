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
	SearchCmd = "s"
	QuitCmd   = "q"
)

// Storage is an interface to store URL items.
type Storage interface {
	Add(args []string) error
	List() ([]models.URL, error)
	Remove(url string) error
	Search(t string) ([]models.URL, error)
	Close(ctx context.Context) error
}

// App is a main structure our application.
type App struct {
	storage Storage
	reader  *bufio.Reader
}

// New is a constructor for App.
func New(s Storage, r *bufio.Reader) *App {
	return &App{
		storage: s,
		reader:  r,
	}
}

// Run is a method which used to start our application.
func (a *App) Run(ctx context.Context) error {
	defer a.storage.Close(ctx)
	var (
		cmd  string
		err  error
		exit = make(chan error)
	)

	fmt.Println(HelloMsg)
	fmt.Println(QuitHelpMsg)

	go func() {
		for {
			fmt.Print(CommandsHelpMsg)
			cmd, err = a.reader.ReadString('\n')
			if err != nil {
				exit <- ErrAppExited
				return
			}

			cmd = strings.TrimSpace(cmd)
			cmd = string(strings.ToLower(cmd)[0])
			switch cmd {
			case AddCmd:
				err = a.AddURL()
			case ListCmd:
				err = a.ListURLs()
			case RemoveCmd:
				err = a.RemoveURL()
			case SearchCmd:
				err = a.Search()
			case QuitCmd:
				exit <- ErrAppExited
				return
			default:
				fmt.Printf(BadCommandErrTemplate, cmd)
			}

			if err != nil {
				fmt.Printf(ErrMsgTemplate, err)
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

// Shutdown is a method which used to stop our application.
func (a *App) Shutdown(cancel context.CancelFunc) {
	cancel()
}
