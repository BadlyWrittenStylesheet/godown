package particles

import (
	"godown/geom"

	"github.com/gdamore/tcell/v2"
)

type Sand struct {
    *BaseParticle
}

func NewSand(x, y int) *Sand {
    return &Sand{
        BaseParticle: &BaseParticle{
            pos: geom.Vec2{X: x, Y: y},
            vel: geom.Vec2{X: 0, Y: 1},
            char: '█',
            style: tcell.StyleDefault.Foreground(tcell.ColorLightGoldenrodYellow),
        },
    }
}


func (s *Sand) Style() tcell.Style {
    return s.style
}

func (s *Sand) Update(w, h int, grid [][]Particle) {
    possiblePositions := []geom.Vec2{
        {X: 0, Y: 1},
        {X: 1, Y: 1},
        {X: -1, Y: 1},
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
        // if newY + 1 < len(grid[newX]) {
            // if _, ok := grid[newX][newY].(*Water); ok {
                // panic(newY)
                // grid[newX][newY], grid[curX][curY] =  grid[curX][curY], grid[newX][newY]
            // }
        // }
    }
}

func (s *Sand) Draw(screen tcell.Screen) {
    pos := s.Pos()
    screen.SetContent(pos.X, pos.Y, s.Char(), nil, s.Style())
}

func (s *Sand) Pos() (geom.Vec2) {
    return s.pos
}

func (s *Sand) Char() rune {
    return s.char
}

