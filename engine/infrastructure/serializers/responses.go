package serializers

import (
	"net/http"
)

// http ответ 400
func BadRequestHttpResponce(err error, errors map[string]string) (int, interface{}) {
	return newHttpError(
		http.StatusBadRequest,
		err.Error(),
		errors,
	)
}

// http ответ 404
func NotFoundHttpResponce(err error, errors map[string]string) (int, interface{}) {
	return newHttpError(
		http.StatusNotFound,
		err.Error(),
		errors,
	)
}

// http ответ 200
func SuccessHttpResponce(data interface{}) (int, map[string]interface{}) {
	if data == nil {
		return http.StatusOK, map[string]interface{}{
			"success": true,
		}
	}

	return http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    data,
	}
}
