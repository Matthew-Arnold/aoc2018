package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Point an x, y cartesian point
type Point struct {
	x int
	y int
}

func manhattanDistance(a Point, b Point) int {
	//why is math.Abs only implemented for float64? who knows
	return int(math.Abs(float64(a.x-b.x))) + int(math.Abs(float64(a.y-b.y)))
}

func minMaxCoords(points []Point) (int, int, int, int) {
	minX := points[0].x
	maxX := points[0].x
	minY := points[0].y
	maxY := points[0].y

	for _, point := range points {
		if point.x < minX {
			minX = point.x
		}

		if point.x > maxX {
			maxX = point.x
		}

		if point.y < minY {
			minY = point.y
		}

		if point.y > maxY {
			maxY = point.y
		}
	}

	return minX, maxX, minY, maxY
}

func contains(s []int, key int) bool {
	for _, val := range s {
		if val == key {
			return true
		}
	}
	return false
}

func main() {
	points := make([]Point, 0)
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		line := stdin.Text()
		var x, y int

		fmt.Sscanf(line, "%d, %d", &x, &y)
		points = append(points, Point{x, y})
	}

	fmt.Println(points)

	_, maxX, _, maxY := minMaxCoords(points)
	grid := make([][]int, maxX+2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, maxY+1)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = -1
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			curPoint := Point{i, j}
			closest := -1
			minDistance := math.MaxInt32

			for num, point := range points {
				distance := manhattanDistance(point, curPoint)
				if distance == 0 {
					closest = -2
					break
				}
				if distance == minDistance {
					closest = -1
				}
				if distance < minDistance {
					closest = num
					minDistance = distance
				}
			}

			grid[i][j] = closest
		}
	}

	infiniteAreas := make([]int, 0)
	//This probably isn't the most efficient, but programmer efficiency!
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if (i == 0 || j == 0) && !contains(infiniteAreas, grid[i][j]) {
				infiniteAreas = append(infiniteAreas, grid[i][j])
			}
		}
	}

	finiteAreas := make([]int, 0)
	for index := range points {
		if !contains(infiniteAreas, index) {
			finiteAreas = append(finiteAreas, index)
		}
	}

	counts := make(map[int]int)
	for _, point := range finiteAreas {
		count := 1 // we're counting the coordinate itself
		for _, row := range grid {
			for _, value := range row {
				if value == point {
					count++
				}
			}
		}

		counts[point] = count
	}

	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	fmt.Println(maxCount)

}
