# My Experience Of Re-inventing Wheel Berry In Golang



# 0. Introduction

## What is Berry?

> Berry colorizing printed string on ANSI terminals

```go
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
```

Copy above code block to a `main.go` file, then `go run main.go`. You will see colorful in your terminal!:heart_eyes: :heart_eyes: 
![Examples](https://raw.githubusercontent.com/tnclong/berry/master/examples.png)

## What is it implements?

>  // This implement reference to:
>  //   https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
>  //   https://en.wikipedia.org/wiki/ANSI_escape_code#Colors

## Where you can find it?

https://github.com/tnclong/berry

# 1. Re-inventing wheel is good.

## Author not maintained his package. https://github.com/fatih/color

His [blog](https://arslan.io/2018/10/09/taking-an-indefinite-sabbatical-from-my-projects/) said:  

> “It Doesn’t Have to Be Crazy at Work”, I decided to make a big change in my life.

## Mysterious strings.

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	color := "39"
	level := "error"
	err := errors.New("Wow! a red error message")
	switch level {
	case "crit":
		color = "35"
	case "error":
		color = "31"
	case "warn":
		color = "33"
	case "info":
		color = "32"
	}

	fmt.Printf("\x1b[%sm%v\x1b[0m\n", color, err)
}

```


# 2. First Step: Create a Github repertory[free].

## Give a name to your baby.

I love ruby.
I find a ruby project. [rainbow 彩虹](https://github.com/sickill/rainbow)
So something colorful? 
See my desktop, strawberry(草莓)! 
Mmm... Only red and blue, Mmm... Berry(浆果)! Yeah!

You are Rainbow, I'm Berry.
You in sky, I'm on land.
You are more beautiful than me, I'm more delicious than you.
You write in ruby, I'm in golang.

## Choose a LICENSE

[MIT](https://github.com/tnclong/berry/blob/master/LICENSE)


# 3. Initiate [implement](https://github.com/tnclong/berry/commit/c48ce76385089ec442a376eb6d58182c8c2f2138) according to standard.

## Mistaken

Divisive `Dye`, `Effect` and `SGR`.
```go
func Dye(str string, d D, c ...uint8) string {}
func Effect(str string, e E) string {}
func SGR(str string, codes ...uint8) string {}
```

Divisive `D` and `E`.
```go
type D uint8
type E uint8
```

## Write test, document and examples.

```
// Dye wraps background color or foreground color arround the str.
//
// foreground color is a type D and start with Bg,
// background color is a type D and start with Bg.
//
// When d is FgRGB or BgRGB, a legal c is required.
//   if length of c is 1, we will treat as 8-bit(256-color),
//   if length of c is 3, we will treat as 24-bit(r,g,b),
//   other lengths of c are ignore silently
func Dye(str string, d D, c ...uint8) string {}
func TestEffect(t *testing.T) {}
func TestDye(t *testing.T) {}
func TestSGR(t *testing.T) {}
// examples/main.go
// README.md
```


## Study from [std library](https://github.com/golang/go/blob/master/src/strconv/itoa.go#L68).

Nothing faster than table search!
[My table search](https://github.com/tnclong/berry/commit/c48ce76385089ec442a376eb6d58182c8c2f2138#diff-a39bdb9731f408fcb3fd66d63e5b3d22R58)
```go
// small returns the string for an i with 0 <= i < nSmalls.
func small(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}

const nSmalls = 100

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"
```

# 4. Write more test cases, examples and document.

```ruby
("d0d1e8cab5c4eac5ee041e941f9d36b45c66221f".."8773b0814a7529688c37ed66260e0b2fa0eaea2c")
```


## Eat your strawberry.

Write test! Write example! write document.
It is foot stone to redesign.
It is a initial of your mind that you make some mistaken.
You eat your strawberry. The strawberry is too sour to the taste.

## An interesting example

See `Introduction`. Based from https://github.com/fatih/color#color--

# 5. [Refactor](https://github.com/tnclong/berry/commit/a6ccb0d09ad348e7e0c06948fcfc85d4062568b1)

## All in `R`.

```go
type R []uint8
```

# 6. User Friendly API.

## Near to std library [fmt](https://github.com/golang/go/blob/master/src/fmt/print.go#L217).

https://github.com/tnclong/berry/commit/c870cd2f79c58652fa3c8bf8abf0a1cdaa56644f
```go
func (r R) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {}
func (r R) Printf(format string, a ...interface{}) (n int, err error) {}
func (r R) Sprintf(format string, a ...interface{}) string {}
func (r R) Fprint(w io.Writer, a ...interface{}) (n int, err error) {}
func (r R) Print(a ...interface{}) (n int, err error) {}
func (r R) Sprint(a ...interface{}) string {}
func (r R) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {}
func (r R) Println(a ...interface{}) (n int, err error) {}
func (r R) Sprintln(a ...interface{}) string {}
```

## Add [x11 color names](https://github.com/tnclong/berry/commit/1082885350d720b45d5c1844772222e68be3c06b).

```
package x11
// https://en.wikipedia.org/wiki/X11_color_names#Color_name_chart
var (
	Aqua              = prepare(0, 255, 255)
	Aquamarine        = prepare(127, 255, 212)
	MediumAquamarine  = prepare(102, 205, 170)
	Azure             = prepare(240, 255, 255)

	// Not show all
)
```

## Add [basic colors](https://github.com/tnclong/berry/commit/e28543e65abc26a00cc7a1e929c73877e65f2501).

```go
package berry

// convenient basic foreground colors, useful in many cases.
var (
	Black   = Prepare(R{FgBlack})
	Red     = Prepare(R{FgRed})
	Green   = Prepare(R{FgGreen})
	Yellow  = Prepare(R{FgYellow})
	Blue    = Prepare(R{FgBlue})
	Magenta = Prepare(R{FgMagenta})
	Cyan    = Prepare(R{FgCyan})
	White   = Prepare(R{FgWhite})
)
```

# 7. Performance(zero cost abstraction)

## Add [benchmarks](https://github.com/tnclong/berry/commit/5168190c062e6caf54ae2d8c0d03557110e4eb39)

```
func BenchmarkS100(b *testing.B) {}
func BenchmarkS100WithoutP(b *testing.B) {}
func BenchmarkBestS100(b *testing.B) {}
// Not show more
```

## High performance at first.

Add new method `SS()`. Default method `S()` has higher performance.
https://github.com/tnclong/berry/commit/a37a2b1b48d32a3ebda29ec8b7d111596f5a5a44
```
BenchmarkS100-8              	20000000	        79.3 ns/op	     112 B/op	       1 allocs/op
BenchmarkSS100-8             	 3000000	       466 ns/op	     368 B/op	       6 allocs/op
```

## `bytes.HasPrefix()` has same performance.

https://github.com/tnclong/berry/commit/f0d4e5fc00dc988680ec3a87208a698614c86378
```go
import (
	"bytes"
)

var sseq = []uint8{'\x1b', '['}

func join(codes []uint8) []uint8 {
	// if len(codes) > 3 &&
	// 	codes[0] == '\x1b' && codes[1] == '[' &&
	// 	codes[len(codes)-1] == 'm' {
	if bytes.HasPrefix(codes, sseq) {
	}
}
```

## `const` `string` will reduce a alloc op when assign to `interface{}`

```go
// var(
// 	tseq    = "\x1b[0m"
// )

const (
	tseq = "\x1b[0m"
)
```

## `type R []uint8` to `type R string`

Reduce `string([]unit8{})` cost.

[make R as string that fast 10+ ns/op](https://github.com/tnclong/berry/commit/fd86c7199b47174084154007fd8d87f111f17ba0)

```
BenchmarkS1-8                	20000000	        57.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkS10-8               	20000000	        61.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkS50-8               	20000000	        67.1 ns/op	      64 B/op	       1 allocs/op
BenchmarkS100-8              	20000000	        79.3 ns/op	     112 B/op	       1 allocs/op
BenchmarkS500-8              	10000000	       134 ns/op	     512 B/op	       1 allocs/op
BenchmarkS1000-8             	10000000	       195 ns/op	    1024 B/op	       1 allocs/op
BenchmarkS10000-8            	 1000000	      1023 ns/op	   10240 B/op	       1 allocs/op
```
```
BenchmarkS1-8                	300000000	        49.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkS10-8               	300000000	        52.4 ns/op	      32 B/op	       1 allocs/op
BenchmarkS50-8               	300000000	        58.2 ns/op	      64 B/op	       1 allocs/op
BenchmarkS100-8              	200000000	        66.2 ns/op	     112 B/op	       1 allocs/op
BenchmarkS500-8              	100000000	       119 ns/op	     512 B/op	       1 allocs/op
BenchmarkS1000-8             	100000000	       203 ns/op	    1024 B/op	       1 allocs/op
BenchmarkS10000-8            	20000000	      1039 ns/op	   10240 B/op	       1 allocs/op
```

## `type R string` to 

```go
type R struct {
	r string

	// https://github.com/golang/go/issues/32305#issuecomment-497051905
	ir interface{}
}
```

Reduce `string` assign to `interface{}` cost.
https://github.com/tnclong/berry/commit/1baeb14e9edbd2f08f32e25513599f0c580eef15
```go
// 	aa[0] = r
	aa[0] = r.ir
```

```
BenchmarkSprint100-8                 	 5000000	       249 ns/op	     176 B/op	       3 allocs/op
```
```
BenchmarkSprint100-8         	50000000	       324 ns/op	     192 B/op	       4 allocs/op
```

## New methods `RI()` and `R()` for you need higher performance

```
// R return a sequence of SGR parameters that able to append to str directly.
//   berry.Red.R() + "red" + berry.RRset.R() => "\x1b[31mred\x1b[0m"
//
// use this method when you need higher performance.
//
func (r R) R() string {
	return r.r
}

// RI is a interface{} that contains R() value.
// this method avoid a alloc op when use R as interface{} in golang.
//
// example:
//     fmt.Sprint(berry.Blue.RI(), "1", 2, berry.RRset.RI())
//
//     aa := make([]interface{}, 2)
//     aa[0] = berry.Yellow.RI()
//
// use this method when you need higher performance.
//
// you can see more detail in https://github.com/golang/go/issues/32305#issuecomment-497051905
func (r R) RI() interface{} {
	return r.ir
}
```

# 8. Write a stroy share to you.

# 9. Summary.

We consider correctness, ease to use, performance when we build a new wheel.
We have emphasis in different phase, but we must keep them in our mind.
We use great tools to implement our goal.
`go` language is your friend.
`go test` is your friend.
`go cover` is your friend.
`go doc` is your friend.
`go benchmark` is your friend.


# 10. THOUGHT

Zero-cost abstraction is impossible?
Re-inventing wheel is awesome! forwarding!
