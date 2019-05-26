package berry

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEffect(t *testing.T) {
	var cases = []struct {
		str  string
		e    E
		want string
	}{
		{
			str:  "Reset",
			e:    Reset,
			want: "\x1b[0mReset\x1b[0m",
		},
		{
			str:  "Bright",
			e:    Reset,
			want: "\x1b[0mBright\x1b[0m",
		},
		{
			str:  "Faint",
			e:    Faint,
			want: "\x1b[2mFaint\x1b[0m",
		},
		{
			str:  "Italic",
			e:    Italic,
			want: "\x1b[3mItalic\x1b[0m",
		},
		{
			str:  "Underline",
			e:    Underline,
			want: "\x1b[4mUnderline\x1b[0m",
		},
		{
			str:  "Blink",
			e:    Blink,
			want: "\x1b[5mBlink\x1b[0m",
		},
		{
			str:  "Inverse",
			e:    Inverse,
			want: "\x1b[7mInverse\x1b[0m",
		},
		{
			str:  "Hide",
			e:    Hide,
			want: "\x1b[8mHide\x1b[0m",
		},
		{
			str:  "CrossOut",
			e:    CrossOut,
			want: "\x1b[9mCrossOut\x1b[0m",
		},
	}

	for _, tc := range cases {
		actual := Effect(tc.str, tc.e)
		t.Log(tc.want, actual)
		if actual != tc.want {
			t.Errorf("Effect(%q) want %q but get %q", tc.str, tc.want, actual)
		}
	}
}

func TestDye(t *testing.T) {
	var cases = []struct {
		str  string
		d    D
		c    []uint8
		want string
	}{
		{
			str:  "FgBlack",
			d:    FgBlack,
			want: "\x1b[30mFgBlack\x1b[0m",
		},
		{
			str:  "BgRed",
			d:    BgRed,
			want: "\x1b[41mBgRed\x1b[0m",
		},
		{
			str:  "FgRGB 51",
			d:    FgRGB,
			c:    []uint8{51},
			want: "\x1b[38;5;51mFgRGB 51\x1b[0m",
		},
		{
			str:  "BgRGB 160",
			d:    BgRGB,
			c:    []uint8{160},
			want: "\x1b[48;5;160mBgRGB 160\x1b[0m",
		},
		{
			str:  "FgRGB 0,0,0",
			d:    FgRGB,
			c:    []uint8{0, 0, 0},
			want: "\x1b[38;2;0;0;0mFgRGB 0,0,0\x1b[0m",
		},
		{
			str:  "BgRGB 255,255,255",
			d:    BgRGB,
			c:    []uint8{255, 255, 255},
			want: "\x1b[48;2;255;255;255mBgRGB 255,255,255\x1b[0m",
		},
	}
	for _, tc := range cases {
		actual := Dye(tc.str, tc.d, tc.c...)
		t.Log(tc.want, actual)
		if actual != tc.want {
			t.Errorf("Dye(%q) want %q but get %q", tc.str, tc.want, actual)
		}
	}
}

func TestDyeDisplay(t *testing.T) {
	var dyes = map[string]D{
		"FgBlack":   FgBlack,
		"FgRed":     FgRed,
		"FgGreen":   FgGreen,
		"FgYellow":  FgYellow,
		"FgBlue":    FgBlue,
		"FgMagenta": FgMagenta,
		"FgCyan":    FgCyan,
		"FgWhite":   FgWhite,

		"BgBlack":   BgBlack,
		"BgRed":     BgRed,
		"BgGreen":   BgGreen,
		"BgYellow":  BgYellow,
		"BgBlue":    BgBlue,
		"BgMagenta": BgMagenta,
		"BgCyan":    BgCyan,
		"BgWhite":   BgWhite,
	}
	for str, d := range dyes {
		t.Log(Dye(str, d), d)
	}

	var bit8Fg, bit8Bg [32][8]string
	var c uint8
	for {
		bit8Fg[c/8][c%8] = Dye(fmt.Sprintf("FgRGB: %v", c), FgRGB, c)
		bit8Bg[c/8][c%8] = Dye(fmt.Sprintf("BgRGB: %v", c), BgRGB, c)

		if c^0xFF == 0 {
			break
		}
		c++
	}
	for i := 0; i < 32; i++ {
		t.Log(bit8Fg[i])
		t.Log(bit8Bg[i])
	}

	var bit24Fg, bit24Bg [8192][8]string
	rand.Seed(time.Now().UnixNano())
	r := uint8(rand.Intn(256))
	var g, b uint8
	for {
		gg := uint32(g)
		for {
			bb := uint32(b)

			bit24Fg[gg*32+bb/8][bb%8] = Dye(fmt.Sprintf("%v,%v,%v", r, g, b), FgRGB, r, g, b)
			bit24Bg[gg*32+bb/8][bb%8] = Dye(fmt.Sprintf("%v,%v,%v", r, g, b), BgRGB, r, g, b)

			if b^0xFF == 0 {
				b = 0
				break
			}
			b++
		}

		if g^0xFF == 0 {
			g = 0
			break
		}
		g++
	}
	for i := 0; i < 8192; i++ {
		t.Log(bit24Fg[i])
		t.Log(bit24Bg[i])
	}
}
