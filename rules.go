package main

// 'n' here means the count of alive cells around a point
func aliveRules(n int) (alive bool) {
	return n == 2 || n == 3

}

func deadRules(n int) (alive bool) {
	return n == 3
}

func countN(prev [][]bool, W, H, x, y int) (n int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 { // don't count the cell itself
				continue
			}
			// y+j and x+i to get the cells around the main cell
			// the % W and % H to wrap around the other side, so if a cell is at the very bottom
			// then y + 1 would equal H, so reminder of H would be 0, which is what we want
			// the +H and +W handles when a cell is at the very start, so 0 => -1 => 2 ( assuming H or W is 3 )
			yAxis := (y + j + H) % H
			xAxis := (x + i + W) % W
			if prev[yAxis][xAxis] {
				n += 1
			}
		}
	}
	return n
}

func (cnf *Config) ApplyRules(x, y int) (alive bool) {
	prev := cnf.history[len(cnf.history)-1]
	n := countN(prev, cnf.W, cnf.H, x, y)
	if cell := prev[y][x]; cell {
		return aliveRules(n)
	} else {
		return deadRules(n)
	}
}
