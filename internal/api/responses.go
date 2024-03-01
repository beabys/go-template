package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
)

// SuccessResponseJSON return a response with status code
func SuccessResponseJSON(w http.ResponseWriter, data map[string]interface{}) {
	ResponseJSON(w, true, http.StatusOK, data)
}

// ErrorResponseJSON return an error response with status code
func ErrorResponseJSON(w http.ResponseWriter, statusCode int, err error) {
	data := map[string]interface{}{
		"error": err.Error(),
	}
	ResponseJSON(w, false, http.StatusOK, data)
}

// ResponseJSON return a response with status code
func ResponseJSON(w http.ResponseWriter, success bool, statusCode int, data map[string]interface{}) {
	var response = v1.Response{
		Success: success,
		Data:    data,
	}
	responseWriter(w, statusCode, &response)
}

// responseWriter print responseWritter
func responseWriter(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong :(")
	}
	fmt.Fprintf(w, string(responseData))
}
