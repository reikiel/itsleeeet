type Stack []int

func (s *Stack) Push(x int) {
    *s = append(*s, x)
}

func (s *Stack) Pop() (int, error) {
    if len(*s) == 0 {
        return 0, errors.New("stack is empty, cannot pop")
    }
    val := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return val, nil
}

func (s *Stack) Peek() int {
    return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
    return len(*s) == 0
}


type MyQueue struct {
    s1 Stack
    s2 Stack
}


func Constructor() MyQueue {
    return MyQueue{
        s1: make(Stack, 0),
        s2: make(Stack, 0),
    }
}


func (this *MyQueue) Push(x int)  {
    for !this.s2.Empty() {
        if val, err := this.s2.Pop(); err == nil {
            this.s1.Push(val)
        } else {
            fmt.Println("Error: ", err)
        }
    }
    this.s1.Push(x)
}


func (this *MyQueue) Pop() int {
    if len(this.s1) == 0 {
        fmt.Println("Pop: Queue is empty")
    }

    for !this.s1.Empty() {
        if val, err := this.s1.Pop(); err == nil {
            this.s2.Push(val)
        } else {
            fmt.Println("Error: ", err)
        }
    }

    popped, _ := this.s2.Pop()

    return popped
}


func (this *MyQueue) Peek() int {
        if len(this.s1) == 0 {
        fmt.Println("Peek: Queue is empty")
    }

    for !this.s1.Empty() {
        if val, err := this.s1.Pop(); err == nil {
            this.s2.Push(val)
        } else {
            fmt.Println("Error: ", err)
        }
    }

    return this.s2.Peek()
}


func (this *MyQueue) Empty() bool {
    return this.s1.Empty() && this.s2.Empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */