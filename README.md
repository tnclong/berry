## Berry

Berry colorizing printed string on ANSI terminals.
[ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code)

## Doc

https://godoc.org/github.com/tnclong/berry

## Install

```
go get -u -v github.com/tnclong/berry
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/tnclong/berry"
)

func main() {
	fmt.Println(berry.Dye("FgYellow", berry.FgYellow))
	fmt.Println(berry.Dye("BgYellow", berry.BgYellow))

	fmt.Println(berry.Dye("FgRed", berry.FgRed))
	fmt.Println(berry.Dye("BgRed", berry.BgRed))

	fmt.Println(berry.Effect("Italic", berry.Italic))
	fmt.Println(berry.Effect("Underline", berry.Underline))

	fmt.Println(berry.Dye("FgRGB 8-bit 88", berry.FgRGB, 88))
	fmt.Println(berry.Dye("BgRGB 8-bit 88", berry.BgRGB, 88))

	fmt.Println(berry.Dye("FgRGB 24-bit 1,1,1", berry.FgRGB, 1, 1, 1))
	fmt.Println(berry.Dye("BgRGB 24-bit 1,1,1", berry.BgRGB, 1, 1, 1))

	berry.Enable(false)
	fmt.Println(berry.Dye("FgYellow(false)", berry.FgYellow))
}
```

## Test

```
go test -v -p=1 -count=1 .
```
