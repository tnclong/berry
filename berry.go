package berry

// Effect wraps SGR parameter arround the str
// detail of e, please see type of E
func Effect(str string, e E) string {
	if !enabled {
		return str
	}

	return SGR(str, uint8(e))
}

// Dye wraps background color or foreground color arround the str.
//
// foreground color is a type D and start with Fg,
// background color is a type D and start with Bg.
//
// When d is FgRGB or BgRGB, a legal c is required.
//   if length of c is 1, we will treat as 8-bit(256-color),
//   if length of c is 3, we will treat as 24-bit(r,g,b),
//   other lengths of c are ignore silently
func Dye(str string, d D, c ...uint8) string {
	if !enabled {
		return str
	}

	if d != FgRGB && d != BgRGB {
		return SGR(str, uint8(d))
	}

	if len(c) == 1 {
		return SGR(str, uint8(d), 5, c[0])
	}

	if len(c) == 3 {
		return SGR(str, uint8(d), 2, c[0], c[1], c[2])
	}

	return str
}

// E is subset of SGR parameters.
// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
type E uint8

const (
	// Reset terminal to default colors/backgrounds(attributes)
	// It shouldn't be needed to use this because all methods
	// append reset code to end of string.
	Reset E = iota
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

// D is color of text
// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
type D uint8

const (
	// 3/4-bit foreground color of text
	FgBlack D = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// 8-bit or 24-bit foreground color of text
	FgRGB

	// 3/4-bit background color of text
	BgBlack = iota + 31
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// 8-bit or 24-bit background color of text
	BgRGB
)
