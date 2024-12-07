package geom

import (
    "math/rand"
)

type Vec2 struct {
    X, Y int
}

func Shuffle(positions []Vec2) {
    n := len(positions)
    for i := n - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        positions[i], positions[j] = positions[j], positions[i]
    }
}
