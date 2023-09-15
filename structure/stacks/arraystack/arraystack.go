package arraystack

import (
    "github.com/emirpasic/gods/stacks/arraystack"
    "github.com/kelvinkuo/gokit/structure/stacks"
)

func NewArrayStack() stacks.Stack {
    return arraystack.New()
}
