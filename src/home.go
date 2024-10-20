package main

import (
	"fmt"
	"net/http"
)

type Home struct {
}

func (h *Home) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!")
}
