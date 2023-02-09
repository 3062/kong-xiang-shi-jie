package queue

import (
	"errors"
)

type Queue[T any] struct {
	head, tail, len int
	contents        []T
}

func (q *Queue[T]) Empty() bool {
	return q.len == 0
}

func (q *Queue[T]) Size() int {
	return q.len
}

func (q *Queue[T]) Front() T {
	return q.contents[q.head]
}

func (q *Queue[T]) Back() T {
	return q.contents[q.tail]
}

func (q *Queue[T]) PushBack(content T) {
	if q.len == len(q.contents) {
		q.dilatation()
	}
	q.tail = (q.tail + 1) % len(q.contents)

	q.contents[q.tail] = content
}

func (q *Queue[T]) PopFront() (T, error) {
	ret := q.contents[q.head]
	if q.len == 0 {
		return ret, errors.New("")
	}

	q.head = (q.head + 1) % len(q.contents)
	q.len -= 1

	return ret, nil
}

func (q *Queue[T]) dilatation() {
	newContents := make([]T, 2*len(q.contents))
	newContents = append(newContents, q.contents[q.head:]...)
	newContents = append(newContents, q.contents[0:q.tail]...)
	q.head = 0
	q.tail = q.len - 1
	q.contents = newContents
}
