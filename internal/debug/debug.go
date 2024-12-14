package debug

import (
	"fmt"
	"os"
	"strings"
)

func Debug(vs ...any) {
	debug(" ", vs...)
}

func Debugln(vs ...any) {
	debug("\n", vs...)
}

func debug(join string, vs ...any) {
	out := make([]string, len(vs))

	for i, v := range vs {
		out[i] = fmt.Sprintf("%v", v)
	}
	_, _ = fmt.Fprintln(os.Stderr, strings.Join(out, join))
}

func Debugf(msg string, vs ...any) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", vs...)
}
