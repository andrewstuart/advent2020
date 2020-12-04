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

		ch := letter[0]
		if str[min-1] == ch && str[max-1] == ch {
			continue
		}
		if !(str[min-1] == ch || str[max-1] == ch) {
			continue
		}
		ct++
	}
	fmt.Println(ct)
}
