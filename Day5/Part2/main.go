package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func sameLetter(a rune, b rune) bool {
	return unicode.ToUpper(a) == unicode.ToUpper(b)
}

func differentCase(a rune, b rune) bool {
	return (unicode.IsUpper(a) && unicode.IsLower(b)) || (unicode.IsLower(a) && unicode.IsUpper(b))
}

func reacts(a rune, b rune) bool {
	return differentCase(a, b) && sameLetter(a, b)
}

func fullyReact(polymer string) string {
	changed := true
	for changed {
		changed = false
		_, lastSize := utf8.DecodeLastRuneInString(polymer)

		for i, w := 0, 0; i < len(polymer)-lastSize; i += w {
			a, aWidth := utf8.DecodeRuneInString(polymer[i:])
			w = aWidth
			b, bWidth := utf8.DecodeRuneInString(polymer[i+w:])

			if reacts(a, b) {
				//fmt.Println("Reaction: ", string(a), string(b))
				//fmt.Println("Was: ", polymer)
				polymer = polymer[0:i] + polymer[i+aWidth+bWidth:]
				//fmt.Println("Polymer now ", polymer)
				changed = true
			}
		}
	}

	return polymer
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	polymer, err := stdin.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	polymer = polymer[:len(polymer)-1]

	theFuckingAlphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	minLength := utf8.RuneCountInString(polymer)

	for _, char := range theFuckingAlphabet {
		testPoly := strings.Replace(polymer, char, "", -1)
		testPoly = strings.Replace(testPoly, strings.ToUpper(char), "", -1)

		testPoly = fullyReact(testPoly)
		length := utf8.RuneCountInString(testPoly)

		if length < minLength {
			minLength = length
		}
	}

	fmt.Println(minLength)
}
