package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zapengine/zapengine"
	"github.com/zapengine/zapengine/components"
	"github.com/zapengine/zapengine/ecs"
	"github.com/zapengine/zapengine/input"
	"github.com/zapengine/zapengine/systems"
)

// MainMenu is our sample scene.
type MainMenu struct {
	world *ecs.World
	input *input.Input
}

func (s *MainMenu) Load() {
	s.world = ecs.NewWorld()
	s.input = input.NewInput()

	// Register systems
	s.world.AddSystem(systems.NewRenderSystem())
	s.world.AddSystem(&PlayerMovementSystem{input: s.input})

	// Create a simple procedural image for our sprite
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{255, 0, 0, 255}) // Red square

	// Create player entity
	player := s.world.CreateEntity()
	s.world.AddComponent(player, components.NewTransform(100, 100))
	s.world.AddComponent(player, components.NewSprite(img))

	// Create another entity to show rotating
	img2 := ebiten.NewImage(64, 16)
	img2.Fill(color.RGBA{0, 255, 0, 255}) // Green rectangle
	ent2 := s.world.CreateEntity()
	t2 := components.NewTransform(300, 200)
	s.world.AddComponent(ent2, t2)
	s.world.AddComponent(ent2, components.NewSprite(img2))
}

func (s *MainMenu) Unload() {
	s.world = nil
}

func (s *MainMenu) Update() error {
	return s.world.Update()
}

func (s *MainMenu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255}) // Dark blue background
	s.world.Draw(screen)
}

// PlayerMovementSystem handles player input and rotating the second entity.
type PlayerMovementSystem struct {
	input *input.Input
	timer float64
}

func (s *PlayerMovementSystem) Update(w *ecs.World) error {
	s.timer += 0.016 // Approximate 60fps dt

	// Manual system logic: loop through entities with transform.
	// In a real engine we'd use view/queries.
	// For demo, we just get the components.
	for _, e := range w.Entities() {
		// Just hack checking if the entity has a transform
		if tc, ok := ecs.GetComponent[*components.Transform](w, e); ok {
			t := *tc
			// If it's entity 1 (red square at 100,100 initially)
			if e == 1 {
				speed := 3.0
				if s.input.IsKeyPressed(ebiten.KeyArrowLeft) {
					t.X -= speed
				}
				if s.input.IsKeyPressed(ebiten.KeyArrowRight) {
					t.X += speed
				}
				if s.input.IsKeyPressed(ebiten.KeyArrowUp) {
					t.Y -= speed
				}
				if s.input.IsKeyPressed(ebiten.KeyArrowDown) {
					t.Y += speed
				}
			}

			// If it's entity 2
			if e == 2 {
				t.Rotation = s.timer * 2.0             // Rotate over time
				t.ScaleX = 1.0 + math.Sin(s.timer)*0.5 // Pulsate
				t.ScaleY = 1.0 + math.Sin(s.timer)*0.5
			}
		}
	}
	return nil
}

func (s *PlayerMovementSystem) Draw(w *ecs.World, screen *ebiten.Image) {}

func main() {
	engine := zapengine.NewEngine(640, 480, "ZapEngine Example")
	engine.SetScene(&MainMenu{})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
