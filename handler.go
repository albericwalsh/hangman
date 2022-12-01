package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Game(w http.ResponseWriter, tmpl string) {
	renderTemplate(w, "game")
}