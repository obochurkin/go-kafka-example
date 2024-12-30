package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckController struct {
}

func (c HealthCheckController) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK");
}