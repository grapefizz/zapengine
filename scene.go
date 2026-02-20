package zapengine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents a distinct state or screen in the game (e.g., Main Menu, Level 1).
type Scene interface {
	// Load is called when the scene becomes active.
	Load()
	// Unload is called when the scene is no longer active.
	Unload()
	// Update is called every game tick. Return error to exit.
	Update() error
	// Draw is called every frame to render the scene.
	Draw(screen *ebiten.Image)
}
