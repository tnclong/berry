package x11

import (
	"testing"

	"github.com/tnclong/berry"
)

func TestX11(t *testing.T) {
	berry.Enable(true)

	actual := Aqua.S("Aqua")
	want := "\x1b[38;2;0;255;255mAqua\x1b[0m"
	t.Log(want, actual)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}
