package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/grapefizz/zapengine"
	"github.com/grapefizz/zapengine/components"
	"github.com/grapefizz/zapengine/ecs"
	"github.com/grapefizz/zapengine/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type StressScene struct {
	world *ecs.World
}

// SpinSystem makes things rotate
type SpinSystem struct{}

func (s *SpinSystem) Update(w *ecs.World) error {
	for _, e := range w.Entities() {
		if tc, ok := ecs.GetComponent[*components.Transform](w, e); ok {
			t := *tc
			t.Rotation += 0.05
		}
	}
	return nil
}
func (s *SpinSystem) Draw(w *ecs.World, screen *ebiten.Image) {}

func (s *StressScene) Load() {
	s.world = ecs.NewWorld()

	s.world.AddSystem(&SpinSystem{})
	s.world.AddSystem(systems.NewRenderSystem())

	img := ebiten.NewImage(8, 8)
	img.Fill(color.RGBA{200, 200, 255, 255})
	sprite := components.NewSprite(img)

	// Create 10,000 entities
	for i := 0; i < 10000; i++ {
		e := s.world.CreateEntity()
		x := rand.Float64() * 640
		y := rand.Float64() * 480
		s.world.AddComponent(e, components.NewTransform(x, y))
		s.world.AddComponent(e, sprite) // Reuse the same sprite component
	}
}

func (s *StressScene) Unload()       {}
func (s *StressScene) Update() error { return s.world.Update() }
func (s *StressScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 20, 255})
	s.world.Draw(screen)
}

func main() {
	engine := zapengine.NewEngine(640, 480, "ZapEngine: Stress Test (10k Entities)")
	engine.SetScene(&StressScene{})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
