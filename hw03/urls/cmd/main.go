package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/app"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/file"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/storage/memory"
	"golang.org/x/sync/errgroup"
)

func main() {
	filepath := flag.String("load", "", "Specify JSON file path")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	// storage := memory.New() // Change in-memory storage
	m, err := file.Load(*filepath)
	if err != nil || m.Urls == nil {
		log.Printf("cannot load data from file: %s\n", err)
		m = memory.New()
	}

	storage := file.New(m, *filepath) // To file storage
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

	err = g.Wait()
	switch err {
	case nil:
	case app.ErrAppExited:
		fmt.Println(app.AppExitedMsg)
	case io.EOF, app.ErrAppKilled:
		fmt.Println(app.AppKilledMsg)
	default:
		fmt.Printf(app.DefaultErrTemplate, err)
	}

	fmt.Println(app.ByeMsg)
}
