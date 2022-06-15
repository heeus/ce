/*
 * Copyright (c) 2022-present unTill Pro, Ltd.
 * @author Maxim Geraskin
 */

package ce

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	logger "github.com/heeus/core-logger"
)

type ce struct {
	cfg Config
	wg  sync.WaitGroup
}

var signals chan os.Signal

func (ce *ce) Run() error {

	logger.Info(fmt.Sprintf("config: %+v", ce.cfg))

	ctx, cancel := context.WithCancel(context.Background())

	signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	ctx, err := ce.start(ctx)
	if nil != err {
		cancel()
		return err
	}

	sig := <-signals
	logger.Info("signal received:", sig)
	cancel()
	return ce.join(ctx)
}

func (ce *ce) start(ctx context.Context) (context.Context, error) {
	ce.wg.Add(1)
	go ce.run(ctx)
	return ctx, nil
}

func (ce *ce) run(ctx context.Context) {
	defer ce.wg.Done()
	logger.Info("Server started")
	for ctx.Err() == nil {
		logger.Info("running")
		time.Sleep(1 * time.Second)
	}
	logger.Info("Server finished")
}

func (ce *ce) join(ctx context.Context) (err error) {
	logger.Info("waiting for the Server...")
	ce.wg.Wait()
	logger.Info("done")
	return nil
}
