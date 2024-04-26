# Net Libram of Random Magical Effects - Go version

Someone shared the 'Net Libram of Random Magical Effects' PDF on our D&D Discord. I felt I should do a basic web version, for the group and for a laugh.


## Getting It

Use one of the following methods:

1. [Download a binary of the latest release](https://github.com/vaughany/nlrme/releases) (Linux or Windows).
2. Clone this repository (requires Go is installed): e.g. `git clone git@github.com:vaughany/nlrme.git`


## Running It

This will generate a random effect and duration, and quit:

* Source code: `go run .`
* Linux binary: `./nlrme`
* Windows binary: run `nlrme.exe`

You can always get help by running it with the flag `-h` for 'help', and there's `-version` for version info too.

If you just want to run it in shell, you can use the following flags:

**-duration** - show a specific 'duration'
**-effect** - choose a specific 'effect'
**-export-csv** - export everything to a CSV file in the current folder
**-export-text** - export everything to a plain text file in the current folder

If you want to run the web server, run it with the flag `-server`. It will start up on `http://localhost:8080` by default but you can specify:

**-host** - host or IP address the web server will run on
**-port** - port the web server will use
