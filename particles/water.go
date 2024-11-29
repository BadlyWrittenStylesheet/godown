package particles

import (
	"godown/geom"

	"github.com/gdamore/tcell/v2"
)

type Water struct {
    *BaseParticle
}

func NewWater(x, y int) *Water {
    return &Water{
        BaseParticle: &BaseParticle{
            pos: geom.Vec2{X: x, Y: y},
            vel: geom.Vec2{X: 0, Y: 1},
            char: '█',
            style: tcell.StyleDefault.Foreground(tcell.ColorCornflowerBlue),
        },
    }
}


func (s *Water) Style() tcell.Style {
    return s.style
}

func (s *Water) Update(w, h int, grid [][]Particle) {
    possiblePositions := []geom.Vec2{
        {X: 0, Y: 1},
        {X: 1, Y: 1},
        {X: -1, Y: 1},
        {X: 1, Y: 0},
        {X: -1, Y: 0},
    }

    curX := s.Pos().X
    curY := s.Pos().Y

    for _, off := range possiblePositions {
        newX := curX + off.X
        newY := curY + off.Y

        if newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[newX]) {
            continue
        }

        if grid[newX][newY] == nil {
            s.pos = geom.Vec2{X: newX, Y: newY}
            grid[curX][curY] = nil
            grid[newX][newY] = s

            return
        }
    }
}

func (s *Water) Draw(screen tcell.Screen) {
    pos := s.Pos()
    screen.SetContent(pos.X, pos.Y, s.Char(), nil, s.Style())
}

func (s *Water) Pos() (geom.Vec2) {
    return s.pos
}

func (s *Water) Char() rune {
    return s.char
}

