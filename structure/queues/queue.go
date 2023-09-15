package queues

import (
    "github.com/kelvinkuo/gokit/structure/container"
)

type Queue interface {
    container.Container
    
    Enqueue(value interface{})
    Dequeue() (value interface{}, ok bool)
    Peek() (value interface{}, ok bool)
}
