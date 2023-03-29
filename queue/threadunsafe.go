package queue

type Queue[T comparable] struct {
	queue []T
}

func BuildQueue[T comparable](data ...T) *Queue[T] {
	que := &Queue[T]{queue: make([]T, 0)}
	que.Enqueue(data...)
	return que
}

func (q *Queue[T]) Enqueue(data ...T) {
	q.queue = append(q.queue, data...)
}

func (q *Queue[T]) Dequeue() T {
	var data T
	if q.IsEmpty() {
		return data
	}

	data = q.queue[0]
	q.queue = q.queue[1:]

	return data
}

func (q *Queue[T]) Peek() T {
	var data T
	if q.IsEmpty() {
		return data
	}

	return q.queue[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.queue)
}

func (q *Queue[T]) Clear() {
	q.queue = []T{}
}
