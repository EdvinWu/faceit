package healtcheck

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type HealthCheck struct {
	db *sqlx.DB
}

func NewHealthCheck(db *sqlx.DB) HealthCheck {
	return HealthCheck{db: db}
}

func (h *HealthCheck) Ping(ctx echo.Context) error {
	timeoutCtx, cancelFunc := context.WithTimeout(ctx.Request().Context(), time.Second*10)
	defer cancelFunc()
	err := h.db.PingContext(timeoutCtx)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "")
	}
	return ctx.String(http.StatusOK, "pong")
}
