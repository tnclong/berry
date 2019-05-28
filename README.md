## Berry

Berry colorizing printed string on ANSI terminals.
[ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code)

![Examples](https://github.com/tnclong/berry/blob/master/examples.png)

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
	fmt.Println(berry.Yellow.S("FgYellow"))
	fmt.Println(berry.R{berry.BgYellow}.S("BgYellow"))

	berry.Red.Println("FgRed")
	berry.R{berry.BgRed}.Println("BgRed")

	fmt.Println(berry.R{berry.Italic}.S("Italic"))
	fmt.Println(berry.R{berry.Underline}.S("Underline"))

	fmt.Println(berry.R{berry.FgSet, berry.Bit8, 88}.S("Fg Bit8"))
	fmt.Println(berry.R{berry.BgSet, berry.Bit8, 88}.S("Bg Bit8"))

	pFg := berry.Prepare(berry.R{berry.FgSet, berry.Bit24, 100, 100, 1})
	pFg.Println("Fg Bit24")
	pBg := berry.Prepare(berry.R{berry.BgSet, berry.Bit24, 100, 100, 1})
	pBg.Println("Bg Bit24")

	berry.Enable(false)
	fmt.Println(berry.Yellow.S("FgYellow(enabled=false)"))
}
```

## Test

```
go test -v -p=1 -count=1 .
```
