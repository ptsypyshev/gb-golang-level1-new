package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/memory"
	"golang.org/x/sync/errgroup"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	storage := memory.New()
	a := app.New(storage, reader)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return a.Run(ctx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		select {
		case <-quit:
			signal.Stop(quit)
		case <-ctx.Done():
			signal.Stop(quit)
			return ctx.Err()
		}

		a.Shutdown(cancel)
		return app.ErrAppKilled
	})

	err := g.Wait()
	switch err {
	case nil:
	case app.ErrExitApp:
		fmt.Println("Завершение по требованию пользователя")
	case io.EOF, app.ErrAppKilled:
		fmt.Println("Завершение по системному сигналу")
	default:
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("До новых встреч!")
}
