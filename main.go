package main

import (
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	grid     [][]bool
	speed    float64
	interval time.Duration
	ch       chan [][]bool
	paused   bool
}

type tickMsg time.Time

func (m *model) updateInterval() {
	m.interval = time.Duration(int(float64(1.0/m.speed)*1000)) * time.Millisecond
}
func (m *model) updateSpeed(speed float64) {
	if speed > 25 || speed <= 0 {
		return
	}
	m.speed = speed
	m.updateInterval()
}

func initialModel(ch chan [][]bool) model {
	m := model{
		speed:  3.0,
		paused: false,
		grid:   <-ch,
		ch:     ch,
	}
	m.updateInterval()
	return m
}

func (m model) Init() tea.Cmd {
	return tea.Tick(m.interval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		case "=":
			m.updateSpeed(m.speed + 1)
		case "-":
			m.updateSpeed(m.speed - 1)
		case "space":
			m.paused = !m.paused
		}

	case tickMsg:
		if !m.paused {
			m.grid = <-m.ch
		}
		return m, tea.Tick(m.interval, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
func (m model) View() tea.View {
	s := ""
	for _, row := range m.grid {
		for _, f := range row {
			if f {
				s += "⬜"
			} else {
				s += "⬛"
			}
		}
		s += "\n"
	}

	// The footer

	p := fmt.Sprint(m.speed)
	if m.paused {
		p = "PAUSED"
	}

	s += fmt.Sprintf("\nSpeed: %v (generations/s) \n[Space] Pause · [+/-] Speed · [q] Quit", p)

	// Send the UI for rendering
	return tea.NewView(s)
}
func main() {
	W, H := 50, 20
	// interval := 300 * time.Millisecond
	config := NewConfig(W, H)
	frames := make(chan [][]bool, 100)

	go config.GenerationsWorker(frames)

	p := tea.NewProgram(initialModel(frames))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// for frame := range frames {
	// 	config.PrintGrid("⬛", "⬜", frame)
	// 	time.Sleep(interval)
	// }
}
