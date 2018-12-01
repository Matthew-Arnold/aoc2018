package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	currentFrequency := 0
	stdin := bufio.NewScanner(os.Stdin)
	frequencies := make([]int, 0)

	history := make(map[int]bool)

	for stdin.Scan() {
		line := stdin.Text()
		change, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Failed to convert line to int", err)
		}

		frequencies = append(frequencies, change)
	}

	for {
		for _, frequency := range frequencies {
			currentFrequency += frequency

			if _, present := history[currentFrequency]; present {
				fmt.Println(currentFrequency)
				return
			}

			history[currentFrequency] = true
		}
	}
}
