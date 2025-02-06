func replaceElements(arr []int) []int {
    greatest := arr[len(arr)-1]

    for i := len(arr)-1 ; i >= 0 ; i-- {
        temp := arr[i]
        arr[i] = greatest
        if temp > greatest {
            greatest = temp
        }
    }
     arr[len(arr)-1] = -1
     return arr
}