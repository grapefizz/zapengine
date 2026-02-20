package components

import "github.com/hajimehoshi/ebiten/v2"

// Sprite represents an image to be rendered for an entity.
type Sprite struct {
	Image   *ebiten.Image
	Visible bool
	// Optional rendering properties could be added like anchors/origins
	OriginX float64 // 0 is left, 0.5 is center, 1 is right
	OriginY float64 // 0 is top, 0.5 is center, 1 is bottom
}

// NewSprite creates a new Sprite component for an image.
func NewSprite(img *ebiten.Image) *Sprite {
	return &Sprite{
		Image:   img,
		Visible: true,
		OriginX: 0.5,
		OriginY: 0.5,
	}
}
