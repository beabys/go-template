package http

import (
	"net/http"

	"github.com/beabys/go-template/internal/application/example/command"
)

// HelloWorld implements the HTTP handler for Hello World
func (hs *HttpServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp, err := hs.ExampleService.GetHelloWorld(r.Context(), &command.GetHelloWorldRequest{})
	if err != nil {
		hs.Logger.Error("error from ExampleService", err)
		errorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}
	response := map[string]interface{}{
		"id":      resp.ID,
		"message": resp.Message,
	}
	successResponseJSON(w, response)
}
