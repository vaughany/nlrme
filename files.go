package main

import (
	"fmt"
	"os"
)

func (cfg *config) export(expType exportType) {
	var (
		filename = "nlrme"
		data     []byte
	)

	switch expType {
	case expCSV:
		filename += ".csv"
		data = []byte(cfg.createCSV())
	case expPlain:
		filename += ".txt"
		data = []byte(cfg.createPlainText())
	}

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(filename, "created.")
}

func (cfg *config) createPlainText() (output string) {
	output += "Effects\n"

	for i := 1; i <= len(cfg.effects.descriptions); i++ {
		newI := i
		if newI == 10000 {
			newI = 0
		}

		output += fmt.Sprintf("%04d: %s\n", newI, cfg.effects.descriptions[i])
	}

	output += "\n"
	output += "Durations\n"

	for i := 1; i <= len(cfg.durations.descriptions); i++ {
		output += fmt.Sprintf("%03d: %s\n", i, cfg.durations.descriptions[i])
	}

	return
}

func (cfg *config) createCSV() (output string) {
	output += "\"Effects\"\n"

	for i := 1; i <= len(cfg.effects.descriptions); i++ {
		newI := i
		if newI == 10000 {
			newI = 0
		}

		output += fmt.Sprintf("\"%04d\",\"%s\"\n", newI, cfg.effects.descriptions[i])
	}

	output += "\n"
	output += "\"Durations\"\n"

	for i := 1; i <= len(cfg.durations.descriptions); i++ {
		output += fmt.Sprintf("\"%03d\",\"%s\"\n", i, cfg.durations.descriptions[i])
	}

	return
}
