package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m string
	q := 0

	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &m)

		if len(m) == 1 {
			fmt.Fprintln(out, "0")
			continue
		}

		s := []rune(m)

		r := false

		for j := 0; j < len(s); j++ {
			if j < len(s)-1 {
				if !r && s[j+1] > s[j] {
					j++
					r = true
				}

				fmt.Fprint(out, string(s[j]))
			} else {
				if r {
					fmt.Fprint(out, string(s[j]))
				}
			}
		}

		fmt.Fprint(out, "\n")
	}
}
