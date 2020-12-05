package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	t := time.Now()
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

	sort.Slice(rows, func(i, j int) bool {
		return getID(rows[i]) < getID(rows[j])
	})

	last := getID(rows[0])
	for _, row := range rows {
		id := getID(row)
		if id-last > 1 {
			fmt.Println(id, last)
		}
		last = id
	}
	fmt.Println(time.Since(t))
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
