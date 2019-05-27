package main

import (
	"fmt"
	"strings"

	"github.com/tnclong/berry"
)

func main() {
	colors := map[string][2]berry.D{
		"Red":     {berry.FgRed, berry.BgRed},
		"Green":   {berry.FgGreen, berry.BgGreen},
		"Yellow":  {berry.FgYellow, berry.BgYellow},
		"Blue":    {berry.FgBlue, berry.BgBlue},
		"Magenta": {berry.FgMagenta, berry.BgMagenta},
		"Cyan":    {berry.FgCyan, berry.BgCyan},
		"White":   {berry.FgWhite, berry.BgWhite},
	}

	for n, c := range colors {
		fmt.Println(
			berry.Dye(n, c[0]),
			strings.Repeat(" ", 10-len(n)),
			berry.Dye(strings.Repeat(" ", 20), c[1]),
			strings.Repeat(" ", 3),
			berry.Effect(
				berry.Effect(
					berry.Effect(berry.Dye(n, c[0]), berry.Italic),
					berry.Underline,
				),
				berry.Bright,
			),
		)
	}
}
