package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
)

func errorResponse(err error) v1.Error {
	errResponse := v1.Error{}
	errMsg := err.Error()
	errResponse.Data.Msg = &errMsg
	return errResponse
}

// errorResponseJSON return an error response with status code
func errorResponseJSON(w http.ResponseWriter, statusCode int, err error) {
	responseWriter(w, statusCode, errorResponse(err))
}

// middlewareRespHandler return an error response with status code for middlewares
func middlewareRespHandler(w http.ResponseWriter, msg string, statusCode int) {
	errFormatted := errorResponse(errors.New(msg))
	responseWriter(w, statusCode, errFormatted)
}

// successResponseJSON return an error response with status code
func successResponseJSON(w http.ResponseWriter, data map[string]interface{}) {
	successResponse := v1.Success{}
	successResponse.Success = true
	successResponse.Data = data
	responseWriter(w, http.StatusOK, successResponse)
}

// responseWriter print responseWritter
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
