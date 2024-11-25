package simulation

import (
	"github.com/gdamore/tcell/v2"
	// "godown/geom"
	"godown/geom"
	"godown/particles"
)

type Simulation struct {
    screen tcell.Screen
    Particles []particles.Particle
    H int
    W int
}

func (s *Simulation) AddParticle(p particles.Particle) {
    s.Particles = append(s.Particles, p)
}

func (s *Simulation) Update() {
    for _, p := range s.Particles {
        p.Update(s.W, s.H, s.Particles)
    }
}

func (s *Simulation) ParticleAt(pos geom.Vec2) bool {
    for _, p := range s.Particles {
        if p.Pos() == pos {
            return true
        }
    }
    return false
}

func (s *Simulation) Draw() {
    s.screen.Clear()
    for _, p := range s.Particles {
        p.Draw(s.screen)
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
    return &Simulation{
        screen: screen,
        W: w,
        H: h,
    }
}
