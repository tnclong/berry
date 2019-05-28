package main

import (
	"fmt"
	"strings"

	"github.com/tnclong/berry"
)

func main() {
	colors := map[string][2]berry.R{
		"Red":     {{berry.FgRed}, {berry.BgRed}},
		"Green":   {{berry.FgGreen}, {berry.BgGreen}},
		"Yellow":  {{berry.FgYellow}, {berry.BgYellow}},
		"Blue":    {{berry.FgBlue}, {berry.BgBlue}},
		"Magenta": {{berry.FgMagenta}, {berry.BgMagenta}},
		"Cyan":    {{berry.FgCyan}, {berry.BgCyan}},
		"White":   {{berry.FgWhite}, {berry.BgWhite}},
	}

	for n, c := range colors {
		fmt.Println(
			c[0].S(n),
			strings.Repeat(" ", 10-len(n)),
			c[1].S(strings.Repeat(" ", 20)),
			strings.Repeat(" ", 3),
			berry.R{
				berry.Italic, berry.Underline, berry.Bright,
			}.S(c[0].S(n)),
		)
	}
}
