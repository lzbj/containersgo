package containersgo

type Set struct {
	items map[interface{}]struct{}

	lock sync.RWMutex
}

func New(items ...interface{}) *Set {
	set := &Set{
		items: make(map[interface{}]struct{}, 48)}
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set

}

func (set *Set) Size() int64 {
	set.lock.RLock()

	size := int64(len(set.items))

	set.lock.RUnlock()

	return size
}

func (set *Set) Clear() {

	set.lock.Lock()

	set.items = map[interface{}]struct{}{}

	st.lock.Unlock()
}

func (set *Set) Empty() bool {

	return set.Size() == 0

}

func (set *Set) Values() []interface{} {

	set.lock.Lock()

	defer set.lock.Unlock()
	values := make([]interface{}, set.Size())

	count := 0

	for item := range set.items {

		values[count] = item
		count++
	}
	return values

}

func (set *Set) Add(items ...interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()

	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

func (set *Set) Remove(items ...interface{}) {

	set.lock.Lock()
	defer set.lock.Unlock()

	for _, item := range items {

		delete(set.items, item)
	}

}

func (set *Set) Exists(item interface{}) bool {

	set.lock.RLock()
	_, ok := set.items[item]
	set.lock.RUnlock()
	return ok

}
