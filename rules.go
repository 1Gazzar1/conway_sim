package main

// 'n' here means the count of alive cells around a point
func aliveRules(n int) (alive bool) {
	if n >= 4 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	return false

}

func deadRules(n int) (alive bool) {
	if n == 3 {
		return true
	}
	return false
}

func (cnf *config) countN(x, y int) (n int) {
	H := len(cnf.grid)
	W := len(cnf.grid[0])

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
			if cnf.grid[yAxis][xAxis] {
				n += 1
			}
		}
	}
	return n
}

func (cnf *config) ApplyRules(x, y int) (alive bool) {
	n := cnf.countN(x, y)
	if cell := cnf.grid[y][x]; cell {
		return aliveRules(n)
	} else {
		return deadRules(n)
	}
}
