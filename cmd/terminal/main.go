package main

import (
	"github.com/go-idp/terminal"
	"github.com/go-idp/terminal/core"
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "terminal",
		Usage:   "Terminal Service for IDP",
		Version: terminal.Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Usage:   "server port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "path",
				Usage:   "server path",
				EnvVars: []string{"WS_PATH"},
				Value:   "/",
			},
			&cli.StringFlag{
				Name:    "username",
				Usage:   "username for basic auth",
				Aliases: []string{"u"},
				EnvVars: []string{"USERNAME"},
			},
			&cli.StringFlag{
				Name:    "password",
				Usage:   "password for basic auth",
				Aliases: []string{"P"},
				EnvVars: []string{"PASSWORD"},
			},
			&cli.StringFlag{
				Name:    "driver",
				Usage:   "Driver runtime, options: host, docker, kubernetes, ssh, default: host",
				Aliases: []string{"d"},
				EnvVars: []string{"DRIVER"},
			},
			&cli.StringFlag{
				Name:    "driver-image",
				Usage:   "Driver image",
				Aliases: []string{"i"},
				EnvVars: []string{"DRIVER_IMAGE"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		cfg := &core.Config{
			Port:        ctx.Int("port"),
			Path:        ctx.String("path"),
			Username:    ctx.String("username"),
			Password:    ctx.String("password"),
			Driver:      ctx.String("driver"),
			DriverImage: ctx.String("driver-image"),
		}

		return core.New().Run(cfg)
	})

	app.Run()
}
