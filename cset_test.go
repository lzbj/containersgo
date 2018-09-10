package containersgo

import (
	"sync"
	"testing"
)

func TestCSetAdd(t *testing.T) {
	set := NewCSet()

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

func BenchmarkHashCSetAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	set := NewCSet()
	b.StartTimer()
	cbenchmark(b, set, size, cadd)
}

// Test add items concurrently.
func TestCSetAddSafe(t *testing.T) {
	set := NewCSet()

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

func cbenchmark(b *testing.B, set cset, size int, op caction) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			switch op {
			case cadd:
				set.Add(n)
			}
		}
	}
}
