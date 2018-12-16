package heap

type Heap struct {
	a     []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	return &Heap{a: make([]int, capacity+1), n: capacity, count: 0}
}

func (h *Heap) Insert(data int) {
	if h.count >= h.n {
		return
	}
	h.count++
	h.a[h.count] = data
	i := h.count
	for i/2 > 0 && h.a[i] > h.a[i/2] {
		swap(h.a, i, i/2)
		i = i / 2
	}

}

func swap(a []int, i, j int) {
	temp := a[i-1]
	a[i-1] = a[j-1]
	a[j-1] = temp
}

func Sort(a []int) {
	N := len(a)
	for k := N / 2; k >= 1; k-- {
		sink(a, k, N)
	}
	for N > 1 {
		swap(a, 1, N)
		N--
		sink(a, 1, N)
	}
}

func sink(a []int, i, n int) {
	for 2*i <= n {
		j := 2 * i
		if j < n && less(a, j, j+1) {
			j++
		}
		if !less(a, i, j) {
			break
		}
		swap(a, i, j)
		i = j
	}
}

func less(a []int, i, j int) bool {
	return a[i-1] < a[j-1]
}
