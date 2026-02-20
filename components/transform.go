package components

// Transform defines the position, rotation, and scale of an entity.
type Transform struct {
	X        float64
	Y        float64
	Rotation float64 // in radians
	ScaleX   float64
	ScaleY   float64
}

// NewTransform creates a new Transform component with default values.
func NewTransform(x, y float64) *Transform {
	return &Transform{
		X:      x,
		Y:      y,
		ScaleX: 1.0,
		ScaleY: 1.0,
	}
}
