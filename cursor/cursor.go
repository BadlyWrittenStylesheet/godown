package cursor

import "godown/particles"

type Direction int

const (
    North = iota
    South
    West
    East
)


type Cursor struct {
    X int
    Y int
    MaxX int
    MaxY int
    direction Direction
    particle particles.Particle 
}

func NewCursor(maxX, maxY int, particle particles.Particle) *Cursor {
    return &Cursor{
        X: 1,
        Y: 1,
        MaxX: maxX,
        MaxY: maxY,
        direction: South,
        particle: particle,
    }
}

func (c *Cursor) Move(to Direction) {
    newX := c.X
    newY := c.Y
    switch to {
    case North:
        newY = c.Y - 1
    case South:
        newY = c.Y + 1
    case East:
        newX = c.X + 1
    case West:
        newX = c.X - 1
    }

    if !(newX > c.MaxX || newX < 0) {
        c.X = newX
    }
    
    if !(newY > c.MaxY || newY < 0) {
        c.Y = newY
    }
}

func (c *Cursor) SpawnParticle() particles.Particle {
    switch c.particle.(type) {
    case *particles.Sand:
        return particles.NewSand(c.X, c.Y)
    }
    return nil
}



