func duplicateZeros(arr []int)  {
    p1,p2 := 0,0
    temp := make([]int, len(arr))
    copy(temp, arr)

    for p2 < len(arr) {
        arr[p2] = temp[p1]
        if temp[p1] == 0 && p2 < len(arr)-1 {
            p2++
            arr[p2] = 0
        }
        p2++
        p1++
    }
}