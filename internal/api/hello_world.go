package api

import (
	"net/http"
)

// Used to implement the methods of helloworld interface
// from openapi

// HelloWorld implements the method Hello World
func (hs *HttpServer) HelloWorld(w http.ResponseWriter, r *http.Request) {
	hs.HelloWorldSvc.GetHelloWorld(r)
	SuccessResponseJSON(w, nil)
}
