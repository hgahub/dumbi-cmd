package main

import (
	"fmt"
	"github.com/hgahub/dumbi-cmd/internal"
	"github.com/hgahub/dumbi-cmd/internal/cli"
	"os"
)

// date is a ldflags parameter - release date
var date string

// version is a ldflags parameter - release version
var version string

// commit is a ldflags parameter - release commit
var commit string

func main() {
	internalVersion := internal.NewVersion(commit, version, date)
	if c, err := cli.CommandLineRead(internalVersion); err == nil && c.Name != "global" {
		fmt.Println(c)
	} else {
		os.Exit(1)
	}
}
