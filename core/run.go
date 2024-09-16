package core

import (
	"net/http"

	"github.com/go-zoox/core-utils/fmt"

	"github.com/go-idp/terminal"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"

	ts "github.com/go-zoox/terminal/server"
)

func (s *server) Run(cfg *Config) error {
	if cfg.Path == "" {
		cfg.Path = "/"
	}
	if cfg.Driver == "" {
		cfg.Driver = "host"
	}

	logger.Infof("config: %+v", cfg)

	app := defaults.Default()

	app.SetBanner(fmt.Sprintf(`
  _____       _______  ___    ______              _           __
 / ___/__    /  _/ _ \/ _ \  /_  __/__ ______ _  (_)__  ___ _/ /
/ (_ / _ \  _/ // // / ___/   / / / -_) __/  ' \/ / _ \/ _ '/ / 
\___/\___/ /___/____/_/      /_/  \__/_/ /_/_/_/_/_//_/\_,_/_/  %s

Terminal Service for Go IDP

____________________________________O/_______
                                    O\
`, terminal.Version))

	app.Use(func(ctx *zoox.Context) {
		s.requests.Inc(1)
		ctx.Next()
	})

	if cfg.Username != "" || cfg.Password != "" {
		app.Use(func(ctx *zoox.Context) {
			user, pass, ok := ctx.Request.BasicAuth()
			if !ok {
				ctx.Set("WWW-Authenticate", `Basic realm="go-idp"`)
				ctx.Status(http.StatusUnauthorized)
				return
			}

			if !(user == cfg.Username && pass == cfg.Password) {
				ctx.Status(http.StatusUnauthorized)
				return
			}

			ctx.Next()
		})
	}

	app.WebSocket(cfg.Path, func(opt *zoox.WebSocketOption) {
		server, err := ts.Serve(&ts.Config{
			Driver:      cfg.Driver,
			DriverImage: cfg.DriverImage,
			// InitCommand:       cfg.InitCommand,
		})
		if err != nil {
			panic(fmt.Errorf("failed to create websocket server: %s", err))
		}

		opt.Server = server

		opt.Middlewares = append(opt.Middlewares, func(ctx *zoox.Context) {
			id := ctx.RequestID()

			initCommand := ctx.Query().Get("init_command")
			if initCommand != "" {
				logger.Infof("[%s] init command: %s", id, initCommand)
			}

			if environments, ok := ctx.Request.URL.Query()["environment"]; ok {
				for _, env := range environments {
					logger.Infof("[%s] environment: %s", id, env)
				}
			}

			ctx.Next()
		})
	})

	app.Get("/web", func(ctx *zoox.Context) {
		ctx.HTML(200, ts.RenderXTerm(zoox.H{
			"wsPath": cfg.Path,
		}))
	})

	return app.Run(fmt.Sprintf(":%d", cfg.Port))
}
