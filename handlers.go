package main

import (
	"fmt"
	"net/http"
)

func (cfg *config) rootHandler(w http.ResponseWriter, r *http.Request) {
	var (
		output    string
		effect    = getItemFromMap(cfg.effects.descriptions, getRandomIndexFromMap(cfg.effects.descriptions))
		durations = getItemFromMap(cfg.durations.descriptions, getRandomIndexFromMap(cfg.durations.descriptions))
	)

	output += fmt.Sprintf("<h1>%s</h1>\n", cfg.app.name)
	output += fmt.Sprintf("<p><strong>Effect:</strong> %s.</p>\n", effect)
	output += fmt.Sprintf("<p><strong>Duration:</strong> %s.</p>\n", durations)
	output += "<p><small>(Original PDF files by Orrex available: <a href='/files/nlrme_v2.pdf'>v2</a> and <a href='/files/nlrme_v1.2.pdf'>v1.2</a>.</small></p>\n"

	w.Write([]byte(output))
}

func (cfg *config) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}
