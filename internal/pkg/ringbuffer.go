package pkg

import (
	"sync"
)

// RingBuffer represents a circular buffer
type RingBuffer[T any] struct {
	StoredStruct []T
	size         int
	index        int
	full         bool
	mutex        sync.Mutex
}

// NewRingBuffer creates a new circular buffer
func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		StoredStruct: make([]T, size),
		size:         size,
	}
}

// Add adds a message to the circular buffer
func (cb *RingBuffer[T]) Add(data T) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.StoredStruct[cb.index] = data
	cb.index = (cb.index + 1) % cb.size

	if cb.index == 0 {
		cb.full = true
	}
}

// Get retrieves the messages from the circular buffer
func (cb *RingBuffer[T]) Get() []T {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if !cb.full {
		return cb.StoredStruct[:cb.index]
	}

	return append(cb.StoredStruct[cb.index:], cb.StoredStruct[:cb.index]...)
}
