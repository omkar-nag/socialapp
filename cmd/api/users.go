package main

import "net/http"

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for creating a new user
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"response":"User created"}`))
}
