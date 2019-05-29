package main

import (
	"fmt"
	"strings"

	"github.com/tnclong/berry"
)

func main() {
	colors := map[string][2]berry.R{
		"Red":     {berry.Red, berry.New(berry.BgRed)},
		"Green":   {berry.Green, berry.New(berry.BgGreen)},
		"Yellow":  {berry.Yellow, berry.New(berry.BgYellow)},
		"Blue":    {berry.Blue, berry.New(berry.BgBlue)},
		"Magenta": {berry.Magenta, berry.New(berry.BgMagenta)},
		"Cyan":    {berry.Cyan, berry.New(berry.BgCyan)},
		"White":   {berry.White, berry.New(berry.BgWhite)},
	}

	for n, c := range colors {
		fmt.Println(
			c[0].S(n),
			strings.Repeat(" ", 10-len(n)),
			c[1].S(strings.Repeat(" ", 20)),
			strings.Repeat(" ", 3),
			berry.New(
				berry.Italic, berry.Underline, berry.Bright,
			).S(c[0].S(n)),
		)
	}
}
