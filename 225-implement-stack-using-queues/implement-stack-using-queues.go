/**
 * Do error handling for peek/pop?
 */

type Queue []int

func (q *Queue) push(x int) {
    *q = append(*q, x)
}

func (q *Queue) pop() int {
    res := (*q)[0]
    *q = (*q)[1:]
    return res
}

func (q *Queue) peek() int {
    return (*q)[0]
}

func (q *Queue) empty() bool {
    return len(*q) == 0
}

func (q *Queue) size() int {
    return len(*q)
}


type MyStack struct {
    q1 Queue
    q2 Queue
}


func Constructor() MyStack {
    return MyStack{
        q1: make(Queue, 0),
        q2: make(Queue, 0),
    }
}


func (this *MyStack) Push(x int)  {
    for !this.q2.empty() {
        x := this.q2.pop()
        this.q1.push(x)
    }

    this.q1.push(x)
}


func (this *MyStack) Pop() int {
    for this.q1.size() > 1 {
        x := this.q1.pop()
        this.q2.push(x)
    }

    y := this.q1.pop()
    
    for !this.q2.empty() {
        x := this.q2.pop()
        this.q1.push(x)
    }

    return y
}


func (this *MyStack) Top() int {
    for this.q1.size() > 1 {
        x := this.q1.pop()
        this.q2.push(x)
    }

    y := this.q1.pop()
    this.q2.push(y)

    for !this.q2.empty() {
        x := this.q2.pop()
        this.q1.push(x)
    }

    return y
}


func (this *MyStack) Empty() bool {
    return this.q1.empty() && this.q2.empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */