package handler

import (
	"github.com/DistributedPlayground/go-lib/dperror"
	"github.com/DistributedPlayground/go-lib/httperror"

	"github.com/labstack/echo/v4"
)

func DefaultErrorHandler(c echo.Context, err error, handlerName string) error {
	if err == nil {
		return nil
	}

	// always log the error
	// common.LogStringError(c, err, handlerName)

	if dperror.Is(err, dperror.NOT_FOUND) {
		return httperror.NotFound404(c)
	}

	if dperror.Is(err, dperror.FORBIDDEN) {
		return httperror.Forbidden403(c, "Invoking member lacks authority")
	}

	if dperror.Is(err, dperror.INVALID_RESET_TOKEN) {
		return httperror.BadRequest400(c, "Invalid password reset token")
	}

	if dperror.Is(err, dperror.INVALID_PASSWORD) {
		return httperror.BadRequest400(c, "Invalid password")
	}

	if dperror.Is(err, dperror.ALREADY_IN_USE) {
		return httperror.Conflict409(c, "Already in use")
	}

	if dperror.Is(err, dperror.INVALID_DATA) {
		return httperror.BadRequest400(c, "Invalid data")
	}

	return httperror.Internal500(c)
}
