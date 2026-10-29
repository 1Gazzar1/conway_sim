package main

var PresetMap map[string][][]bool = map[string][][]bool{
	"blinker": Pad(blinker, 2, 1),
	"square":  Pad(square, 1, 1),
	"glider":  Pad(glider, 1, 1),
	"lwss":    Pad(lwss, 3, 2),
	"toad":    Pad(toad, 1, 2),
	"beacon":  Pad(beacon, 1, 1),
	"pulsar":  Pad(pulsar, 3, 3),
	"diehard": Pad(diehard, 8, 11),
	"ggg":     Pad(gosperGliderGun, 1, 6),
}

var blinker = [][]bool{
	{true},
	{true},
	{true},
}

var square = [][]bool{
	{true, true},
	{true, true},
}

var glider = [][]bool{
	{false, true, false},
	{false, false, true},
	{true, true, true},
}

var lwss = [][]bool{
	{false, true, true, true, true},
	{true, false, false, false, true},
	{false, false, false, false, true},
	{true, false, false, true, false},
}

var toad = [][]bool{
	{false, true, true, true},
	{true, true, true, false},
}

var beacon = [][]bool{
	{true, true, false, false},
	{true, true, false, false},
	{false, false, true, true},
	{false, false, true, true},
}

var pulsar = [][]bool{
	{false, false, true, true, true, false, false, false, true, true, true, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{false, false, true, true, true, false, false, false, true, true, true, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, true, true, true, false, false, false, true, true, true, false, false},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{true, false, false, false, false, true, false, true, false, false, false, false, true},
	{false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, true, true, true, false, false, false, true, true, true, false, false},
}

var diehard = [][]bool{
	{false, false, false, false, false, false, true, false},
	{true, true, false, false, false, false, false, false},
	{false, true, false, false, false, true, true, true},
}

var gosperGliderGun = [][]bool{
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
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
