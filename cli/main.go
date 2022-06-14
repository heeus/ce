/*
* Copyright (c) 2022-present unTill Pro, Ltd.
* @author Maxim Geraskin
 */

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/heeus/ce"
)

func init() {

	flag.Usage = func() {
		w := flag.CommandLine.Output() // may be os.Stderr - but not necessarily

		fmt.Println()

		fmt.Fprintf(w, "Usage: %s <command> [<options>]\n", os.Args[0])
		fmt.Fprintf(w, "Commands:\n")
		fmt.Fprintf(w, "  help\t\tPrint help\n")
		fmt.Fprintf(w, "  server\tStart server\n")
		fmt.Fprintf(w, "  version\tPrint version\n")
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
	}
}

func main() {

	var cfg ce.Config

	flag.IntVar(&cfg.AdminPort, "aport", 8080, "admin port, will be used for 127.0.0.1 only")
	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Println(ce.Version)
		return
	}
	if flag.Arg(0) == "server" {
		ce, cleanup, err := ce.Provide(cfg)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer cleanup()
		_ = ce.Run()
		return
	}
	flag.Usage()
}
