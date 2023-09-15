package liststack

import (
    "github.com/emirpasic/gods/stacks/linkedliststack"
    "github.com/kelvinkuo/gokit/structure/stacks"
)

func NewListStack() stacks.Stack {
    return linkedliststack.New()
}
