package particles

import (
	"godown/geom"

	"github.com/gdamore/tcell/v2"
)

type Particle interface {
    Update(w, h int, grid [][]Particle)
    Draw(screen tcell.Screen)
    Pos() geom.Vec2
    Char() rune
    Style() tcell.Style
}

type BaseParticle struct {
    pos geom.Vec2
    vel geom.Vec2
    char rune
    style tcell.Style
}

func (p *BaseParticle) Style() tcell.Style {
    return tcell.StyleDefault
}

func (p *BaseParticle) Update(w, h int, grid [][]Particle) {}

func (p *BaseParticle) Draw(screen tcell.Screen) {
    pos := p.Pos()
    screen.SetContent(pos.X, pos.Y, p.char, nil, tcell.StyleDefault)
}

func (p *BaseParticle) Pos() geom.Vec2 {
    return p.pos
}

func NewParticle(x, y, vx, vy int, char rune) *BaseParticle {
    return &BaseParticle{
        pos: geom.Vec2{X: x, Y: y},
        vel: geom.Vec2{X: vx, Y: vy},
        char: char,
        style: tcell.StyleDefault,
    }
}


