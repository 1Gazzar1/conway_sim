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

	return grid, aliveLocs
}

func NewConfig(W, H int, grid [][]bool, initPos map[cell]struct{}) Config {
	return Config{aliveLocations: initPos, history: [][][]bool{grid}, W: W, H: H, mu: &sync.Mutex{}}
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
