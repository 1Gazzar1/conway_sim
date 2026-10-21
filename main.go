package main

import (
	"time"
)

type cell struct {
	x, y int
}
type config struct {
	grid           [][]bool
	aliveLocations map[cell]struct{} // set of (x,y) tuples
	history        [][][]bool
}

func main() {
	W, H := 5, 5
	g, s := genRandomGrid(W, H)
	config := config{
		grid:           g,
		aliveLocations: s,
		history:        [][][]bool{g}, // initial grid
	}
	config.PrintGrid("❌", "🤡")
	for {
		time.Sleep(1 * time.Second)

		locationToCheck := make(map[cell]struct{})

		// time complex: no. alive cells * 9
		for point := range config.aliveLocations {
			// adding the 9 locations around the cell and the cell itself.
			for i := point.x - 1; i <= point.x+1; i++ {
				for j := point.y - 1; j <= point.y+1; j++ {
					y := (j + H) % H
					x := (i + W) % W
					locationToCheck[cell{x: x, y: y}] = struct{}{}
				}
			}
		}

		next := make([][]bool, H)
		for i := 0; i < H; i++ {
			next[i] = make([]bool, W)
		}

		// calculation the next generation
		for p := range locationToCheck {
			next[p.y][p.x] = config.ApplyRules(p.x, p.y)

			// fmt.Println(p.x, p.y, r)
		}

		// swapping the current grid with the next one
		config.grid = next
		config.history = append(config.history, next)

		config.PrintGrid("❌", "🤡")

	}

}

func genRandomGrid(W, H int) (grid [][]bool, aliveLocs map[cell]struct{}) {

	// apparently this is how you init a matrix
	// grid = make([][]bool, H)
	// for i := 0; i < H; i++ {
	// 	grid[i] = make([]bool, W)
	// }

	// aliveLocs = make(map[cell]struct{})

	// for i := range W {
	// 	for j := range H {
	// 		if r := rand.Float64(); r > 0.5 {
	// 			grid[j][i] = true
	// 			aliveLocs[cell{x: i, y: j}] = struct{}{}
	// 		} else {
	// 			grid[j][i] = false
	// 		}

	// 	}
	// }
	grid = [][]bool{
		{false, false, false, false, false}, // y = 0
		{false, false, false, false, false}, // y = 1
		{false, true, true, true, false},    // y = 2 (Left: 1,2 | Center: 2,2 | Right: 3,2)
		{false, false, false, false, false}, // y = 3
		{false, false, false, false, false}, // y = 4
	}

	// 2. Initialize the set with the (x, y) tuples directly using a literal statement
	aliveLocs = map[cell]struct{}{
		{x: 1, y: 2}: {},
		{x: 2, y: 2}: {},
		{x: 3, y: 2}: {},
	}
	return grid, aliveLocs
}
