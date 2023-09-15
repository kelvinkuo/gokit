package stacks

import (
    "github.com/kelvinkuo/gokit/structure/container"
)

type Stack interface {
    container.Container
    
    Push(value interface{})
    Pop() (value interface{}, ok bool)
    Peek() (value interface{}, ok bool)
}
