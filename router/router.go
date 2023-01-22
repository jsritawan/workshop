package router

import (
	"database/sql"
	"net/http"

	"github.com/kkgo-software-engineering/workshop/account"
	"github.com/kkgo-software-engineering/workshop/cloudpocket"
	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/kkgo-software-engineering/workshop/featflag"
	"github.com/kkgo-software-engineering/workshop/healthchk"
	mw "github.com/kkgo-software-engineering/workshop/middleware"
	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func RegRoute(cfg config.Config, logger *zap.Logger, db *sql.DB) *echo.Echo {
	e := echo.New()
	e.Use(mlog.Middleware(logger))
	e.Use(middleware.BasicAuth(mw.Authenicate()))

	hHealthChk := healthchk.New(db)
	e.GET("/healthz", hHealthChk.Check)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	hCloudPocket := cloudpocket.New(db)
	e.PUT("/cloud-pockets/:id", hCloudPocket.HandleUpdatePocket)
	e.POST("/cloud-pockets", hCloudPocket.HandleCreatePocket)
	e.POST("/cloud-pocket/:id/transfer", hCloudPocket.Transfer)
	e.GET("/cloud-pockets/:id", hCloudPocket.GetCloudpocketByID)
	e.GET("/cloud-pockets", hCloudPocket.GetCloudpocket)
	e.DELETE("/cloud-pockets/:id", hCloudPocket.Delete)

	hCloudPockets := cloudpocket.New(db)
	e.GET("/cloud-pockets", hCloudPockets.GetAll)

	hAccount := account.New(cfg.FeatureFlag, db)
	e.POST("/accounts", hAccount.Create)

	hFeatFlag := featflag.New(cfg)
	e.GET("/features", hFeatFlag.List)

	return e
}
