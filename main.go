// Net Libram of Random Magical Effects - Go version!

// Net Libram: https://centralia.aquest.com/downloads/NLRMEv2.pdf
// Reddit post: https://www.reddit.com/r/Roll20/comments/a1d1le/net_libram_of_random_magical_effects_v2_as_a/
// Gist (of v1.2?): https://gist.github.com/slugnet/6985fff9c4e09a9176c456f63a13999f
// Roll20 code: https://drive.google.com/drive/folders/1FBgz48isMzw8qpZ5RF3_opUf5hPLQYUv
// Online generator with downloadable html: https://perchance.org/random-magical-effect

// To Do:
//     web page
//     export effects and durations as plain text/csv, in much the same format as the PDF.
//     locate any strings of the format 'number-d-number' and roll that many dice of that type.

package main

import (
	"embed"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//go:embed files/nlrme_v1.2.pdf
//go:embed files/nlrme_v2.pdf
var embeddedFiles embed.FS

type config struct {
	app struct {
		name string
	}
	flags struct {
		version    bool
		exportCSV  bool
		exportText bool
	}
	effects struct {
		descriptions map[int]string
		count        int
		number       int
	}
	durations struct {
		descriptions map[int]string
		count        int
		number       int
	}
}

func main() {
	var cfg config
	cfg.app.name = "Net Libram of Random Magical Effects, v2"

	cfg.effects.descriptions = cfg.getEffects()
	cfg.effects.count = len(cfg.effects.descriptions)

	cfg.durations.descriptions = cfg.getDurations()
	cfg.durations.count = len(cfg.durations.descriptions)

	flag.BoolVar(&cfg.flags.version, "version", false, "version info")
	flag.IntVar(&cfg.effects.number, "effect", 0, "choose a specific effect number")
	flag.IntVar(&cfg.durations.number, "duration", 0, "choose a specific duration number")
	flag.BoolVar(&cfg.flags.exportCSV, "export-csv", false, "export to CSV file")
	flag.BoolVar(&cfg.flags.exportText, "export-text", false, "export to plain text file")
	flag.Parse()

	fmt.Println(cfg.app.name)

	// Sanity testing of inputs.
	if cfg.effects.number < 0 || cfg.effects.number > cfg.effects.count {
		fmt.Printf("Sorry, effect must be between 1 and %d (except 0, which means 'random').\n", cfg.effects.count)
		os.Exit(1)
	}
	if cfg.durations.number < 0 || cfg.durations.number > cfg.durations.count {
		fmt.Printf("Sorry, duration must be between 1 and %d (except 0, which means 'random').\n", cfg.effects.count)
		os.Exit(1)
	}

	// Print version info and quit.
	if cfg.flags.version {
		fmt.Println("Effects:", prettyNumbers(cfg.effects.count))
		fmt.Println("Durations:", prettyNumbers(cfg.durations.count))
		os.Exit(0)
	}

	// Export as plain text.
	if cfg.flags.exportText {
		cfg.exportPlainText()
		os.Exit(0)
	}

	// Export as CSV.
	if cfg.flags.exportCSV {
		cfg.exportCSV()
		os.Exit(0)
	}

	if cfg.effects.number == 0 {
		cfg.effects.number = rand.IntN(cfg.effects.count) + 1
	}
	fmt.Printf("Effect #%s:\t%s\n", prettyNumbers(cfg.effects.number), cfg.effects.descriptions[cfg.effects.number])

	if cfg.durations.number == 0 {
		cfg.durations.number = rand.IntN(cfg.durations.count) + 1
	}
	fmt.Printf("Duration #%s:\t%s\n", prettyNumbers(cfg.durations.number), cfg.durations.descriptions[cfg.durations.number])

	fmt.Println("Good luck!")
}

func (cfg *config) exportPlainText() {
	plainText := cfg.createPlainText()
	data := []byte(plainText)

	err := os.WriteFile("nlrme.txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Print(plainText)
}

func (cfg *config) createPlainText() (output string) {
	output += "Effects\n"

	for i := 1; i <= cfg.effects.count; i++ {
		newI := i
		if newI == 10000 {
			newI = 0
		}

		output += fmt.Sprintf("%04d: %s\n", newI, cfg.effects.descriptions[i])
	}

	output += "\n"
	output += "Durations\n"

	for i := 1; i <= cfg.durations.count; i++ {
		output += fmt.Sprintf("%03d: %s\n", i, cfg.durations.descriptions[i])
	}

	return
}

func (cfg *config) exportCSV() {
	csv := cfg.createCSV()
	data := []byte(csv)

	err := os.WriteFile("nlrme.csv", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Print(csv)
}

func (cfg *config) createCSV() (output string) {
	output += "\"Effects\"\n"

	for i := 1; i <= cfg.effects.count; i++ {
		newI := i
		if newI == 10000 {
			newI = 0
		}

		output += fmt.Sprintf("\"%04d\",\"%s\"\n", newI, cfg.effects.descriptions[i])
	}

	output += "\n"
	output += "\"Durations\"\n"

	for i := 1; i <= cfg.durations.count; i++ {
		output += fmt.Sprintf("\"%03d\",\"%s\"\n", i, cfg.durations.descriptions[i])
	}

	return
}

func prettyNumbers(in int) string {
	return message.NewPrinter(language.English).Sprint(in)
}
