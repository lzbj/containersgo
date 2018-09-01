package containersgo

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

func TestListBasic(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}

func TestListBasicUnSafe(t *testing.T) {
	// Create a new list and put some numbers in it concurrently
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		l.InsertBefore(3, e4)
		wg.Done()
	}()
	go func() {
		l.InsertBefore(5, e4)
		wg.Done()
	}()
	go func() {
		l.InsertAfter(2, e1)
		wg.Done()
	}()
	wg.Wait()

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
