package httpassistant

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(body io.ReadCloser, object any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(object)
}

func Reply(w http.ResponseWriter, r *http.Request, statusCode int, data any) {
	log.Println(r.URL.Path, r.Method, statusCode)
	w.WriteHeader(statusCode)

	if data != nil {
		response := map[string]any{
			"data": data,
		}

		bytes, _ := json.Marshal(&response)
		w.Write(bytes)
	}
}
