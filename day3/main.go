package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	out := []string{}
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		out = append(out, strings.TrimSpace(st))
	}

	down, right := 1, 3

	maxL := len(out[0])
	x, trees := 0, 0
	for i := 0; i < len(out); i += down {
		if out[i][x] == '#' {
			trees++
		}
		x = (x + right) % maxL
	}
	fmt.Println(trees)
}
