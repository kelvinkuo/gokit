package arrayqueue

import (
    "github.com/emirpasic/gods/queues/arrayqueue"
    "github.com/kelvinkuo/gokit/structure/queues"
)

func NewArrayQueue() queues.Queue {
    return arrayqueue.New()
}
