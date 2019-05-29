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

## Benchmark

```
$ go test -bench=. -benchmem -count=1
goos: darwin
goarch: amd64
pkg: github.com/tnclong/berry
BenchmarkS1-8                	20000000	        57.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkS10-8               	20000000	        61.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkS50-8               	20000000	        67.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkS100-8              	20000000	        79.3 ns/op	     112 B/op	       1 allocs/op
BenchmarkSS100-8             	 3000000	       466 ns/op	     368 B/op	       6 allocs/op
BenchmarkS100WithoutP-8      	20000000	       104 ns/op	     117 B/op	       2 allocs/op
BenchmarkBestS100-8          	20000000	        65.2 ns/op	     112 B/op	       1 allocs/op
BenchmarkSprint100-8         	 5000000	       315 ns/op	     197 B/op	       5 allocs/op
BenchmarkBestSprint100-8     	10000000	       171 ns/op	     128 B/op	       2 allocs/op
BenchmarkSprintf100-8        	10000000	       213 ns/op	     144 B/op	       3 allocs/op
BenchmarkBestSprintf100-8    	10000000	       143 ns/op	     128 B/op	       2 allocs/op
BenchmarkS500-8              	10000000	       134 ns/op	     512 B/op	       1 allocs/op
BenchmarkS1000-8             	10000000	       195 ns/op	    1024 B/op	       1 allocs/op
BenchmarkSS1000-8            	 2000000	       867 ns/op	    3104 B/op	       6 allocs/op
BenchmarkS1000WithoutP-8     	10000000	       224 ns/op	    1029 B/op	       2 allocs/op
BenchmarkBestS1000-8         	10000000	       183 ns/op	    1024 B/op	       1 allocs/op
BenchmarkSprint1000-8        	 3000000	       453 ns/op	    1109 B/op	       5 allocs/op
BenchmarkBestSprint1000-8    	 5000000	       317 ns/op	    1040 B/op	       2 allocs/op
BenchmarkSprintf1000-8       	 5000000	       353 ns/op	    1056 B/op	       3 allocs/op
BenchmarkBestSprintf1000-8   	 5000000	       286 ns/op	    1040 B/op	       2 allocs/op
BenchmarkS10000-8            	 1000000	      1023 ns/op	   10240 B/op	       1 allocs/op
PASS
ok  	github.com/tnclong/berry	38.438s
```
