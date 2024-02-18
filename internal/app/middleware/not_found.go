package middleware

import (
	"net/http"

	"gitlab.com/beabys/quetzal"
)

// NotFound handler middleware
func NotFound(w http.ResponseWriter, r *http.Request) {
	quetzal.ResponseJSON(w, http.StatusNotFound, "Not Found")
}
