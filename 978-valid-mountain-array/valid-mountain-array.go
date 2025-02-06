func validMountainArray(arr []int) bool {
    inc := true

    if len(arr) < 3 || arr[1] <= arr[0] {
        return false
    }

    for i:=1; i<len(arr); i++ {
        if arr[i] < arr[i-1] {
            inc = false
        }
        
        if (inc && arr[i] <= arr[i-1]) || (!inc && arr[i] >= arr[i-1]) {
            return false
        }
    }

    return true && !inc
}