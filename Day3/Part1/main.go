package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid [1000][1000]int
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		line := stdin.Text()
		var id, x, y, width, height int

		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)
		//fmt.Printf("ID: %d -- x: %d -- y: %d -- width: %d -- height: %d\n", id, x, y, width, height)

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {
				//fmt.Println(i, j)
				grid[i][j]++
			}
		}
	}

	overlaps := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}
