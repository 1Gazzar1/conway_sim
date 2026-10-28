package main
var PresetMap map[string][][]bool = map[string][][]bool{
	"blinker": Pad(blinker, 2, 1),
	"square":  Pad(square, 1, 1),
	"glider":  Pad(glider, 1, 1),
	"lwss":    Pad(lwss, 3, 2),
	"toad":    Pad(toad, 1, 2),
	"beacon":  Pad(beacon, 1, 1),
	"pulsar":  Pad(pulsar, 3, 3),
	"diehard": Pad(diehard, 8, 18),
	"ggg":     Pad(gosperGliderGun, 1, 6),
}

var blinker = [][]bool{
	[]bool{true},
	[]bool{true},
	[]bool{true},
}

var square = [][]bool{
	[]bool{true, true},
	[]bool{true, true},
}

var glider = [][]bool{
	[]bool{false, true, false},
	[]bool{false, false, true},
	[]bool{true, true, true},
}

var lwss = [][]bool{
	[]bool{false, true, true, true, true},
	[]bool{true, false, false, false, true},
	[]bool{false, false, false, false, true},
	[]bool{true, false, false, true, false},
}

var toad = [][]bool{
	[]bool{false, true, true, true},
	[]bool{true, true, true, false},
}

var beacon = [][]bool{
	[]bool{true, true, false, false},
	[]bool{true, true, false, false},
	[]bool{false, false, true, true},
	[]bool{false, false, true, true},
}

var pulsar = [][]bool{
	[]bool{false, false, true, true, true, false, false, false, true, true, true, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{false, false, true, true, true, false, false, false, true, true, true, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, true, true, true, false, false, false, true, true, true, false, false},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{true, false, false, false, false, true, false, true, false, false, false, false, true},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, true, true, true, false, false, false, true, true, true, false, false},
}

var diehard = [][]bool{
	[]bool{false, false, false, false, false, false, true, false},
	[]bool{true, true, false, false, false, false, false, false},
	[]bool{false, true, false, false, false, true, true, true},
}

var gosperGliderGun = [][]bool{
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	[]bool{true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	[]bool{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
}

func getInitPresetAliveLocsSet(preset [][]bool) (set map[cell]struct{}) {
	set = make(map[cell]struct{}, 0)
	for i, row := range preset {
		for j, val := range row {
			if val {
				set[cell{x: j, y: i}] = struct{}{}
			}
		}
	}
	return set
}

// Pad returns a copy of p with n blank cells added on every side.
func Pad(p [][]bool, x, y int) [][]bool {
	h, w := len(p), len(p[0])
	out := make([][]bool, h+2*y)
	for s := range out {
		out[s] = make([]bool, w+2*x)
	}
	for s := range p {
		copy(out[s+y][x:], p[s])
	}
	return out
}
