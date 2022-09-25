package utils

import (
	"ebiten_mouse_move_effect/vec2d"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func Clamp(number, min, max float64) float64 {
	if number < min {
		return min
	}
	if number > max {
		return max
	}
	return number
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func MouseHasMoved(mouse, previousMouse *vec2d.Vec2D) bool {
	x, y := ebiten.CursorPosition()
	mouse.X = float64(x)
	mouse.Y = float64(y)

	returnValue := !mouse.IsEqualTo(*previousMouse)

	*previousMouse = *mouse

	return returnValue
}
