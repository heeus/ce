/*
 * Copyright (c) 2022-present unTill Pro, Ltd.
 * @author Maxim Geraskin
 */

package main

import (
	"os"

	"github.com/heeus/core/cmd/ce"

	_ "embed"
)

//go:embed version
var Version string

func main() {
	os.Exit(ce.CLI(Version))
}
