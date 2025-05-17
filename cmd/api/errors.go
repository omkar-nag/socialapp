package main

import (
	"fmt"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("internal server error: %s\npath: %s\nerror: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("Bad request error: %s\npath: %s\nerror: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}
