package berry

import (
	"testing"
)

func TestColor(t *testing.T) {
	Enable(true)

	actual := Green.S("Green")
	want := "\x1b[32mGreen\x1b[0m"
	t.Log(want, actual)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}
