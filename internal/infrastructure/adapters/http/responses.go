package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	v1 "github.com/beabys/go-template/internal/api/v1"
)

func errorResponse(err error) v1.Error {
	errResponse := v1.Error{}
	errMsg := err.Error()
	errResponse.Data.Error = &errMsg
	return errResponse
}

func errorResponseJSON(w http.ResponseWriter, statusCode int, err error) {
	responseWriter(w, statusCode, errorResponse(err))
}

func middlewareRespHandler(w http.ResponseWriter, msg string, statusCode int) {
	errFormatted := errorResponse(errors.New(msg))
	responseWriter(w, statusCode, errFormatted)
}

func successResponseJSON(w http.ResponseWriter, data map[string]interface{}) {
	successResponse := v1.Success{}
	successResponse.Success = true
	successResponse.Data = data
	responseWriter(w, http.StatusOK, successResponse)
}

func responseWriter(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseData = []byte("Something went wrong :(")
	}
	fmt.Fprintln(w, string(responseData))
}
