package bbref

type BaseState struct {
	arr [3]*Player
}

func NewBaseState() *BaseState {
	return &BaseState{
		arr: [3]*Player{
			nil, nil, nil,
		},
	}
}

func (b *BaseState) Empty() bool {
	if on := b.GuysOn(); on == 0 {
		return true
	}
	return false
}

func (b *BaseState) GuysOn() int {
	var on = 0
	for _, base := range b.arr {
		if base != nil {
			on += 1
		}
	}
	return on
}

func (b *BaseState) GuyOn(i int) bool {
	if on := b.arr[i]; on != nil {
		return true
	}
	return false
}

func (b *BaseState) Clear() {
	b.arr = [3]*Player{
		nil, nil, nil,
	}
}

func (b *BaseState) Move(numBases int) int {
	if b.Empty() {
		return 0
	}

	var runs int

	for i := 3; i != 1; i-- {
		if !b.GuyOn(i - 1) {
			continue
		}

		if i+numBases < 3 {
			b.arr[i-1+numBases] = b.arr[i-1]
			b.arr[i-1] = nil
		} else {
			b.arr[i] = nil
			runs += 1
		}
	}
	return runs
}

var (
	runs = 0
	nums = map[string]int{
		"1B":  1,
		"2B":  2,
		"3B":  3,
		"HR":  3,
		"BB":  1,
		"HBP": 1,
	}
)

// Handle -> responsible for calculating runs scored
// on a given result from a plate appearance and also modifying
// the base state in place (moving the runners)
func (b *BaseState) Handle(hitter *Player, result string) (int, error) {

	if _, isIn := nums[result]; !isIn {
		return 0, nil
	}

	numBases := nums[result]
	if result == "HR" {
		b.Clear()
		return b.Move(numBases) + 1, nil
	}

	runs += b.Move(numBases)
	b.arr[numBases] = hitter
	return runs, nil
}
