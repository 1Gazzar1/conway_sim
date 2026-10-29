package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

func main() {
	// user requested Width/Height
	W := flag.Int("w", 30, "width of the grid")
	H := flag.Int("h", 20, "height of the grid")
	a := flag.String("a", "⬜", "value of alive cell in the grid")
	d := flag.String("d", "⬛", "value of dead cell in the grid")
	preset := flag.String("preset", "", `starting pattern, could be one of the following: 
	- blinker
	- square
	- glider
	- lwss
	- toad
	- beacon
	- pulsar
	- diehard
	- ggg  `)

	flag.Parse()

	if *W < 3 || *H < 3 {
		fmt.Println("Your Width and Height have to be atleast 3.")
		os.Exit(1)
	}
	// height is the same whether it's ASCII or UTF-8 (emoji)
	// width differes, an emoji is wider than a english language letter
	// look "xx"
	// look "🤡"
	lenA := runewidth.StringWidth(*a)
	lenD := runewidth.StringWidth(*d)

	if lenA == 0 || lenD == 0 {
		fmt.Println("Enter valid characters and not empty ones!")
		os.Exit(1)
	}
	if lenA != lenD {
		fmt.Println("Please use characters of the same width, e.g. '⬜' and '⬛' or '.' and 'x' ")
		os.Exit(1)
	}
	// terminal width/height
	tW, tH, err := term.GetSize(int(os.Stdout.Fd())) // the width and height of the terminal used

	if err != nil {
		fmt.Printf("something went wrong, couldn't get terminal size, err: %v", err)
		os.Exit(1)
	}

	g := make([][]bool, 0)
	s := make(map[cell]struct{}, 0)
	// final width/height
	h, w := 0, 0

	if grid, exists := PresetMap[*preset]; exists {
		g = grid
		s = getInitPresetAliveLocsSet(grid)
		h = len(grid)
		w = len(grid[0])
	} else {
		g, s = genRandomGrid(*W, *H)
		w, h = *W, *H
	}

	if w*lenA > tW || h+3 > tH {
		fmt.Printf("There's isn't enough space to show the grid, chose smaller w,h or smaller a,d or make your terminal bigger!")
		os.Exit(1)
	}
	wMargin := (tW - w*lenA) / 2 // normal amount of size so like this: " ", not 2 characters wide

	config := NewConfig(w, h, g, s)
	signals := make(chan struct{}, 100) // 100 buffred channel

	go config.GenerationsWorker(signals)

	p := tea.NewProgram(initialModel(&config, signals, cellTypes{
		dead:  *d,
		alive: *a,
	}, wMargin))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
