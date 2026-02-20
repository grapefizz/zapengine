package main

import (
	"image/color"
	"log"

	"github.com/grapefizz/zapengine"
	"github.com/grapefizz/zapengine/components"
	"github.com/grapefizz/zapengine/ecs"
	"github.com/grapefizz/zapengine/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type BouncingScene struct {
	world *ecs.World
}

// BouncingSystem makes things bounce off the walls
type BouncingSystem struct {
	width, height float64
}

type Velocity struct {
	X, Y float64
}

func (s *BouncingSystem) Update(w *ecs.World) error {
	for _, e := range w.Entities() {
		tc, hasTransform := ecs.GetComponent[*components.Transform](w, e)
		vc, hasVel := ecs.GetComponent[*Velocity](w, e)

		if hasTransform && hasVel {
			t := *tc
			v := *vc
			t.X += v.X
			t.Y += v.Y

			// Simple bouncing logic
			if t.X <= 0 || t.X >= s.width-32 {
				v.X = -v.X
			}
			if t.Y <= 0 || t.Y >= s.height-32 {
				v.Y = -v.Y
			}
		}
	}
	return nil
}

func (s *BouncingSystem) Draw(w *ecs.World, screen *ebiten.Image) {}

func (s *BouncingScene) Load() {
	s.world = ecs.NewWorld()

	s.world.AddSystem(&BouncingSystem{width: 640, height: 480})
	s.world.AddSystem(systems.NewRenderSystem())

	img1 := ebiten.NewImage(32, 32)
	img1.Fill(color.RGBA{255, 100, 100, 255})

	e1 := s.world.CreateEntity()
	s.world.AddComponent(e1, components.NewTransform(100, 100))
	s.world.AddComponent(e1, components.NewSprite(img1))
	s.world.AddComponent(e1, &Velocity{X: 4.0, Y: 3.0})

	img2 := ebiten.NewImage(32, 32)
	img2.Fill(color.RGBA{100, 255, 100, 255})

	e2 := s.world.CreateEntity()
	s.world.AddComponent(e2, components.NewTransform(400, 300))
	s.world.AddComponent(e2, components.NewSprite(img2))
	s.world.AddComponent(e2, &Velocity{X: -5.0, Y: 2.0})
}

func (s *BouncingScene) Unload()       {}
func (s *BouncingScene) Update() error { return s.world.Update() }
func (s *BouncingScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{30, 30, 30, 255})
	s.world.Draw(screen)
}

func main() {
	engine := zapengine.NewEngine(640, 480, "ZapEngine: Bouncing Example")
	engine.SetScene(&BouncingScene{})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}
