package utils

type node[T any] struct {
	val  T
	next *node[T]
}

type Queue[T any] struct {
	maxSize int
	head    *node[T]
	tail    *node[T]
	size    int
}

func (q Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) Push(elem T) {
	newNode := node[T]{
		val: elem,
	}
	if q.head == nil {
		q.head = &newNode
		q.tail = &newNode
	} else {
		q.head.next = &newNode
		q.head = &newNode
		if q.size == q.maxSize {
			q.Pop()
		}
	}

	q.size += 1
}

func (q *Queue[T]) Pop() T {
	prev := q.tail
	q.tail = q.tail.next
	q.size -= 1
	return prev.val
}

type Iterator[T any] struct {
	pointer *node[T]
}

func (i *Iterator[T]) Next() T {
	v := i.pointer
	i.pointer = i.pointer.next
	return v.val
}

func (i *Iterator[T]) HasNext() bool {
	return i.pointer != nil
}

func (q *Queue[T]) Iterator() Iterator[T] {
	return Iterator[T]{pointer: q.tail}
}

func NewQueue[T any](size int) Queue[T] {
	return Queue[T]{
		maxSize: size,
	}
}

func NewUnboundedQueue[T any]() Queue[T] {
	return Queue[T]{
		maxSize: -1,
	}
}
