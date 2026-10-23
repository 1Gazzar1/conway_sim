package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	grid     [][]bool
	c        *Config
	speed    float64
	interval time.Duration
	signals  chan struct{}
	curGen   int
	maxGen   int
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

func initialModel(c *Config, ch chan struct{}) model {
	frame, found := c.getFrame(1)
	if !found {
		log.Fatal("couldn't get frame from config history")
	}
	m := model{
		speed:   3.0,
		paused:  false,
		grid:    frame,
		curGen:  1,
		c:       c,
		signals: ch,
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
		case "right":
			m.paused = true
			frame, found := m.c.getFrame(m.curGen + 1)
			if found {
				m.grid = frame
				m.curGen += 1
			}

			if m.curGen > m.maxGen {
				m.maxGen = m.curGen
				// recieve a signal, representing a slot has opened from the channel queue
				// so generator can generate more
				<-m.signals
			}
		case "left":
			m.paused = true
			frame, found := m.c.getFrame(m.curGen - 1)
			if found {
				m.grid = frame
				m.curGen -= 1
			}
		}

	case tickMsg:
		if !m.paused {
			frame, found := m.c.getFrame(m.curGen + 1)
			if found {
				m.grid = frame
				m.curGen += 1
			}

			if m.curGen > m.maxGen {
				m.maxGen = m.curGen
				// recieve a signal, representing a slot has opened from the channel queue
				// so generator can generate more
				<-m.signals
			}

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
	var b strings.Builder
	b.Grow(m.c.W*m.c.H*3 + 128) // each emoji is 3 bytes in UTF-8
	for _, row := range m.grid {
		for _, alive := range row {
			if alive {
				b.WriteString("⬜")
			} else {
				b.WriteString("⬛")
			}
		}
		b.WriteByte('\n')
	}

	p := fmt.Sprint(m.speed)
	if m.paused {
		p = "PAUSED"
	}

	fmt.Fprintf(&b, "\nGeneration: %v · Speed: %v (generations/s) \n[Space] Pause · [+/-] Speed · [q] Quit", m.curGen, p)

	// Send the UI for rendering
	return tea.NewView(b.String())
}
func main() {
	W, H := 80, 30
	// interval := 300 * time.Millisecond
	config := NewConfig(W, H)
	signals := make(chan struct{}, 100) // 100 buffred channel

	go config.GenerationsWorker(signals)

	p := tea.NewProgram(initialModel(&config, signals))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// for frame := range frames {
	// 	config.PrintGrid("⬛", "⬜", frame)
	// 	time.Sleep(interval)
	// }
}
