package main

import (
	"fmt"
	"math/rand"
	"time"

	// "godown/geom"
	"godown/particles"
	"godown/simulation"

	"github.com/gdamore/tcell/v2"
)

func main() {
    screen, err := tcell.NewScreen()
    if err != nil {
        panic(err)
    }

    err = screen.Init()
    if err != nil {
        panic(err)
    }
    defer screen.Fini()

    sim := simulation.NewSimulation(screen)

    s := particles.NewSand(1, 1)

    sim.AddParticle(s)

    quit := make(chan struct{})

    start := time.Now()

    go func() {
        ticker := time.NewTicker(time.Second / 10)
        defer ticker.Stop()

        loop:
        for {
            select {
            case <-ticker.C:

                if len(sim.Particles) > 3 * sim.W {
                    clearBottomRow(sim)
                }

                status := []rune(fmt.Sprintf("sim.H = %d | sim.W = %d | len(sim.Particles) = %d | timePassed = %f", sim.H, sim.W, len(sim.Particles), time.Now().Sub(start).Seconds()))
                randX := rand.Intn(sim.W)
                randY := rand.Intn(sim.H / 2)
                s = particles.NewSand(randX, randY)
                sim.AddParticle(s)
                sim.Draw()
                sim.Update()
                screen.SetContent(0, 0, ' ', status, tcell.StyleDefault)
                screen.Show()
            case <- quit:
                break loop
            }
        }
    }()

    for {
        switch ev := screen.PollEvent().(type) {
        case *tcell.EventKey:
            switch ev.Rune() {
            case 'q':
                quit <- struct{}{}
                screen.Fini()
                return
            case ' ':
                // randX := rand.Intn(sim.W)
                // randY := rand.Intn(sim.H / 2)
                s = particles.NewSand(sim.W / 2, 0)
                sim.AddParticle(s)
            case 's':
                clearBottomRow(sim)
            }
        }
    }
}

func clearBottomRow(sim *simulation.Simulation) {
    var newParticles []particles.Particle
    for _, p := range sim.Particles {
        if p.Pos().Y != sim.H - 1 {
            newParticles = append(newParticles, p)
        }
    }

    sim.Particles = newParticles
}

