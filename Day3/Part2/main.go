package main

import (
	"bufio"
	"fmt"
	"os"
)

type coordinate struct {
	x int
	y int
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	locations := make(map[int][]coordinate)

	for stdin.Scan() {
		line := stdin.Text()
		var id, x, y, width, height int

		_, _ = fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)
		//fmt.Printf("ID: %d -- x: %d -- y: %d -- width: %d -- height: %d\n", id, x, y, width, height)

		current, present := locations[id]
		if !present {
			current = make([]coordinate, 0)
		}

		for i := x; i < x+width; i++ {
			for j := y; j < y+height; j++ {

				current = append(current, coordinate{x: i, y: j})
			}
		}

		locations[id] = current
	}

	for id, coordinates := range locations {
		found := true

		for otherID, otherCoords := range locations {
			if id != otherID {
				for _, coord := range coordinates {
					for _, otherCoord := range otherCoords {
						if coord == otherCoord {
							//fmt.Printf("%v == %v\n", coord, otherCoord)
							found = false
						}
					}
				}
			}
		}

		if found {
			fmt.Println(id)
			return
		}
	}
}
