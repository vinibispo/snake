package helpers

import rl "github.com/gen2brain/raylib-go/raylib"

func ElementInSlice(element rl.Vector2, list []rl.Vector2) bool {
	for _, b := range list {
		if rl.Vector2Equals(element, b) {
			return true
		}
	}
	return false
}
