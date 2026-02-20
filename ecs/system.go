package ecs

import "github.com/hajimehoshi/ebiten/v2"

// System defines the interface for logic that processes entities with specific components.
type System interface {
	// Update is called every game tick.
	Update(w *World) error

	// Draw is called every frame. Not all systems need to draw,
	// but rendering systems will utilize this.
	Draw(w *World, screen *ebiten.Image)
}
