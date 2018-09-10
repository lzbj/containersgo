package containersgo

type cset chan commandData

type caction int
type commandData struct {
	action  caction
	keys    []interface{}
	result  chan<- interface{}
	data    chan<- map[interface{}]struct{}
	updater UpdateFunc
}

const (
	cadd caction = iota
	cremove
	cexist
	csize
)

type findResult struct {
	value struct{}
	found bool
}

type UpdateFunc func(interface{}, bool) interface{}

func NewCSet(items ...interface{}) cset {
	set := make(cset, 1000)
	go set.run()
	set.Add(items...)

	return set
}

func (cs cset) run() {
	items := make(map[interface{}]struct{})

	for command := range cs {
		switch command.action {
		case cadd:
			for _, item := range command.keys {
				items[item] = struct{}{}
			}

		case csize:

			command.result <- int64(len(items))
		}

	}

}

func (cs cset) Add(items ...interface{}) {
	cs <- commandData{action: cadd, keys: items}

}

func (cs cset) Size() int64 {
	reply := make(chan interface{})
	cs <- commandData{action: csize, result: reply}
	return (<-reply).(int64)
}

func (cs *cset) Empty() bool {
	return cs.Size() == 0

}
