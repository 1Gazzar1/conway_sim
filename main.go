package main

import (
	"time"
)

func main() {
	W, H := 50, 20
	interval := 300 * time.Millisecond
	config := NewConfig(W, H)
	frames := make(chan [][]bool, 100)

	go config.GenerationsWorker(frames)

	for frame := range frames {
		config.PrintGrid("⬛", "⬜", frame)
		time.Sleep(interval)
	}
}
