package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func Qu() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(5)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
