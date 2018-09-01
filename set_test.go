package containersgo

import (
	"sync"
	"testing"
)

func TestSetAdd(t *testing.T) {

	set := New()
	set.Add()
	set.Add(1)
	set.Add(2, 1)
	set.Add(2, 3)
	set.Add()
	if actual := set.Empty(); actual != false {
		t.Errorf("Got %v expected %v", actual, false)
	}
	if actualSize := set.Size(); actualSize != 3 {
		t.Errorf("Got %v expected %v", actualSize, 3)
	}
}

func TestSetRemove(t *testing.T) {
	set := New()
	set.Add(3, 1, 2, 4)
	set.Remove()
	if actual := set.Size(); actual != 4 {
		t.Errorf("Got %v expected %v", actual, 4)
	}
	set.Remove(1)
	if actual := set.Size(); actual != 3 {
		t.Errorf("Got %v expected %v", actual, 3)
	}
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actual := set.Size(); actual != 1 {
		t.Errorf("Got %v expected %v", actual, 1)
	}
}

// Test add items concurrently.
func TestSetAddSafe(t *testing.T) {
	set := New()

	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			set.Add(i)
			wg.Done()
		}(i)

	}

	wg.Wait()

	if actual := set.Size(); actual != 1000 {
		t.Errorf("Got %v expected %v", actual, 1000)
	}
}

// Test add items concurrently.
func TestSetRemoveSafe(t *testing.T) {
	set := New()

	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func(i int) {
				set.Add(i)
				wg.Done()
			}(i)

		}
	}()

	wg.Wait()

	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go func(i int) {
				set.Remove(i)
				wg.Done()
			}(i)

		}
	}()
	wg.Wait()

	if actual := set.Size(); actual != 0 {
		t.Errorf("Got %v expected %v", actual, 0)
	}
}

const (
	add = iota
	remove
)

func BenchmarkHashSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	b.StartTimer()
	benchmark(b, set, size, add)
}

func BenchmarkHashSetRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := New()
	b.StartTimer()
	benchmark(b, set, size, remove)
}

func benchmark(b *testing.B, set *Set, size int, op int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			switch op {
			case add:
				set.Add(n)
			case remove:
				set.Remove(n)
			}

		}
	}
}
