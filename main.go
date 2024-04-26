// Net Libram of Random Magical Effects - Go version!

// Net Libram: https://centralia.aquest.com/downloads/NLRMEv2.pdf
// Reddit post: https://www.reddit.com/r/Roll20/comments/a1d1le/net_libram_of_random_magical_effects_v2_as_a/
// Gist (of v1.2?): https://gist.github.com/slugnet/6985fff9c4e09a9176c456f63a13999f
// Roll20 code: https://drive.google.com/drive/folders/1FBgz48isMzw8qpZ5RF3_opUf5hPLQYUv
// Online generator with downloadable html: https://perchance.org/random-magical-effect
// orrex@excite.com: https://www.reddit.com/r/DnD/comments/1527j6u/ultimate_dd_dm_resource_compilation_essential/

// To Do:
//     web page
//     export effects and durations as plain text/csv, in much the same format as the PDF.
//     locate any strings of the format 'number-d-number' and roll that many dice of that type.

package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//go:embed files/nlrme_v1.2.pdf
//go:embed files/nlrme_v2.pdf
var embeddedFiles embed.FS

type exportType int

const (
	expPlain exportType = iota
	expCSV   exportType = iota
)

type config struct {
	app struct {
		name string
	}
	flags struct {
		version    bool
		exportCSV  bool
		exportText bool
		webServer  bool
	}
	effects struct {
		descriptions map[int]string
		number       int
	}
	durations struct {
		descriptions map[int]string
		number       int
	}
	web struct {
		host string
		port string
	}
}

func main() {
	var cfg config
	cfg.app.name = "Net Libram of Random Magical Effects, v2"

	cfg.effects.descriptions = cfg.getEffects()
	cfg.durations.descriptions = cfg.getDurations()

	flag.BoolVar(&cfg.flags.version, "version", false, "version info")
	flag.IntVar(&cfg.effects.number, "effect", 0, "choose a specific effect number")
	flag.IntVar(&cfg.durations.number, "duration", 0, "choose a specific duration number")
	flag.BoolVar(&cfg.flags.exportCSV, "export-csv", false, "export to CSV file")
	flag.BoolVar(&cfg.flags.exportText, "export-text", false, "export to plain text file")
	flag.StringVar(&cfg.web.host, "host", "localhost", "host or IP address of web server")
	flag.StringVar(&cfg.web.port, "port", "8080", "port of web server")
	flag.BoolVar(&cfg.flags.webServer, "server", false, "run a web server")

	flag.Parse()

	fmt.Println(cfg.app.name)

	// Sanity testing of inputs.
	if cfg.effects.number < 0 || cfg.effects.number > len(cfg.effects.descriptions) {
		fmt.Printf("Sorry, effect must be between 1 and %d (except 0, which means 'random').\n", len(cfg.effects.descriptions))
		os.Exit(1)
	}
	if cfg.durations.number < 0 || cfg.durations.number > len(cfg.durations.descriptions) {
		fmt.Printf("Sorry, duration must be between 1 and %d (except 0, which means 'random').\n", len(cfg.durations.descriptions))
		os.Exit(1)
	}

	// Print version info and quit.
	if cfg.flags.version {
		fmt.Println("Effects:", prettyNumbers(len(cfg.effects.descriptions)))
		fmt.Println("Durations:", prettyNumbers(len(cfg.durations.descriptions)))
		os.Exit(0)
	}

	// Export as plain text.
	if cfg.flags.exportText {
		cfg.export(expPlain)
		os.Exit(0)
	}

	// Export as CSV.
	if cfg.flags.exportCSV {
		cfg.export(expCSV)
		os.Exit(0)
	}

	// Run a web server.
	if cfg.flags.webServer {
		mux := http.NewServeMux()
		mux.HandleFunc("/", cfg.rootHandler)
		mux.Handle("/files/", http.FileServer(http.FS(embeddedFiles)))
		mux.HandleFunc("/favicon.ico", cfg.notFoundHandler)

		serveURL := net.JoinHostPort(cfg.web.host, cfg.web.port)
		log.Printf("running web server on http://%s\n", serveURL)
		log.Fatal(http.ListenAndServe(serveURL, mux))

	} else {
		// If not running a web server, spit out random (or specific) 'effects' and 'durations' as configured by flags.
		if cfg.effects.number == 0 {
			cfg.effects.number = getRandomIndexFromMap(cfg.effects.descriptions)
		}
		fmt.Printf("Effect: %s.\n", getItemFromMap(cfg.effects.descriptions, cfg.effects.number))

		if cfg.durations.number == 0 {
			cfg.durations.number = getRandomIndexFromMap(cfg.durations.descriptions)
		}
		fmt.Printf("Duration: %s.\n", getItemFromMap(cfg.durations.descriptions, cfg.durations.number))

		fmt.Println("Good luck!")
	}
}

func getRandomIndexFromMap(in map[int]string) int {
	return rand.IntN(len(in)) + 1
}

func getItemFromMap(in map[int]string, item int) string {
	return fmt.Sprintf("#%s: %s", prettyNumbers(item), in[item])
}

func prettyNumbers(in int) string {
	return message.NewPrinter(language.English).Sprint(in)
}
