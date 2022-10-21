type Vector2D struct {
	queue []int
}

func Constructor(vec [][]int) Vector2D {
	var queue []int
	for _, ones := range vec {
		queue = append(queue, ones...)
	}
	return Vector2D{queue: queue}
}

func (v2 *Vector2D) Next() int {
	t := v2.queue[0]
	v2.queue = v2.queue[1:]
	return t
}

func (v2 *Vector2D) HasNext() bool {
	return len(v2.queue) != 0
}
