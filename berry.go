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
type R []uint8

// Prepare create a new R has better performace
// If r already prepared, r will be returned.
// examples:
//   string(Prepare(berry.R{berry.FgSet, berry.Bit8, 1})) => "\x1b[38;5;1m"
func Prepare(r R) R {
	return join(r)
}

// S wraps str around a sequence of SGR parameters that store in r.
//
// When the length of R is 0, the S will clear all surrounding in str.
func (r R) S(str string) string {
	return r.s(str, false)
}

// SS is strict S method.
//
// This method is about 7x slower than S in benchmark result.
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

func (r R) s(str string, strict bool) string {
	if !enabled {
		return str
	}

	if len(r) == 0 {
		return seqReg.ReplaceAllString(str, "")
	}

	if !strict {
		return string(join(r)) + str + tseq
	}

	hseq := string(join(r))
	str = hseqReg.ReplaceAllStringFunc(str, func(m string) string {
		return m + hseq
	})
	if strings.HasSuffix(str, tseq) {
		return str
	}
	return str + tseq
}

var (
	seqReg  = regexp.MustCompile(`\x1b\[[0-9;]*m`)
	hseqReg = regexp.MustCompile(`^(\x1b\[([\d;]+)m)*`)
)

const (
	tseq = "\x1b[0m"
)

// A subset of SGR parameters.
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

// 3/4-bit foreground color of text
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

// 3/4-bit background color of text
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
	//    berry.R{berry.FgSet, berry.Bit8, 1}.S("Red") => "\x1b[38;5;1mRed\x1b[0m"
	//    berry.R{berry.BgSet, berry.Bit8, 11}.S("Yellow") => "\x1b[48;5;11mYellow\x1b[0m"
	Bit8 uint8 = 5
	// Bit24 is a flag that be used after FgSet and BgSet for specify a 24-bit color.
	// examples:
	//    berry.R{berry.FgSet, berry.Bit24, 0, 0, 0}.S("Black") => "\x1b[38;2;0;0;0mBlack\x1b[0m"
	//    berry.R{berry.BgSet, berry.Bit24, 0, 0, 0}.S("Black") => "\x1b[48;2;0;0;0mBlack\x1b[0m"
	Bit24 uint8 = 2

	// Color256 is alias of Bit8
	Color256 = Bit8
	// RGB is alias of bit24
	RGB = Bit24
	// TrueColor is alias of bit24
	TrueColor = Bit24
)
