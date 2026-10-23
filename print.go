package main

import "fmt"

func (cnf *Config) PrintGrid(deadValue, aliveValue string, frame [][]bool) {
	// Source - https://stackoverflow.com/a/22892171
	// Posted by Kavu, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-10-21, License - CC BY-SA 4.0

	fmt.Print("\033[H\033[3J")
	fmt.Print("\033[H")

	for _, row := range frame {
		for _, f := range row {
			fmt.Print("")
			if f {
				fmt.Print(string(aliveValue))
			} else {
				fmt.Print(string(deadValue))
			}
		}
		fmt.Println()
	}
}
