package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var reqA = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	"cid",
}

func maina() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	passports := []map[string]string{}
	next := map[string]string{}

	valid := 0
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			if checkA(next) {
				valid++
			}
			passports = append(passports, next)
			break
		}
		if strings.TrimSpace(st) == "" {
			if check(next) {
				valid++
			}
			passports = append(passports, next)
			next = map[string]string{}
			continue
		}
		for _, f := range strings.Fields(st) {
			fs := strings.Split(f, ":")
			next[fs[0]] = fs[1]
		}
	}

	fmt.Println(valid)
}
func checkA(next map[string]string) bool {
	for _, r := range reqA[:len(req)-1] { // skip cid
		if _, ok := next[r]; !ok {
			return false
		}
	}
	return true
}
