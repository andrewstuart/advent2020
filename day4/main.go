package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	hcl = regexp.MustCompile("^#[0-9a-f]{6}$")
	ecl = regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
	pid = regexp.MustCompile(`^\d{9}$`)
)

var req = map[string]func(string) bool{
	"byr": func(s string) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return 1920 <= n && n <= 2002
	},
	"iyr": func(s string) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return 2010 <= n && n <= 2020
	},
	"eyr": func(s string) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return 2020 <= n && n <= 2030
	},
	"hgt": func(s string) bool {
		switch {
		case strings.HasSuffix(s, "cm"):
			n, err := strconv.Atoi(strings.TrimSuffix(s, "cm"))
			if err != nil {
				return false
			}
			return 150 <= n && n <= 193
		case strings.HasSuffix(s, "in"):
			n, err := strconv.Atoi(strings.TrimSuffix(s, "in"))
			if err != nil {
				return false
			}
			return 59 <= n && n <= 76
		}
		return false
	},
	"hcl": func(s string) bool {
		return hcl.MatchString(s)
	},
	"ecl": func(s string) bool {
		return ecl.MatchString(s)
	},
	"pid": func(s string) bool {
		return pid.MatchString(s)
	},
}

func main() {
	n := time.Now()
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	next := map[string]string{}

	valid := 0
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			if check(next) {
				valid++
			}
			break
		}
		if strings.TrimSpace(st) == "" {
			if check(next) {
				valid++
			}
			next = map[string]string{}
			continue
		}
		for _, f := range strings.Fields(st) {
			fs := strings.Split(f, ":")
			next[strings.TrimSpace(fs[0])] = strings.TrimSpace(fs[1])
		}
	}

	fmt.Println(valid)
	fmt.Println(time.Since(n))
}

func check(next map[string]string) bool {
	for r, fn := range req {
		if val, ok := next[r]; !ok || !fn(val) {
			return false
		}
	}
	return true
}
