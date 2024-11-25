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


    // p = particles.NewParticle()
    s := particles.NewSand(1, 1)

    sim.AddParticle(s)

    quit := make(chan struct{})

    start := time.Now()

    go func() {
        ticker := time.NewTicker(time.Second / 10)
        tickerReset := time.NewTicker(time.Minute)
        defer ticker.Stop()

        loop:
        for {
            select {
            case <-tickerReset.C:
                var newParticles []particles.Particle
                for _, p := range sim.Particles {
                    if p.Pos().Y != sim.H - 1 {
                        newParticles = append(newParticles, p)
                    }
                }

                sim.Particles = newParticles
            case <-ticker.C:
                status := []rune(fmt.Sprintf("sim.H = %d | sim.W = %d | len(sim.Particles) = %d | timePassed = %f", sim.H, sim.W, len(sim.Particles), time.Now().Sub(start).Seconds()))
                randX := rand.Intn(sim.W)
                randY := rand.Intn(sim.H / 2)
                s = particles.NewSand(randX, randY)
                sim.AddParticle(s)
                sim.Draw()
                sim.Update()
                screen.SetContent(0, 0, '_', status, tcell.StyleDefault)
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
                randX := rand.Intn(sim.W)
                randY := rand.Intn(sim.H / 2)
                s = particles.NewSand(randX, randY)
                sim.AddParticle(s)
            }
        }
    }
}

