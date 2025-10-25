package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/selfharm-enj/otus_m2_hw9/internal/model"
	"github.com/selfharm-enj/otus_m2_hw9/internal/repository"
	"github.com/selfharm-enj/otus_m2_hw9/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	if err := repository.InitData(); err != nil {
		fmt.Printf("Error initializing data: %v\n", err)
	}

	dataCh := make(chan model.IDReader)

	service.StartService(dataCh)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataCh)
				return
			case data, ok := <-dataCh:
				if !ok {
					return
				}
				repository.AddData(data)
			}
		}
	}()

	go func() {
		service.LogChanges()
	}()

	<-sigCh
	cancel()

	time.Sleep(500 * time.Millisecond)
}
