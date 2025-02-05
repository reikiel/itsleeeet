func checkIfExist(arr []int) bool {
    m := make(map[int]struct{})

    for _, v := range arr {
        if _, exists := m[v*2]; exists {
            return true
        }

        if v%2 == 0 {
            if _, exists := m[v/2]; exists {
                return true
            }
        }

        if _, exists := m[v]; !exists {
            m[v] = struct{}{}
        }
    }
    return false
}