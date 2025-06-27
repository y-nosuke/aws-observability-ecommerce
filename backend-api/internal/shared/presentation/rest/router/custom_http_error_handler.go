package router

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	ctx := c.Request().Context()

	var he *echo.HTTPError
	if errors.As(err, &he) {
		slog.WarnContext(ctx, "echo.HTTPError", "error", err)

		now := time.Now()
		res := &openapi.ErrorResponse{
			Code:      "",
			Details:   nil,
			Message:   http.StatusText(he.Code),
			Timestamp: &now,
			TraceId:   nil,
		}
		if resErr := c.JSON(he.Code, res); resErr != nil {
			slog.ErrorContext(ctx, "c.JSON()", "original error", err, "resErr", resErr)
		}
	} else {
		slog.ErrorContext(ctx, err.Error(), "error", err)

		now := time.Now()
		res := &openapi.ErrorResponse{
			Code:      "",
			Details:   nil,
			Message:   "internal server error",
			Timestamp: &now,
			TraceId:   nil,
		}
		if resErr := c.JSON(http.StatusInternalServerError, res); resErr != nil {
			slog.ErrorContext(ctx, "c.JSON()", "original error", err, "resErr", resErr)
		}
	}
}
