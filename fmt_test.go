package berry

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestAround(t *testing.T) {
	Enable(false)

	arr := []interface{}{1, "1", true}
	yellow := Yellow

	actual := yellow.around(arr)
	if !reflect.DeepEqual(actual, arr) {
		t.Errorf("want same")
	}

	Enable(true)
	actual = yellow.around(arr)
	if len(actual) != len(arr)+2 {
		t.Errorf("want size increase 2")
	}
	if actual[0] != yellow {
		t.Errorf("want first is %q but get %q", string(yellow), actual[0])
	}
	if !reflect.DeepEqual(actual[1:len(actual)-1], arr) {
		t.Errorf("want same")
	}
	if actual[len(actual)-1] != tseq {
		t.Errorf("want first is %q but get %q", tseq, actual[len(actual)-1])
	}
}

func TestFprintf(t *testing.T) {
	Enable(true)

	green := Green
	var buf bytes.Buffer
	_, err := green.Fprintf(&buf, "g%vg", 1)
	if err != nil {
		t.Fatal(err)
	}
	actual := buf.String()
	want := "\x1b[32mg1g\x1b[0m"
	t.Log(want, actual)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func TestPrintf(t *testing.T) {
	Enable(true)

	stdout(
		t,
		func() {
			green := Green
			_, err := green.Printf("g%vg", 2)
			if err != nil {
				t.Fatal(err)
			}
		},
		func(actual string) {
			want := "\x1b[32mg2g\x1b[0m"
			t.Log(want, actual)
			if actual != want {
				t.Errorf("want %q but get %q", want, actual)
			}
		},
	)
}

func TestSprintf(t *testing.T) {
	Enable(true)

	green := Green
	actual := green.Sprintf("g%vg", 3)
	want := "\x1b[32mg3g\x1b[0m"
	t.Log(actual, want)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func TestFprint(t *testing.T) {
	Enable(true)

	green := Green
	var buf bytes.Buffer
	_, err := green.Fprint(&buf, "g", 4, "g")
	if err != nil {
		t.Fatal(err)
	}
	actual := buf.String()
	want := "\x1b[32mg4g\x1b[0m"
	t.Log(want, actual)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func TestPrint(t *testing.T) {
	Enable(true)

	stdout(
		t,
		func() {
			green := Green
			_, err := green.Print("g", 5, "g")
			if err != nil {
				t.Fatal(err)
			}
		},
		func(actual string) {
			want := "\x1b[32mg5g\x1b[0m"
			t.Log(want, actual)
			if actual != want {
				t.Errorf("want %q but get %q", want, actual)
			}
		},
	)
}

func TestSprint(t *testing.T) {
	Enable(true)

	green := Green
	actual := green.Sprint("g", 6, "g")
	want := "\x1b[32mg6g\x1b[0m"
	t.Log(actual, want)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func TestFprintln(t *testing.T) {
	green := New(BgGreen)
	var buf bytes.Buffer
	_, err := green.Fprintln(&buf, "g", 7, "g")
	if err != nil {
		t.Fatal(err)
	}
	actual := buf.String()
	want := "\x1b[42m g 7 g \x1b[0m\n"
	t.Log(want, actual)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

func TestPrintln(t *testing.T) {
	stdout(
		t,
		func() {
			green := Green
			_, err := green.Println("g", 8, "g")
			if err != nil {
				t.Fatal(err)
			}
		},
		func(actual string) {
			want := "\x1b[32m g 8 g \x1b[0m\n"
			t.Log(want, actual)
			if actual != want {
				t.Errorf("want %q but get %q", want, actual)
			}
		},
	)
}

func TestSprintln(t *testing.T) {
	Enable(true)

	green := Green
	actual := green.Sprintln("g", 9, "g")
	want := "\x1b[32m g 9 g \x1b[0m\n"
	t.Log(actual, want)
	if actual != want {
		t.Errorf("want %q but get %q", want, actual)
	}
}

var stdoutLock = new(sync.Mutex)

func stdout(t *testing.T, wFunc func(), assertFunc func(string)) {
	stdoutLock.Lock()
	defer stdoutLock.Unlock()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	out := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = out
	}()

	wFunc()
	w.Close()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	assertFunc(string(data))
}
