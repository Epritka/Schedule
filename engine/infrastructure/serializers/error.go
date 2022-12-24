package serializers

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newHttpError(code int, message string, errors map[string]string) (int, map[string]any) {
	if len(errors) > 0 {
		return code, map[string]any{
			"success": false,
			"error": &httpError{
				Code:    code,
				Message: message,
			},
			"data": &errors,
		}
	}

	return code, map[string]any{
		"success": false,
		"error": &httpError{
			Code:    code,
			Message: message,
		},
	}
}
