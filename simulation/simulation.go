package simulation

import (
	"github.com/gdamore/tcell/v2"
	"godown/particles"
)

type Simulation struct {
    screen tcell.Screen
    Particles []*particles.Particle
    Grid [][]particles.Particle
    H int
    W int
}

func (s *Simulation) AddParticle(p particles.Particle) {
    s.Particles = append(s.Particles, &p)
    s.Grid[p.Pos().X][p.Pos().Y] = p
}

func (s *Simulation) Update() {
    for _, p := range s.Particles {
        (*p).Update(s.W, s.H, s.Grid)
    }
}

func (s *Simulation) Draw() {
    s.screen.Clear()
    for _, p := range s.Particles {
        (*p).Draw(s.screen)
    }
}

// func (p *Particle) Update(w, h int, particles []*Particle) {
//     newX, newY := p.pos.x + p.vel.x, p.pos.y + p.vel.y
//     if newX > w {
//         newX = w - 1
//     }
//     if newX < 0 {
//         newX = 0
//     }

//     if newY > h {
//         newY = h - 1
//     }

//     if newY < 0 {
//         newY = 0
//     }
//     canMove := true
//     for _, other := range particles {
//         if other != p && other.pos.x == newX && other.pos.y == newY {
//             canMove = false
//             break
//         }
//     }
//     if canMove {
//         p.pos.x = newX
//         p.pos.y = newY
//     }
// }

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
    }
}

