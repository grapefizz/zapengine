package ecs

// Component is a generic interface for all component data.
type Component any

// Storage manages components of a specific type.
// For simplicity in this demo, we use a map from Entity to Component.
// A more performant approach would use dense/sparse arrays.
type Storage struct {
	components map[Entity]Component
}

// NewStorage creates a new component storage.
func NewStorage() *Storage {
	return &Storage{
		components: make(map[Entity]Component),
	}
}

// Add attaches a component to an entity.
func (s *Storage) Add(e Entity, c Component) {
	s.components[e] = c
}

// Get retrieves a component for an entity.
func (s *Storage) Get(e Entity) (Component, bool) {
	c, ok := s.components[e]
	return c, ok
}

// Remove removes a component from an entity.
func (s *Storage) Remove(e Entity) {
	delete(s.components, e)
}

// Has checks if an entity has this component.
func (s *Storage) Has(e Entity) bool {
	_, ok := s.components[e]
	return ok
}
