package main

import (
	"github.com/hgahub/dumbi-cmd/internal"
	"github.com/hgahub/dumbi-cmd/internal/cli"
	"github.com/hgahub/dumbi-cmd/internal/cli/command"
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
	if c, err := cli.CommandLineRead(internalVersion); err == nil && c.Name != command.GlobalFlagCommandName {
		os.Exit(c.Execute())
	} else {
		os.Exit(128)
	}
}
