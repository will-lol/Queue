package walker

type Walker struct {
	data []any
	pos  uint8
}

type Equal func(i, j any) bool

func (walker *Walker) Init(len int) *Walker {
	walker.data = make([]any, len, len)
	walker.pos = 0
	return walker
}

func (walker *Walker) Walk(item any) {
	if int(walker.pos) > (len(walker.data) - 2) {
		walker.pos = 0
	} else {
		walker.pos++
	}
	walker.data[walker.pos] = item
}

func (walker *Walker) Check(equalFunc Equal) bool {
	prev := walker.data[0]
	for i := 1; i < len(walker.data); i++ {
		if equalFunc(walker.data[i], prev) {
			return true
		}
		prev = walker.data[i]
	}
	return false
}

func NewWalker(len int) *Walker {
	return new(Walker).Init(len)
}