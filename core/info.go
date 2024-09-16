package core

import (
	"github.com/go-idp/terminal"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/zoox"
)

func (s *server) Info() zoox.HandlerFunc {
	runningAt := datetime.Now().Format("YYYY-MM-DD HH:mm:ss")

	return func(ctx *zoox.Context) {
		ctx.JSON(200, zoox.H{
			"name":    "terminal service for idp",
			"version": terminal.Version,
			"status": map[string]any{
				"requests":   s.requests.Get(),
				"running_at": runningAt,
			},
		})
	}
}
