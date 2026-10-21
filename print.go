package main

import "fmt"

func (cnf *config) PrintGrid(deadValue, aliveValue string) {
	for _, row := range cnf.grid {
		for _, f := range row {
			fmt.Print(" ")
			if f {
				fmt.Print(string(aliveValue))
			} else {
				fmt.Print(string(deadValue))
			}
		}
		fmt.Println()
	}
}
