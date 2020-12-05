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

	rows := []string{}
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		rows = append(rows, strings.TrimSpace(st))
	}

	test := "FBFBBFFRLR"
	fmt.Println(get(test, 7, 'B'))

	max := 0
	for _, row := range rows {
		id := getID(row)
		if id > max {
			max = id
		}
	}

	fmt.Println(max)
}

func get(row string, num int, char byte) int {
	var n int
	for i := 0; i < num; i++ {
		if row[i] == char {
			n |= (1 << (num - i - 1))
		}
	}
	return n
}

func getID(row string) int {
	return 8*get(row, 7, 'B') + get(row[7:], 3, 'R')
}
