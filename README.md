## Berry

Berry colorizing printed string on ANSI terminals.
[ANSI escape codes](http://en.wikipedia.org/wiki/ANSI_escape_code)

![Examples](https://github.com/tnclong/berry/blob/master/examples.png)

## Aim

1. Near zero abstraction cost.
2. Test coverage 100%.
3. Easy to use.
4. Good document and examples.

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
	fmt.Println(berry.New(berry.BgYellow).S("BgYellow"))

	berry.Red.Println("FgRed")
	berry.New(berry.BgRed).Println("BgRed")

	fmt.Println(berry.New(berry.Italic).S("Italic"))
	fmt.Println(berry.New(berry.Underline).S("Underline"))

	fmt.Println(berry.New(berry.FgSet, berry.Bit8, 88).S("Fg Bit8"))
	fmt.Println(berry.New(berry.BgSet, berry.Bit8, 88).S("Bg Bit8"))

	pFg := berry.New(berry.FgSet, berry.Bit24, 100, 100, 1)
	pFg.Println("Fg Bit24")
	pBg := berry.New(berry.BgSet, berry.Bit24, 100, 100, 1)
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
BenchmarkS1-8                        	30000000	        51.3 ns/op	      16 B/op	       1 allocs/op
BenchmarkS10-8                       	30000000	        54.0 ns/op	      32 B/op	       1 allocs/op
BenchmarkS50-8                       	20000000	        59.9 ns/op	      64 B/op	       1 allocs/op
BenchmarkS100-8                      	20000000	        68.5 ns/op	     112 B/op	       1 allocs/op
BenchmarkSS100-8                     	 3000000	       433 ns/op	     360 B/op	       5 allocs/op
BenchmarkBestS100-8                  	20000000	        60.8 ns/op	     112 B/op	       1 allocs/op
BenchmarkSprint100-8                 	 5000000	       249 ns/op	     176 B/op	       3 allocs/op
BenchmarkBestSprintUseBerry100-8     	10000000	       169 ns/op	     128 B/op	       2 allocs/op
BenchmarkBestSprint100-8             	10000000	       168 ns/op	     128 B/op	       2 allocs/op
BenchmarkSprintf100-8                	10000000	       200 ns/op	     144 B/op	       3 allocs/op
BenchmarkBestSprintfUseBerry100-8    	10000000	       169 ns/op	     128 B/op	       2 allocs/op
BenchmarkBestSprintf100-8            	10000000	       139 ns/op	     128 B/op	       2 allocs/op
BenchmarkS500-8                      	10000000	       119 ns/op	     512 B/op	       1 allocs/op
BenchmarkS1000-8                     	10000000	       190 ns/op	    1024 B/op	       1 allocs/op
BenchmarkSS1000-8                    	 2000000	       831 ns/op	    3096 B/op	       5 allocs/op
BenchmarkBestS1000-8                 	10000000	       181 ns/op	    1024 B/op	       1 allocs/op
BenchmarkSprint1000-8                	 5000000	       390 ns/op	    1088 B/op	       3 allocs/op
BenchmarkBestSprintUseBerry1000-8    	 5000000	       314 ns/op	    1040 B/op	       2 allocs/op
BenchmarkBestSprint1000-8            	 5000000	       311 ns/op	    1040 B/op	       2 allocs/op
BenchmarkSprintf1000-8               	 5000000	       346 ns/op	    1056 B/op	       3 allocs/op
BenchmarkBestSprintfUseBerry1000-8   	 5000000	       310 ns/op	    1040 B/op	       2 allocs/op
BenchmarkBestSprintf1000-8           	 5000000	       277 ns/op	    1040 B/op	       2 allocs/op
BenchmarkS10000-8                    	 1000000	      1016 ns/op	   10240 B/op	       1 allocs/op
PASS
ok  	github.com/tnclong/berry	40.922s
```
