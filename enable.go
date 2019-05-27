package berry

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Enable is global setting that enabled/disabled berry.
// typically, call when application init.
func Enable(e bool) {
	if enabled != e {
		enabled = e
	}
}

var enabled = true

func init() {
	ini()
}

func ini() {
	// https://github.com/golang/go/issues/18153
	if !isTerminal(os.Stdout) || !isTerminal(os.Stderr) {
		enabled = false
	}

	if os.Getenv("TERM") == "dump" {
		enabled = false
	}

	if os.Getenv("CLICOLOR_FORCE") == "1" {
		enabled = true
	}
}

func isTerminal(f *os.File) bool {
	return terminal.IsTerminal(int(f.Fd()))
}
