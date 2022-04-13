package server

import (
	userHandler "faceit-test/internal/domain/user/handler"
	healthCheckHandler "faceit-test/internal/healtcheck"
	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"

	"net/http"
)

func Echo(addr string, userHandler userHandler.User, healthCheckHandler healthCheckHandler.HealthCheck) *echo.Echo {
	e := echo.New()
	e.Use(echoMW.Recover())
	e.HideBanner = true

	e.Use(echoMW.CORSWithConfig(echoMW.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Server.Addr = addr
	e.GET("/ping", healthCheckHandler.Ping)
	group := e.Group("api/v1")
	registerUserRoutes(group, userHandler)
	return e
}

func registerUserRoutes(group *echo.Group, handler userHandler.User) {
	userGroup := group.Group("/user")
	userGroup.POST("", handler.Create)
	userGroup.PUT("/:id", handler.Update)
	userGroup.DELETE("/:id", handler.Delete)
	userGroup.GET("", handler.Paged)
}
