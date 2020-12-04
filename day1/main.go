package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func maina() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	seen := map[int]struct{}{}
	for {
		st, err := br.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.Atoi(strings.TrimSpace(st))
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := seen[2020-num]; ok {
			fmt.Println(num, 2020-num)
			fmt.Println(num * (2020 - num))
		}
		seen[num] = struct{}{}
	}
}
