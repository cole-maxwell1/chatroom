package pkg

import (
	"sync"
	"testing"

)

// TestRingBuffer_Add tests the Add method of RingBuffer
func TestRingBuffer_Add(t *testing.T) {
	bufferSize := 3
	rb := NewRingBuffer[int](bufferSize)

	rb.Add(1)
	rb.Add(2)
	rb.Add(3)

	if got, want := rb.StoredStruct, []int{1, 2, 3}; !equal(got, want) {
		t.Errorf("StoredStruct = %v, want %v", got, want)
	}

	rb.Add(4)

	if got, want := rb.StoredStruct, []int{4, 2, 3}; !equal(got, want) {
		t.Errorf("StoredStruct = %v, want %v", got, want)
	}
}

// TestRingBuffer_Get tests the Get method of RingBuffer
func TestRingBuffer_Get(t *testing.T) {
	bufferSize := 3
	rb := NewRingBuffer[int](bufferSize)

	rb.Add(1)
	rb.Add(2)

	if got, want := rb.Get(), []int{1, 2}; !equal(got, want) {
		t.Errorf("Get() = %v, want %v", got, want)
	}

	rb.Add(3)
	rb.Add(4)

	if got, want := rb.Get(), []int{2, 3, 4}; !equal(got, want) {
		t.Errorf("Get() = %v, want %v", got, want)
	}
}

// TestRingBuffer_Concurrency tests concurrent access to the RingBuffer
func TestRingBuffer_Concurrency(t *testing.T) {
	bufferSize := 5
	rb := NewRingBuffer[int](bufferSize)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			rb.Add(i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			rb.Get()
		}
	}()

	wg.Wait()
}

// equal is a helper function to compare two slices for equality
func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
