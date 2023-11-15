package main

import (
	"os"

	"github.com/urfave/cli/v2"

	_ "go.uber.org/automaxprocs"
)

var app = cli.NewApp()

func main() {
	app.Usage = "gtstart 使用方式"

	_ = app.Run(os.Args)
}
