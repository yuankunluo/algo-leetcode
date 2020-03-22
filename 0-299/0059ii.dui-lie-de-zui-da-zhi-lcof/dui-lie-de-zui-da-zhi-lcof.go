package problem0059ii

type MaxQueue struct {
	q   []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		q:   make([]int, 0),
		max: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	if len(this.max) != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q) != 0 {
		n = this.q[0]
		this.q = this.q[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
