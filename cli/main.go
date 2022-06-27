/*
 * Copyright (c) 2022-present unTill Pro, Ltd.
 * @author Maxim Geraskin
 */

package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/heeus/ce"
	logger "github.com/heeus/core-logger"
	"github.com/heeus/core/ibus"
	"github.com/heeus/core/ihttp"
	"github.com/heeus/core/iservices"
	"github.com/heeus/core/iservicesce"
	"github.com/heeus/core/iservicesctl"
)

func init() {

	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "\t%s [options] <command>\n", filepath.Base(os.Args[0]))
		fmt.Fprintln(w)
		fmt.Fprintf(w, "Commands:\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "\thelp\t\tprint help\n")
		fmt.Fprintf(w, "\tserver\t\tstart Server\n")
		fmt.Fprintf(w, "\tversion\t\tprint version\n")
		fmt.Fprintln(w)
		fmt.Fprintf(w, "Options:\n")
		fmt.Fprintln(w)
		flag.PrintDefaults()
	}
}

func main() {

	var httpCLIParams ihttp.CLIParams
	flag.IntVar(&httpCLIParams.Port, "ihttp.Port", Default_ihttp_Port, "")

	var busCLIParams ibus.CLIParams
	flag.IntVar(&busCLIParams.MaxNumOfConcurrentRequests, "ibus.MaxNumOfConcurrentRequests", Default_ibus_MaxNumOfConcurrentRequests, "")
	busCLIParams.ReadWriteTimeout = time.Nanosecond * Default_ibus_ReadWriteTimeoutNS

	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Println(ce.Version)
		return
	}
	if flag.Arg(0) == "server" {
		services, cleanup, err := iservicesce.ProvideCEServices(busCLIParams, httpCLIParams)
		if err != nil {
			fmt.Println("services not provided:", err)
			return
		}
		defer cleanup()
		run(services)
		return
	}
	flag.Usage()
}

var signals chan os.Signal

func run(services map[string]iservices.IService) {

	signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	ctl := iservicesctl.New()
	join, err := ctl.PrepareAndRun(ctx, services)
	if err != nil {
		cancel()
		fmt.Println("services preparation error:", err)
		return
	}
	defer join(ctx)

	sig := <-signals
	logger.Info("signal received:", sig)
	cancel()

}
