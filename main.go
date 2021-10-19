package main

import (
	"fmt"
	"github.com/hgahub/dumbi-cmd/internal"
	"github.com/hgahub/dumbi-cmd/internal/cli"
	"os"
)

var date string
var version string
var commit string

func main() {
	internalVersion := internal.NewVersion(commit, version, date)
	c, err := cli.Parse(internalVersion)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if c != nil {
		fmt.Println(c)
	}
}
