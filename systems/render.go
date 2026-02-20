package systems

import (
	"reflect"

	"github.com/grapefizz/zapengine/components"
	"github.com/grapefizz/zapengine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// RenderSystem handles drawing entities with Transform and Sprite components.
type RenderSystem struct{}

// NewRenderSystem creates a RenderSystem.
func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

// Update is unused for RenderSystem but required by ecs.System.
func (s *RenderSystem) Update(w *ecs.World) error {
	return nil
}

// Draw renders entities to the screen.
func (s *RenderSystem) Draw(w *ecs.World, screen *ebiten.Image) {
	// Look up storage directly
	// Using values (not pointers) in reflecting type is standard for how we save components.
	// In our ECS we assume we stored *Transform and *Sprite
	transformType := reflect.TypeOf(&components.Transform{})
	spriteType := reflect.TypeOf(&components.Sprite{})

	transforms := w.GetStorage(transformType)
	sprites := w.GetStorage(spriteType)

	if transforms == nil || sprites == nil {
		return // Component types not registered
	}

	for _, e := range w.Entities() {
		// Optimization: check if both components exist
		if !transforms.Has(e) || !sprites.Has(e) {
			continue
		}

		tc, _ := transforms.Get(e)
		sc, _ := sprites.Get(e)

		t := tc.(*components.Transform)
		sp := sc.(*components.Sprite)

		if !sp.Visible || sp.Image == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}

		// 1. Move origin
		bounds := sp.Image.Bounds()
		w, h := float64(bounds.Dx()), float64(bounds.Dy())

		// Translate image to origin
		op.GeoM.Translate(-w*sp.OriginX, -h*sp.OriginY)

		// 2. Scale
		op.GeoM.Scale(t.ScaleX, t.ScaleY)

		// 3. Rotate
		op.GeoM.Rotate(t.Rotation)

		// 4. Translate back to position
		op.GeoM.Translate(t.X, t.Y)

		screen.DrawImage(sp.Image, op)
	}
}
