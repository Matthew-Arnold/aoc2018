package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	twoCount := 0
	threeCount := 0

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		line := stdin.Text()
		counts := make(map[rune]int)

		for _, char := range line {
			current, present := counts[char]

			if present {
				counts[char] = current + 1
			} else {
				counts[char] = 1
			}
		}

		//fmt.Printf("%s: %v\n", line, counts)

		hasDouble := false
		hasTriple := false

		for _, count := range counts {
			if count == 2 {
				hasDouble = true
			}

			if count == 3 {
				hasTriple = true
			}
		}

		if hasDouble {
			twoCount++
		}

		if hasTriple {
			threeCount++
		}
	}

	checksum := twoCount * threeCount
	fmt.Println(checksum)
}
