package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	stdin := bufio.NewReader(os.Stdin)
	polymer, err := stdin.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	polymer = polymer[:len(polymer)-1]

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

	//fmt.Println(polymer)
	fmt.Println(utf8.RuneCountInString(polymer))
}
