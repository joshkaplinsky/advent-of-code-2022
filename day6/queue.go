package main

type Queue struct {
	q     []rune
	count int
}

func NewQueue(i []rune) Queue {
	return Queue{
		q:     i,
		count: 14,
	}
}

func (q *Queue) Enqueue(i rune) {
	q.count++

	if len(q.q) == 14 {
		q.Dequeue()
	}

	q.q = append(q.q, i)
}

func (q *Queue) Dequeue() {
	q.q = q.q[1:]
}

func (q *Queue) Count() int {
	return q.count
}

func (q *Queue) IsQueueUnique() bool {
	unique := make(map[rune]bool)
	for _, i := range q.q {
		if unique[i] {
			return false
		}
		unique[i] = true
	}

	return true
}
