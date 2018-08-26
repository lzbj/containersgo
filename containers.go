package containersgo

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}
