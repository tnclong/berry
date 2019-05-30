package berry

import (
	"regexp"
	"strings"
)

// R is a sequence of SGR parameters and Colors.
//
// When the length of R is 0, the S will clear all surrounding in str.
//
// This implement reference to:
//   https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
//   https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
//
// See exmaples in:
//   https://github.com/tnclong/berry#usage
//   exmaples/*.go
//   *_test.go
//
type R struct {
	r string

	// https://github.com/golang/go/issues/32305#issuecomment-497051905
	ri interface{}
}

// New create a new R
//
// If r already matched `^(\x1b\[([\d;]+)m)`, original r will be returned.
//
// examples:
//   r := berry.New(berry.FgSet, berry.Bit8, 1) => "\x1b[38;5;1m"
//   r.S("s") => "\x1b[38;5;1ms\x1b[0m"
func New(r ...uint8) R {
	rr := string(join(r))
	return R{r: rr, ri: rr}
}

// S wraps str around a sequence of SGR parameters that store in r.
//
// When the length of R is 0, the S will clear all surrounding in str.
func (r R) S(str string) string {
	return r.s(str, false)
}

// SS is strict S method.
//
// This method is about 4~6x slower than S in benchmark result.
//
// If you have a str already arrounded with SGR:
//     Red.S("\x1b[3mItalic then BgRed\x1b[0m")
//       => "\x1b[41m\x1b[3mItalic then BgRed\x1b[0m\x1b[0m"
//
//     Red.SS("\x1b[3mItalic then BgRed\x1b[0m")
//       => "\x1b[3m\x1b[41mItalic then BgRed\x1b[0m"
func (r R) SS(str string) string {
	return r.s(str, true)
}

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
	return r.ri
}

func (r R) s(str string, strict bool) string {
	if !enabled {
		return str
	}

	if len(r.r) == 0 {
		return seqReg.ReplaceAllString(str, "")
	}

	if !strict {
		return r.r + str + tseq
	}

	str = hseqReg.ReplaceAllStringFunc(str, func(m string) string {
		return m + r.r
	})
	if strings.HasSuffix(str, tseq) {
		return str
	}
	return str + tseq
}

var (
	seqReg  = regexp.MustCompile(`\x1b\[[0-9;]*m`)
	hseqReg = regexp.MustCompile(`^(\x1b\[([\d;]+)m)*`)
	mseqReg = regexp.MustCompile(`^(\x1b\[([\d;]+)m)`)
)

const (
	tseq = "\x1b[0m"
)

// A subset of SGR parameters that be used in New.
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
const (
	// Reset terminal to default colors/backgrounds(attributes)
	// It shouldn't be needed to use this because all methods
	// append reset code to end of string.
	Reset uint8 = iota
	Bright
	Faint
	Italic
	Underline
	Blink
	_
	Inverse
	Hide
	CrossOut
)

// 3/4-bit foreground color of text that be used in New.
const (
	FgBlack uint8 = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// 8-bit or 24-bit foreground color of text, see Bit8 and Bit24
	FgSet
)

// 3/4-bit background color of text that be used in New.
const (
	BgBlack uint8 = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// 8-bit or 24-bit background color of text, see  Bit8 and Bit24
	BgSet
)

const (
	// Bit8 is a flag that be used after FgSet and BgSet for specify a 8-bit color.
	// examples:
	//    berry.New(berry.FgSet, berry.Bit8, 1).S("Red") => "\x1b[38;5;1mRed\x1b[0m"
	//    berry.New(berry.BgSet, berry.Bit8, 11).S("Yellow") => "\x1b[48;5;11mYellow\x1b[0m"
	Bit8 uint8 = 5
	// Bit24 is a flag that be used after FgSet and BgSet for specify a 24-bit color.
	// examples:
	//    berry.New(berry.FgSet, berry.Bit24, 0, 0, 0).S("Black") => "\x1b[38;2;0;0;0mBlack\x1b[0m"
	//    berry.New(berry.BgSet, berry.Bit24, 0, 0, 0).S("Black") => "\x1b[48;2;0;0;0mBlack\x1b[0m"
	Bit24 uint8 = 2

	// Color256 is alias of Bit8
	Color256 = Bit8
	// RGB is alias of bit24
	RGB = Bit24
	// TrueColor is alias of bit24
	TrueColor = Bit24
)
