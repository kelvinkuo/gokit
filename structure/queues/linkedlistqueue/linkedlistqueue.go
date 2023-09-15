package linkedlistqueue

import (
    "github.com/emirpasic/gods/queues/linkedlistqueue"
    "github.com/kelvinkuo/gokit/structure/queues"
)

func NewLinkedListQueue() queues.Queue {
    return linkedlistqueue.New()
}
