package http

import (
	"github.com/gorilla/mux"
)

// New creates and return an http router
func NewRouter() *mux.Router{
	return  mux.NewRouter()

}