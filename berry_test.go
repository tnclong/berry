package berry

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestS(t *testing.T) {
	var cases = []struct {
		str  string
		r    R
		r2   R
		want string
	}{
		{
			// when no codes(nil) given
			str:  "\x1b[1mhello\x1b[0m",
			r:    New(),
			want: "hello",
		},
		{
			str:  "Reset",
			r:    New(Reset),
			want: "\x1b[0mReset\x1b[0m",
		},
		{
			str:  "Bright",
			r:    New(Bright),
			want: "\x1b[1mBright\x1b[0m",
		},
		{
			str:  "Faint",
			r:    New(Faint),
			want: "\x1b[2mFaint\x1b[0m",
		},
		{
			str:  "Italic",
			r:    New(Italic),
			want: "\x1b[3mItalic\x1b[0m",
		},
		{
			str:  "Underline",
			r:    New(Underline),
			want: "\x1b[4mUnderline\x1b[0m",
		},
		{
			str:  "Blink",
			r:    New(Blink),
			want: "\x1b[5mBlink\x1b[0m",
		},
		{
			str:  "Inverse",
			r:    New(Inverse),
			want: "\x1b[7mInverse\x1b[0m",
		},
		{
			str:  "Hide",
			r:    New(Hide),
			want: "\x1b[8mHide\x1b[0m",
		},
		{
			str:  "CrossOut",
			r:    New(CrossOut),
			want: "\x1b[9mCrossOut\x1b[0m",
		},
		{
			str:  "FgBlack",
			r:    New(FgBlack),
			want: "\x1b[30mFgBlack\x1b[0m",
		},
		{
			str:  "BgRed",
			r:    New(BgRed),
			want: "\x1b[41mBgRed\x1b[0m",
		},
		{
			str:  "FgSet, Bit8, 51",
			r:    New(FgSet, Bit8, 51),
			want: "\x1b[38;5;51mFgSet, Bit8, 51\x1b[0m",
		},
		{
			str:  "BgSet, Bit8, 160",
			r:    New(BgSet, Bit8, 160),
			want: "\x1b[48;5;160mBgSet, Bit8, 160\x1b[0m",
		},
		{
			str:  "FgSet, Bit24, 0, 0, 0",
			r:    New(FgSet, Bit24, 0, 0, 0),
			want: "\x1b[38;2;0;0;0mFgSet, Bit24, 0, 0, 0\x1b[0m",
		},
		{
			str:  "BgSet, Bit24, 255, 255, 255",
			r:    New(BgSet, Bit24, 255, 255, 255),
			want: "\x1b[48;2;255;255;255mBgSet, Bit24, 255, 255, 255\x1b[0m",
		},

		{
			str:  "Italic then BgRed",
			r:    New(Italic),
			r2:   New(BgRed),
			want: "\x1b[41m\x1b[3mItalic then BgRed\x1b[0m\x1b[0m",
		},
	}

	Enable(true)
	for _, tc := range cases {
		actual := tc.r.S(tc.str)
		if len(tc.r2.r) != 0 {
			actual = tc.r2.S(actual)
		}
		t.Log(tc.want, actual)
		if actual != tc.want {
			t.Errorf("S(%q) want %q but get %q", tc.str, tc.want, actual)
		}
	}

	Enable(false)
	for _, tc := range cases {
		actual := tc.r.S(tc.str)
		if len(tc.r2.r) != 0 {
			actual = tc.r2.S(actual)
		}
		if actual != tc.str {
			t.Errorf("when Enable(false), Effect(%q) want %q but get %q", tc.str, tc.str, actual)
		}
	}
}

func TestSS(t *testing.T) {
	Enable(true)

	var cases = []struct {
		str  string
		r    R
		r2   R
		want string
	}{
		{
			str:  "Italic then BgRed",
			r:    New(Italic),
			r2:   New(BgRed),
			want: "\x1b[3m\x1b[41mItalic then BgRed\x1b[0m",
		},
	}

	Enable(true)
	for _, tc := range cases {
		actual := tc.r.SS(tc.str)
		if len(tc.r2.r) != 0 {
			actual = tc.r2.SS(actual)
		}
		t.Log(tc.want, actual)
		if actual != tc.want {
			t.Errorf("S(%q) want %q but get %q", tc.str, tc.want, actual)
		}
	}
}

func TestSDisplay(t *testing.T) {
	Enable(true)

	var dyes = map[string]uint8{
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
		t.Log(New(d).S(str))
	}

	var bit8Fg, bit8Bg [32][8]string
	var c uint8
	for {
		bit8Fg[c/8][c%8] = New(FgSet, Bit8, c).S(fmt.Sprintf("FgSet: %v", c))
		bit8Bg[c/8][c%8] = New(BgSet, Bit8, c).S(fmt.Sprintf("BgSet: %v", c))

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

			bit24Fg[gg*32+bb/8][bb%8] = New(FgSet, Bit24, r, g, b).S(fmt.Sprintf("%v,%v,%v", r, g, b))
			bit24Bg[gg*32+bb/8][bb%8] = New(BgSet, Bit24, r, g, b).S(fmt.Sprintf("%v,%v,%v", r, g, b))

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
		t.Log(fmt.Sprintf("%q", bit24Fg[i]))
		t.Log(bit24Bg[i])
		t.Log(fmt.Sprintf("%q", bit24Bg[i]))
	}

	t.Log(New(Italic, Underline, FgRed).S("multi1"))
	t.Log(New(FgSet, Bit8, 201, BgSet, Bit8, 46).S("multi2"))
	t.Log(New(Italic, Underline, FgSet, Bit8, 201, BgSet, Bit8, 46).S("multi3"))
}

func TestNew(t *testing.T) {
	r := New(FgSet, Bit8, 1)
	pr := New([]uint8(r.R())...)

	want := "\x1b[38;5;1m"
	t.Logf("%q %q", want, pr)
	if want != string(pr.r) {
		t.Logf("want %q but get %q", want, pr)
	}

	for i := 0; i < 10; i++ {
		want := "\x1b[38;5;1ma\x1b[0m"
		actual := pr.S("a")
		t.Log(want, actual)
		if want != actual {
			t.Errorf("want %q but get %q", want, actual)
		}
	}

	ppr := New([]uint8(pr.R())...)
	if want != string(ppr.R()) {
		t.Errorf("want %q but get %q", want, string(ppr.R()))
	}
}

func TestR(t *testing.T) {
	rs := Yellow.R()
	want := "\x1b[33m"
	if rs != want {
		t.Errorf("want %q but get %q", want, rs)
	}

	t.Log(rs + "y" + RReset.R())
}

func TestRI(t *testing.T) {
	rsi := Yellow.RI()
	rsi2 := Yellow.RI()
	if rsi != rsi2 {
		t.Error("want same")
	}

	_, ok := rsi.(string)
	if !ok {
		t.Error("want rsi type is string")
	}

	actual := fmt.Sprint(rsi)
	want := "\x1b[33m"
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func BenchmarkS1(b *testing.B) {
	benchmarkS(b, Green, 1)
}

func BenchmarkS10(b *testing.B) {
	benchmarkS(b, Green, 10)
}

func BenchmarkS50(b *testing.B) {
	benchmarkS(b, Green, 50)
}

func BenchmarkS100(b *testing.B) {
	benchmarkS(b, Green, 100)
}

func BenchmarkSS100(b *testing.B) {
	benchmarkSS(b, Green, 100)
}

func BenchmarkBestS100(b *testing.B) {
	benchmarkBS(b, 100)
}

func BenchmarkSprint100(b *testing.B) {
	benchmarkSprint(b, Green, 100)
}

func BenchmarkBestSprintUseBerry100(b *testing.B) {
	benchmarkBestSprintUseBerry(b, Green, 100)
}

func BenchmarkBestSprint100(b *testing.B) {
	benchmarkBSprint(b, 100)
}

func BenchmarkSprintf100(b *testing.B) {
	benchmarkSprintf(b, Green, 100)
}

func BenchmarkBestSprintfUseBerry100(b *testing.B) {
	benchmarkBestSprintfUseBerry(b, Green, 100)
}

func BenchmarkBestSprintf100(b *testing.B) {
	benchmarkBSprintf(b, 100)
}

func BenchmarkS500(b *testing.B) {
	benchmarkS(b, Green, 500)
}

func BenchmarkS1000(b *testing.B) {
	benchmarkS(b, Green, 1000)
}

func BenchmarkSS1000(b *testing.B) {
	benchmarkSS(b, Green, 1000)
}

func BenchmarkBestS1000(b *testing.B) {
	benchmarkBS(b, 1000)
}

func BenchmarkSprint1000(b *testing.B) {
	benchmarkSprint(b, Green, 1000)
}

func BenchmarkBestSprintUseBerry1000(b *testing.B) {
	benchmarkBestSprintUseBerry(b, Green, 1000)
}

func BenchmarkBestSprint1000(b *testing.B) {
	benchmarkBSprint(b, 1000)
}

func BenchmarkSprintf1000(b *testing.B) {
	benchmarkSprintf(b, Green, 1000)
}

func BenchmarkBestSprintfUseBerry1000(b *testing.B) {
	benchmarkBestSprintfUseBerry(b, Green, 1000)
}

func BenchmarkBestSprintf1000(b *testing.B) {
	benchmarkBSprintf(b, 1000)
}

func BenchmarkS10000(b *testing.B) {
	benchmarkS(b, Green, 10000)
}

func benchmarkS(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.S(str)
	}
}

func benchmarkSS(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.SS(str)
	}
}

// func benchmarkBS(b *testing.B, count int) {
// 	str := strings.Repeat("1", count)
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		strings.Join([]string{"\x1b[32m", str, "\x1b[0m"}, "")
// 	}
// }

func benchmarkBS(b *testing.B, count int) {
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = "\x1b[32m" + str + "\x1b[0m"
	}
}

// func benchmarkBS(b *testing.B, count int) {
// 	str := strings.Repeat("1", count)
// 	prev := []byte("\x1b[32m")
// 	tail := []byte("\x1b[0m")
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		var buf = make([]byte, 5+len(str)+4)
// 		n := 0
// 		copy(buf[n:], prev)
// 		n += 5
// 		copy(buf[n:], []byte(str))
// 		n += len(str)
// 		copy(buf[n:], tail)
// 		_ = string(buf)
// 	}
// }

func benchmarkSprint(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.Sprint(str)
	}
}

func benchmarkBSprint(b *testing.B, count int) {
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fmt.Sprint("\x1b[32m", str, "\x1b[0m")
	}
}

func benchmarkBestSprintUseBerry(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fmt.Sprint(r.RI(), str, RReset.RI())
	}
}

func benchmarkSprintf(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		r.Sprintf("%s", str)
	}
}

func benchmarkBSprintf(b *testing.B, count int) {
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fmt.Sprintf("\x1b[32m%s\x1b[0m", str)
	}
}

func benchmarkBestSprintfUseBerry(b *testing.B, r R, count int) {
	Enable(true)
	str := strings.Repeat("1", count)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%s%s%s", r.RI(), str, RReset.RI())
	}
}
