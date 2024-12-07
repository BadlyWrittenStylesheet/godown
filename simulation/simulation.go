package simulation

import (
	"godown/particles"

	"github.com/gdamore/tcell/v2"
)

type Simulation struct {
    screen tcell.Screen
    Grid [][]particles.Particle
    H int
    W int
    IterCount int
    ParticleCount int
}

func (s *Simulation) AddParticle(p particles.Particle) {
    s.Grid[p.Pos().X][p.Pos().Y] = p
}

func (s *Simulation) Update() {
    newCount := 0
    for i := len(s.Grid) - 1; i >= 0 ; i-- {
        for j := len(s.Grid[i]) - 1; j >= 0; j--{
            p := s.Grid[i][j]
            if p != nil {
                newCount += 1
                p.Update(s.W, s.H, s.Grid)
            }
        }
    }
    s.ParticleCount = newCount
}

func (s *Simulation) Draw() {
    s.screen.Clear()
    for _, r := range s.Grid {
        for _, p := range r {
            if p != nil {
                p.Draw(s.screen)
            }
        }
    }
}

func NewSimulation(screen tcell.Screen) *Simulation {
    w, h := screen.Size()
    grid := make([][]particles.Particle, w)
    for i := range grid {
        grid[i] = make([]particles.Particle, h)
    }

    return &Simulation{
        screen: screen,
        Grid: grid,
        W: w,
        H: h,
        IterCount: 0,
    }
}

