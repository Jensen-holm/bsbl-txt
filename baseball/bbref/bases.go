package bbref

type BaseState struct {
	m map[string]bool
	s []bool
}

func NewBaseState() *BaseState {
	return &BaseState{
		map[string]bool{
			"1B": false,
			"2B": false,
			"3B": false,
		},
		[]bool{
			false,
			false,
			false,
		},
	}
}

func (b *BaseState) ClearBases() {
	b.m = map[string]bool{
		"1B": false,
		"2B": false,
		"3B": false,
	}
	b.s = []bool{
		false,
		false,
		false,
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

func (b *BaseState) AdvanceOnHit(r string) int {
	var runs = 0
	for i, base := range b.s {
		if base && r == "1B" {
			newIndex := i + 1
			if newIndex > len(b.s) {
				runs += 1
				b.s[i] = false
				b.s[0] = true
			} else {
				b.s[newIndex] = true
				b.s[0] = true
			}
		} else if base && r == "2B" {
			newIndex := i + 2
			if newIndex > len(b.s) {
				runs += 1
				b.s[i] = false
				b.s[1] = true
			} else {
				b.s[newIndex] = true
				b.s[1] = true
			}
		} else if base && r == "3B" {
			runs += b.GuysOn()
			b.ClearBases()
			b.s[2] = true
		}
	}
	return runs
}

func (b *BaseState) HandleBases(r string) (int, error) {

	if r == "2B" || r == "3B" || r == "1B" {
		runs := b.AdvanceOnHit(r)
		return runs, nil
	}

	if r == "HR" {
		runs := 1 + b.GuysOn()
		b.ClearBases()
		return runs, nil
	}

	if r == "BB" || r == "HBP" {
		if b.GuysOn() == 3 {
			return 1, nil
		}

		// this is not correct and I know ...
		if b.GuysOn() > 0 {
			runs := b.AdvanceOnHit("1B")
			return runs, nil
		}

	}
	return 0, nil
}
