package main

import "fmt"

type Queue struct {
	queue []int
	size  int
}

func (q *Queue) init(s int) {
	q.queue = make([]int, 0)
	q.size = s
}

func (q *Queue) enqueue(i int) error {
	if len(q.queue) < q.size {
		q.queue = append(q.queue, i)
		return nil
	}
	return fmt.Errorf("Queue is full.")
}

func (q *Queue) dequeue() (int, error) {
	if len(q.queue) == 0 {
		return -1, fmt.Errorf("Queue is empty")
	}
	value := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return value, nil
}

func QueueDriver() {
	var q = new(Queue)
	q.init(5)
	for i := 1; i <= 6; i++ {
		err := q.enqueue(i * i)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Println(q.queue)
	for i := 0; i < q.size+1; i++ {
		value, err := q.dequeue()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(value)
	}
	fmt.Println(q.queue)
}
