package cache

type Queue struct {
	head *QueueItem
	tail *QueueItem
}

// Move the queue and return the key to be removed (head of the queue)
func (q *Queue) dropHead() string {
	h := q.head
	q.head = q.head.next
	return h.val
}

func (q *Queue) add(next string) {
	newItem := QueueItem{
		val:  next,
		next: nil,
	}
	if q.tail == nil {
		// Empty queue
		q.tail = &newItem
		q.head = &newItem
	} else {
		oldTail := q.tail
		oldTail.next = &newItem
		q.tail = &newItem
	}
}

func EmptyQueue() Queue {
	return Queue{
		head: nil,
		tail: nil,
	}
}

type QueueItem struct {
	val  string
	next *QueueItem
}
