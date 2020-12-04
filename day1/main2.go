package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile("input", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	var seen []int
	for {
		st, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.Atoi(strings.TrimSpace(st))
		if err != nil {
			log.Fatal(err)
		}
		seen = append(seen, num)
	}
	sort.Ints(seen)

	out := search(3, 2020, 0, seen, count(seen))

	fmt.Printf("out = %+v\n", out)
	n := 1
	for _, n1 := range out {
		n *= n1
	}
	fmt.Println(n)
}

func count(ints []int) map[int]int {
	out := map[int]int{}
	for _, n := range ints {
		out[n]++
	}
	return out
}

func search(iters int, tgt, soFar int, nums []int, countTrack map[int]int) []int {
	if iters <= 0 {
		if soFar == tgt {
			return []int{}
		}
		return nil
	}

	maxInd := sort.SearchInts(nums, tgt-soFar) // can't be any number past the target - total so far
	nums = nums[:maxInd+1]
	for _, n := range nums {
		if countTrack[n] < 1 {
			continue
		}
		if soFar+n > tgt {
			return nil
		}
		countTrack[n]--
		next := search(iters-1, tgt, soFar+n, nums, countTrack)
		countTrack[n]++
		if next == nil {
			continue
		}
		return append(next, n)
	}
	return nil
}
