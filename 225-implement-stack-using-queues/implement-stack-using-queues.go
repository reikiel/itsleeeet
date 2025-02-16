type MyStack struct {
    q1 []int
    q2 []int
}


func Constructor() MyStack {
    return MyStack{
        q1: []int{},
        q2: []int{},
    }
}


func (this *MyStack) Push(x int)  {
    if len(this.q2) != 0 {
        this.q2 = append(this.q2, x)
        return  
    }
    this.q1 = append(this.q1, x) 
}


func (this *MyStack) Pop() int {
    if len(this.q2) != 0 {
        for len(this.q2) > 1 {
            this.q1 = append(this.q1, this.q2[0])
            this.q2 = this.q2[1:]
        }
        x := this.q2[0]
        this.q2 = this.q2[:0]
        return x
    }

    // else do for q1
    for len(this.q1) > 1 {
        this.q2 = append(this.q2, this.q1[0])
        this.q1 = this.q1[1:]
    }
    x := this.q1[0]
    this.q1 = this.q1[:0]
    return x
}


func (this *MyStack) Top() int {
    if len(this.q2) != 0 {
        for len(this.q2) > 1 {
            this.q1 = append(this.q1, this.q2[0])
            this.q2 = this.q2[1:]
        }
        x := this.q2[0]
        this.q1 = append(this.q1, this.q2[0])
        this.q2 = this.q2[:0]
        return x
    }

    // else do for q1
    for len(this.q1) > 1 {
        this.q2 = append(this.q2, this.q1[0])
        this.q1 = this.q1[1:]
    }
    x := this.q1[0]
    this.q2 = append(this.q2, this.q1[0])
    this.q1 = this.q1[:0]
    return x
}


func (this *MyStack) Empty() bool {
    return len(this.q1) == 0 && len(this.q2) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */