package debug

import (
	"fmt"
	"strings"
)

func Debug(vs ...any) {
	out := make([]string, len(vs))

	for i, v := range vs {
		out[i] = fmt.Sprintf("%v", v)
	}
	println(strings.Join(out, " | "))
}

func Debugf(msg string, vs ...any) {
	fmt.Printf(msg+"\n", vs...)
}
