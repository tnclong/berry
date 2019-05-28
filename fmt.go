package berry

import (
	"fmt"
	"io"
)

// Simple wrap std fmt package
// add SGR parameters to head and tail of out string
// detail of how to use these method, please see fmt package directly
// https://golang.org/pkg/fmt/

func (r R) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, r.S(format), a...)
}

func (r R) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(r.S(format), a...)
}

func (r R) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(r.S(format), a...)
}

func (r R) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, r.around(a)...)
}

func (r R) Print(a ...interface{}) (n int, err error) {
	return fmt.Print(r.around(a)...)
}

func (r R) Sprint(a ...interface{}) string {
	return fmt.Sprint(r.around(a)...)
}

func (r R) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, r.around(a)...)
}

func (r R) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(r.around(a)...)
}

func (r R) Sprintln(a ...interface{}) string {
	return fmt.Sprintln(r.around(a)...)
}

func (r R) around(a []interface{}) []interface{} {
	if !enabled {
		return a
	}

	aa := make([]interface{}, len(a)+2)
	aa[0] = join(r)
	copy(aa[1:], a)
	aa[len(a)+1] = tseq
	return aa
}
