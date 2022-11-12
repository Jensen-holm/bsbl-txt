package bbref

type BaseState struct {
	first  bool
	second bool
	third  bool
	m      map[string]bool
}

func NewBaseState() *BaseState {
	return &BaseState{
		false,
		false,
		false,
		map[string]bool{
			"1B": false,
			"2B": false,
			"3B": false,
		},
	}
}

func (b *BaseState) ClearBases() {
	b.first, b.second, b.third = false, false, false
	b.m = map[string]bool{
		"1B": b.first,
		"2B": b.second,
		"3B": b.third,
	}
}

func (b *BaseState) GuysOn() int {
	var on = 0
	for _, st := range b.m {
		if st {
			on += 1
		}
	}
	return on
}