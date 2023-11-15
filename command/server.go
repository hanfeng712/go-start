package main

import (
	"gtstart/command/apigateway"
	"gtstart/config"

	"github.com/urfave/cli/v2"
)

func serveCommand() *cli.Command {
	var cfg string
	cmd := &cli.Command{
		Name:  "serve",
		Usage: "启动 HTTP 服务",
		Action: func(context *cli.Context) error {
			config.ParseConfig(cfg)
			setups := Setup{
				SetupDB,
				SetupJwt,
			}
			setups.apply()
			apigateway.RunServer(&config.C)
			return nil
		},
	}
	cmd.Flags = append(cmd.Flags, &cli.StringFlag{
		Name: "config", Usage: "配置文件路径", Aliases: []string{"c"}, Destination: &cfg,
	})
	return cmd
}

func init() {
	app.Commands = append(app.Commands, serveCommand())
}
