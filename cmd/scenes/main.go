package main

import (
	"image/color"
	"log"

	"github.com/grapefizz/zapengine"
	"github.com/grapefizz/zapengine/input"
	"github.com/hajimehoshi/ebiten/v2"
)

var globalEngine *zapengine.Engine

// SceneA is the first scene (Red)
type SceneA struct {
	inp *input.Input
}

func (s *SceneA) Load() {
	s.inp = input.NewInput()
	log.Println("Scene A Loaded")
}

func (s *SceneA) Unload() {
	log.Println("Scene A Unloaded")
}

func (s *SceneA) Update() error {
	if s.inp.IsKeyJustPressed(ebiten.KeySpace) {
		globalEngine.SetScene(&SceneB{})
	}
	return nil
}

func (s *SceneA) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{200, 50, 50, 255}) // Red
}

// SceneB is the second scene (Blue)
type SceneB struct {
	inp *input.Input
}

func (s *SceneB) Load() {
	s.inp = input.NewInput()
	log.Println("Scene B Loaded")
}

func (s *SceneB) Unload() {
	log.Println("Scene B Unloaded")
}

func (s *SceneB) Update() error {
	if s.inp.IsKeyJustPressed(ebiten.KeySpace) {
		globalEngine.SetScene(&SceneA{})
	}
	return nil
}

func (s *SceneB) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 200, 255}) // Blue
}

func main() {
	globalEngine = zapengine.NewEngine(640, 480, "ZapEngine: Scene Switching (Press Space)")
	globalEngine.SetScene(&SceneA{})

	if err := globalEngine.Run(); err != nil {
		log.Fatal(err)
	}
}
