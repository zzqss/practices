package heap

// slice must start by 1
type Heap struct {
	slice     []int
	length    int
	MaxLength int
}

func NewHeap(length int) *Heap {
	length += 1
	return &Heap{make([]int, length, length), 1, length}
}

func (this *Heap) ShiftUp(value int) {
	if this.length >= this.MaxLength {
		return
	}
	this.length += 1
	this.slice[this.length] = value
	last := position
	now := last / 2
	for now > 0 {
		if this.slice[now] > value {
			this.slice[last] = this.slice[now]
		} else {
			break
		}
		last = now
		now = last / 2
	}
	this.slice[last] = value
}

func (this *Heap) ShiftDown(value int) {
	last := 1
	now := 1
	for last <= this.length {

	}
}
