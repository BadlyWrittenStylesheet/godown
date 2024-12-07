package main

import (
	"fmt"
	"math/rand"
	"time"

	// "godown/geom"
	"godown/cursor"
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
    // w, h := screen.Size()
    // 626 132

    s := particles.NewSand(1, 1)
    sim.AddParticle(s)
    fmt.Println(s)

    quit := make(chan struct{})
    start := time.Now()

    cur := cursor.NewCursor(sim.W, sim.H, particles.NewSand(1, 1))

    go func() {
        ticker := time.NewTicker(time.Second / 30)
        particleTicker := time.NewTicker(time.Second / 5)
        defer ticker.Stop()

        var tsB time.Time
        var tsA time.Time

        loop:
        for {
            select {
            case <- particleTicker.C:
                randX := rand.Intn(sim.W)
                // randY := rand.Intn(sim.H / 2)
                s = particles.NewSand(randX, 1)
                // sim.AddParticle(s)

            case <-ticker.C:
                // if len(sim.Particles) >= sim.H * sim.W {
                //     sim.Particles = []*particles.Particle{}
                //     sim.Grid = [][]particles.Particle{}
                // }
                stats := []rune(generateStats(start, tsA.Sub(tsB), sim))
                // iterTime = time.Now()
                sim.Draw()
                tsB = time.Now()
                sim.Update()
                tsA = time.Now()
                screen.SetContent(cur.X, cur.Y, '+', nil, tcell.StyleDefault)
                screen.SetContent(0, 0, ' ', stats, tcell.StyleDefault)
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
                // s = particles.NewSand(randX, randY)
                s := cur.SpawnParticle()
                sim.AddParticle(s)
            case 'n':
                switch cur.Particle.(type) {
                case *particles.Water:
                    cur.Particle = particles.NewSand(0, 0)
                case *particles.Sand:
                    cur.Particle = particles.NewWater(0, 0)
                }
            case 'x':
                sim.Grid[cur.X][cur.Y] = nil
            case 'h':
                cur.Move(cursor.West)
            case 'j':
                cur.Move(cursor.South)
            case 'k':
                cur.Move(cursor.North)
            case 'l':
                cur.Move(cursor.East)
            }
        }
    }
}

func generateStats(simStart time.Time, updateTime time.Duration, sim *simulation.Simulation) []rune {
    stats := []rune(fmt.Sprintf("sim.H = %d | sim.W = %d | particles = %d | timePassed = %f | updateTime = %f", sim.H, sim.W, sim.ParticleCount, time.Now().Sub(simStart).Seconds(), updateTime.Seconds()))
    return stats
}

