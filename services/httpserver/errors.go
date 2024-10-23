package httpserver

import (
	"errors"
	"net/http"
	apperrors "telegraminput/lib/errors"
)

func mapInternalErrorToHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, apperrors.ErrTooManyRequest):
		return http.StatusBadRequest // 400
	case errors.Is(err, apperrors.ErrNotFound):
		return http.StatusNotFound // 404
	case errors.Is(err, apperrors.ErrTimout):
		return http.StatusRequestTimeout // 408
	case errors.Is(err, apperrors.ErrExternal):
		return http.StatusBadGateway // 502
	default:
		return http.StatusInternalServerError // 500
	}
}
