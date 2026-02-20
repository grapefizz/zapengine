package main

import (
	"image/color"
	"log"

	"github.com/grapefizz/zapengine"
	"github.com/grapefizz/zapengine/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type InputScene struct {
	inp         *input.Input
	fontFace    text.Face
	lastPressed string
}

func (s *InputScene) Load() {
	s.inp = input.NewInput()

	// We'll just use a basic text rendering approach without a font file for simplicity in this example
	// Ebitengine requires a font face for text, so we'll just draw colored rectangles based on input instead
	// if a font isn't available. But let's build a visual representation.
}

func (s *InputScene) Unload() {}

func (s *InputScene) Update() error {
	if s.inp.IsKeyJustPressed(ebiten.KeySpace) {
		s.lastPressed = "Space just pressed!"
	} else if s.inp.IsKeyPressed(ebiten.KeySpace) {
		s.lastPressed = "Space is being held down."
	}

	if s.inp.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.lastPressed = "Left Mouse Clicked!"
	}

	return nil
}

func (s *InputScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 60, 255})

	// Draw mouse indicator
	mx, my := s.inp.CursorPosition()
	cursor := ebiten.NewImage(8, 8)
	cursor.Fill(color.RGBA{255, 255, 255, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(mx)-4, float64(my)-4)
	screen.DrawImage(cursor, op)
}

func main() {
	engine := zapengine.NewEngine(640, 480, "ZapEngine: Input Example")
	engine.SetScene(&InputScene{})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
