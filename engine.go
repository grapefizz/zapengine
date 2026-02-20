package zapengine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Engine is the core structure for the ZapEngine game application.
type Engine struct {
	width, height int
	title         string
	currentScene  Scene
}

// NewEngine creates a new Engine instance.
func NewEngine(width, height int, title string) *Engine {
	return &Engine{
		width:  width,
		height: height,
		title:  title,
	}
}

// SetScene changes the currently active scene.
func (e *Engine) SetScene(scene Scene) {
	if e.currentScene != nil {
		e.currentScene.Unload()
	}
	e.currentScene = scene
	if e.currentScene != nil {
		e.currentScene.Load()
	}
}

// Update is called every tick by Ebitengine.
func (e *Engine) Update() error {
	if e.currentScene != nil {
		return e.currentScene.Update()
	}
	return nil
}

// Draw is called every frame by Ebitengine.
func (e *Engine) Draw(screen *ebiten.Image) {
	if e.currentScene != nil {
		e.currentScene.Draw(screen)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return e.width, e.height
}

// Run starts the game loop.
func (e *Engine) Run() error {
	ebiten.SetWindowSize(e.width, e.height)
	ebiten.SetWindowTitle(e.title)
	return ebiten.RunGame(e)
}
