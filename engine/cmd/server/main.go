package main

import (
	"encoding/json"
	"engine/internal/core/entity"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func parseBody(body io.ReadCloser, object any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(object)
}

func UploadNewSchedule(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusCreated
	defer log.Println(r.URL.Path, r.Method, statusCode)

	schedule := entity.Schedule{}

	err := parseBody(r.Body, &schedule)
	if err != nil {
		fmt.Println(err)
		statusCode = http.StatusBadRequest
	}

	w.WriteHeader(statusCode)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/upload", UploadNewSchedule).Methods("POST")
	http.ListenAndServe(":2513", router)
}
