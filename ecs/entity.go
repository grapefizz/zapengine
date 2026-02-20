package ecs

import "sync/atomic"

// Entity is a unique identifier.
type Entity uint32

// entityCounter generates unique entity IDs.
var entityCounter uint32

// NewEntity creates a new unique Entity.
func NewEntity() Entity {
	return Entity(atomic.AddUint32(&entityCounter, 1))
}
