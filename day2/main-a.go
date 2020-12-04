package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func maina() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	ct := 0
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var min, max int
		var letter, str string
		_, err = fmt.Sscanf(strings.TrimSpace(st), "%d-%d %s %s", &min, &max, &letter, &str)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		n := 0
		ch := rune(letter[0])
		for _, ch1 := range str {
			if ch1 == ch {
				n++
			}
		}
		if min <= n && n <= max {
			ct++
		}
	}
	fmt.Println(ct)
}
