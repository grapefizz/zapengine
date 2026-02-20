package ecs

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

// World is the main container for the ECS.
type World struct {
	entities   []Entity
	components map[reflect.Type]*Storage
	systems    []System
}

// NewWorld creates a new ECS World.
func NewWorld() *World {
	return &World{
		entities:   make([]Entity, 0),
		components: make(map[reflect.Type]*Storage),
		systems:    make([]System, 0),
	}
}

// CreateEntity creates a new entity in the world.
func (w *World) CreateEntity() Entity {
	e := NewEntity()
	w.entities = append(w.entities, e)
	return e
}

// AddComponent adds a component to an entity in the world.
func (w *World) AddComponent(e Entity, c Component) {
	t := reflect.TypeOf(c)
	storage, ok := w.components[t]
	if !ok {
		storage = NewStorage()
		w.components[t] = storage
	}
	storage.Add(e, c)
}

// GetComponent retrieves a component of type T for the given entity.
// This uses generics for a cleaner API.
func GetComponent[T any](w *World, e Entity) (*T, bool) {
	var zero T
	t := reflect.TypeOf(zero)
	storage, ok := w.components[t]
	if !ok {
		return nil, false
	}

	// Check pointer type as well if value wasn't found
	// If we passed by value, reflect.TypeOf(zero) is struct{}
	// If we passed by pointer to AddComponent, reflect.TypeOf(c) was *struct{}
	// Let's standardise on passing values or pointers consistently. We assume values for now.

	c, ok := storage.Get(e)
	if !ok {
		return nil, false
	}

	val, ok := c.(T)
	if ok {
		return &val, true
	}

	valPtr, ok := c.(*T)
	if ok {
		return valPtr, true
	}

	return nil, false
}

// AddSystem registers a system with the world.
func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

// Update runs all systems' Update methods.
func (w *World) Update() error {
	for _, s := range w.systems {
		if err := s.Update(w); err != nil {
			return err
		}
	}
	return nil
}

// Draw runs all systems' Draw methods.
func (w *World) Draw(screen *ebiten.Image) {
	for _, s := range w.systems {
		s.Draw(w, screen)
	}
}

// View returns a list of entities that have all the specified component types.
// Since Go doesn't support variadic type parameters yet, we can't make this fully generic.
// Alternatively, systems can just query standard storage maps.
// For a simple ECS, we'll let systems manually query what they need.
func (w *World) Entities() []Entity {
	return w.entities
}

// GetStorage returns the storage for a specific component type.
func (w *World) GetStorage(t reflect.Type) *Storage {
	return w.components[t]
}
