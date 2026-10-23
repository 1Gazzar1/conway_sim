package main

func getLocationsToCheck(ogSet map[cell]struct{}, H, W int) (locationToCheck map[cell]struct{}) {
	locationToCheck = make(map[cell]struct{})
	// time complex: no. alive cells * 9
	for point := range ogSet {
		// adding the 9 locations around the cell and the cell itself.
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				y := (point.y + j + H) % H
				x := (point.x + i + W) % W
				locationToCheck[cell{x: x, y: y}] = struct{}{}
			}
		}
	}
	return locationToCheck
}

func (config *Config) GenerationsWorker(signal chan struct{}) {
	// send a signal
	signal <- struct{}{}

	for {

		locationsToCheck := getLocationsToCheck(config.aliveLocations, config.H, config.W)

		next := make([][]bool, config.H)
		for i := 0; i < config.H; i++ {
			next[i] = make([]bool, config.W) // zero value is 'false'
		}
		config.aliveLocations = make(map[cell]struct{}) // reset the set

		// calculation the next generation
		for p := range locationsToCheck {
			// if it's alive
			if config.ApplyRules(p.x, p.y) {
				next[p.y][p.x] = true
				config.aliveLocations[cell{x: p.x, y: p.y}] = struct{}{} // populate the set for the next interation

			}
			// fmt.Println(p.x, p.y, r)
		}

		// swapping the current grid with the next one
		config.grid = next
		config.writeToHistory(next)

		// send a signal, meaning you're ahead by 1 more
		signal <- struct{}{}
	}
}
