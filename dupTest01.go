package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outPutt := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		outPutt[input.Text()]++
	}
	for line, n := range outPutt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
