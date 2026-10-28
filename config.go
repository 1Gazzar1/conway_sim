package main

import (
	"math/rand"
	"sync"
)

type cell struct {
	x, y int
}
type Config struct {
	aliveLocations map[cell]struct{} // set of (x,y) tuples
	history        [][][]bool
	W, H           int
	mu             *sync.Mutex
	stIndex        int
}

func genRandomGrid(W, H int) (grid [][]bool, aliveLocs map[cell]struct{}) {

	// apparently this is how you init a matrix
	grid = make([][]bool, H)
	for i := 0; i < H; i++ {
		grid[i] = make([]bool, W)
	}

	aliveLocs = make(map[cell]struct{})

	for i := range W {
		for j := range H {
			if r := rand.Float64(); r > 0.8 {
				grid[j][i] = true
				aliveLocs[cell{x: i, y: j}] = struct{}{}
			} else {
				grid[j][i] = false
			}

		}
	}
	// grid = [][]bool{
	// 	{false, false, false, false, false}, // y = 0
	// 	{false, false, false, false, false}, // y = 1
	// 	{false, true, true, true, false},    // y = 2 (Left: 1,2 | Center: 2,2 | Right: 3,2)
	// 	{false, false, false, false, false}, // y = 3
	// 	{false, false, false, false, false}, // y = 4
	// }

	// // 2. Initialize the set with the (x, y) tuples directly using a literal statement
	// aliveLocs = map[cell]struct{}{
	// 	{x: 1, y: 2}: {},
	// 	{x: 2, y: 2}: {},
	// 	{x: 3, y: 2}: {},
	// }
	return grid, aliveLocs
}

func NewConfig(W, H int, preset string) Config {
	g := make([][]bool, 0)
	s := make(map[cell]struct{}, 0)
	w, h := 0, 0
	if grid, exists := PresetMap[preset]; exists {
		g = grid
		s = getInitPresetAliveLocsSet(grid)
		h = len(grid)
		w = len(grid[0])
	} else {
		g, s = genRandomGrid(W, H)
		w, h = W, H
	}
	return Config{aliveLocations: s, history: [][][]bool{g}, W: w, H: h, mu: &sync.Mutex{}}
}

func (c *Config) getFrame(gen int) (frame [][]bool, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if gen <= c.stIndex || gen > len(c.history)+c.stIndex {
		return frame, false
	}
	return c.history[gen-1-c.stIndex], true
}
func (c *Config) writeToHistory(newFrame [][]bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.history = append(c.history, newFrame)

	n := len(c.history)
	if n > 200 {
		removed := n - 200
		c.history = c.history[removed:]
		c.stIndex += removed
	}
}
