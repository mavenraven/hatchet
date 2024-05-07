package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/hatchet-dev/hatchet/pkg/cmdutils"
	"github.com/hatchet-dev/hatchet/pkg/worker"
)

type printOutput struct{}

func print(ctx context.Context, num int) (result *printOutput, err error) {
	fmt.Printf("called print:print for %v\n", num)

	return &printOutput{}, nil
}

func printWithNum(num int) func(ctx context.Context) (result *printOutput, err error) {
	return func(ctx context.Context) (result *printOutput, err error) {
		return print(ctx, num)
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	client, err := client.New()

	if err != nil {
		panic(err)
	}

	w, err := worker.NewWorker(
		worker.WithClient(
			client,
		),
	)

	if err != nil {
		panic(err)
	}

	for i := 0; i < 5000; i++ {
		j := i
		err = w.On(
			worker.Cron("* * * * *"),
			worker.Fn(func(ctx context.Context) (result *printOutput, err error) {
				return print(ctx, j)
			}).SetName(fmt.Sprintf("chron%v", i)))
	}

	if err != nil {
		panic(err)
	}

	interrupt := cmdutils.InterruptChan()

	cleanup, err := w.Start()
	if err != nil {
		panic(err)
	}

	<-interrupt

	if err := cleanup(); err != nil {
		panic(fmt.Errorf("error cleaning up: %w", err))
	}
}
