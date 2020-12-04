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

	in := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	n := 1
	for _, input := range in {
		n *= fn0(input[0], input[1], out)
	}
	fmt.Println(n)
}

func fn0(down, right int, out []string) int {
	maxL := len(out[0])
	x, trees := 0, 0
	for i := 0; i < len(out); i += down {
		if out[i][x] == '#' {
			trees++
		}
		x = (x + right) % maxL
	}
	return trees
}
