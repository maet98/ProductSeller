package main

import "net/http"

type buyers struct {
}

func (*buyers) getAll() http.HandlerFunc {
	return http.HandlerFunc(w http.ResponseWriter, 
}
