package serializers

import (
	"net/http"
)

// http ответ 400
func BadRequestHttpResponce(err error, errors map[string]string) (int, any) {
	return newHttpError(
		http.StatusBadRequest,
		err.Error(),
		errors,
	)
}

// http ответ 404
func NotFoundHttpResponce(err error, errors map[string]string) (int, any) {
	return newHttpError(
		http.StatusNotFound,
		err.Error(),
		errors,
	)
}

// http ответ 200
func SuccessHttpResponce(data any) (int, map[string]any) {
	if data == nil {
		return http.StatusOK, map[string]any{
			"success": true,
		}
	}

	return http.StatusOK, map[string]any{
		"success": true,
		"data":    data,
	}
}
