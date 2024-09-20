type MyHashMap struct {
    Data [1000001]int
    Keys [1000001]bool
}


func Constructor() MyHashMap {
    return MyHashMap {
        Data: [1000001]int{},
        Keys: [1000001]bool{},
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    this.Data[key] = value
    this.Keys[key] = true
}


func (this *MyHashMap) Get(key int) int {
    if this.Keys[key] {
        return this.Data[key]
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    this.Keys[key] = false
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */