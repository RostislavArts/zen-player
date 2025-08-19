package utils

import (
	"os"
	"fmt"
	"flag"
	"strings"
)

type Config struct {
	NoiseList []string
	Loop bool
	Shuffle bool
	Path string
}

var parsed bool
var isHelp bool
var isList bool

func ParseFlags() *Config {
	if parsed { return nil }
	parsed = true

	var noiseArg string
	var loopFlag bool
	var shuffleFlag bool

	flag.StringVar(&noiseArg, "noise", "", "Comma-separated list of background noises (or 'list' to show available noises)")
	flag.BoolVar(&loopFlag, "loop", false, "Loop the track")
	flag.BoolVar(&loopFlag, "l", false, "Loop the track (shorthand)")
	flag.BoolVar(&shuffleFlag, "shuffle", false, "Shuffle tracks")
	flag.BoolVar(&shuffleFlag, "s", false, "Shuffle tracks (shorthand)")
	flag.BoolVar(&isHelp, "help", false, "Show help message")
	flag.BoolVar(&isHelp, "h", false, "Show help message (shorthand)")

	flag.Parse()

	if isHelp {
		printHelp()
		os.Exit(0)
	}

	// Handle noise flag
	noises := []string{}
	if noiseArg == "list" {
		isList = true
	} else if noiseArg != "" {
		noises = strings.Split(noiseArg, ",")
	}

	if isList {
		fmt.Println("Noises included: night, rain, river, sea, thunder, wind")
		os.Exit(1)
	}

	// Positional argument (path)
	args := flag.Args()
	var path string
	if len(args) > 0 {
		path = args[0]
	}

	return &Config{
		NoiseList:  noises,
		Loop:       loopFlag,
		Shuffle:    shuffleFlag,
		Path:       path,
	}
}

func printHelp() {
	fmt.Println(`Usage: zen [options] <path>

A minimal CLI-based ambient audio player with optional background noises.

Arguments:
  <path>             Path to a file or folder with audio tracks to play.
                     Supported formats: WAV, MP3, FLAC.

Options:
  --noise <list>     Comma-separated list of background noises to play in parallel.
                     Available: night, rain, river, sea, thunder, wind.
                     Default: rain
                     Example: --noise rain,wind

  -l, --loop         Repeat the entire playlist or single file in a loop.

  -s, --shuffle      Shuffle tracks from provided folder.

  --noise list       Print the list of available noise options and exit.

  -h, --help         Show this help message and exit.

Behaviour:
  By default, if no --noise option is provided, 'rain' noise will be played.

Examples:
  zen --shuffle music/
  zen --loop sound.mp3
  zen --noise sea,wind ambient/
  zen -s -l --noise night playlist/`)
}

