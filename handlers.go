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

	output += fmt.Sprintf(`<h1>%s</h1>`, cfg.app.name)
	output += fmt.Sprintf(`<p><strong>Effect:</strong> %s.</p>`, effect)
	output += fmt.Sprintf(`<p><strong>Duration:</strong> %s.</p>`, durations)

	w.Write([]byte(output))
}
