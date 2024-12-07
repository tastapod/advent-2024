package debug

import (
	"fmt"
	"os"
	"strings"
)

func Debug(vs ...any) {
	out := make([]string, len(vs))

	for i, v := range vs {
		out[i] = fmt.Sprintf("%v", v)
	}
	_, _ = fmt.Fprintln(os.Stderr, strings.Join(out, " "))
}

func Debugf(msg string, vs ...any) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", vs...)
}
