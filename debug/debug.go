package debug

import "fmt"

func Debug(vs ...any) {
	for _, v := range vs {
		fmt.Printf("%v\n", v)
	}
}
