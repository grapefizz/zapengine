package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input provides a simplified interface for checking Ebitengine input state.
type Input struct {
	// Usually we'd track more complex state, but Ebitengine and inpututil
	// handle most of this internally. We expose helper methods for the engine.
}

// NewInput creates a new Input instance.
func NewInput() *Input {
	return &Input{}
}

// IsKeyPressed returns true if the key is currently held down.
func (i *Input) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

// IsKeyJustPressed returns true only on the frame the key was pressed.
func (i *Input) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

// IsMouseButtonPressed returns true if the mouse button is currently held down.
func (i *Input) IsMouseButtonPressed(button ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(button)
}

// IsMouseButtonJustPressed returns true only on the frame the mouse button was pressed.
func (i *Input) IsMouseButtonJustPressed(button ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(button)
}

// CursorPosition returns the current mouse cursor position.
func (i *Input) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}
