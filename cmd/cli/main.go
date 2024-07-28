package main

import (
	"github.com/yuchou87/rest-agent/internal/cli"
)

var (
	version = "dev"
	commit  = "HEAD"
	date    = "unknown"
)

func main() {
	cli.Execute(version, commit, date)
}
