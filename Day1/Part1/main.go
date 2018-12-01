package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		line := stdin.Text()
		change, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Failed to convert line to int", err)
		}

		frequency += change
	}

	fmt.Println(frequency)
}
