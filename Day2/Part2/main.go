package main

import (
	"bufio"
	"fmt"
	"os"
)

func differentLetters(s1 string, s2 string) int {
	count := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}

	return count
}

func commonLetters(s1 string, s2 string) string {
	result := ""

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			result += string(s1[i])
		}
	}

	return result
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	boxIds := make([]string, 0)

	for stdin.Scan() {
		line := stdin.Text()
		boxIds = append(boxIds, line)
	}

	for _, id := range boxIds {
		for _, otherID := range boxIds {
			if differentLetters(id, otherID) == 1 {
				common := commonLetters(id, otherID)
				fmt.Println(common)
				return
			}
		}
	}
}
