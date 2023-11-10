package pluto

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var HTTPAdmin = func() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	return e
}()

func init() {
	HTTPAdmin.POST("/reload", func(ctx echo.Context) error {
		reloadConfig := ctx.QueryParam("config") == "true"

		if reloadConfig {
			var cfg Config
			if err := ctx.Bind(&cfg); err != nil {
				return err
			}

			ReloadExecutionCache(ResolveConfig(cfg))
		}

		return ctx.NoContent(http.StatusOK)
	})

	go Log.Fatal("Running HTTP admin server", zap.Error(HTTPAdmin.Start(Env.HTTPAdmin)))
}