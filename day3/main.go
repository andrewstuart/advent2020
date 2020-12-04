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

	in := []string{}
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		in = append(in, strings.TrimSpace(st))
	}

	slopes := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	n := 1
	for _, slope := range slopes {
		n *= countTrees(slope[0], slope[1], in)
	}
	fmt.Println(n)
}

func countTrees(down, right int, input []string) int {
	maxL := len(input[0])
	x, trees := 0, 0
	for i := 0; i < len(input); i += down {
		if input[i][x] == '#' {
			trees++
		}
		x = (x + right) % maxL
	}
	return trees
}
