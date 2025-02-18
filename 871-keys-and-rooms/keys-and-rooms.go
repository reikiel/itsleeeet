func canVisitAllRooms(rooms [][]int) bool {
    num := len(rooms)
    q := []int{}
    visited := make([]int, num)

    // put room 0 keys in queue
    for _, key := range rooms[0] {
        q = append(q, key)
    }
    visited[0] = 1
    num -= 1

    for len(q) > 0 {
        // pop front
        curr := q[0]
        q = q[1:]

        if visited[curr] == 0 {
            visited[curr] = 1
            num -= 1
            for _, key := range rooms[curr] {
                q = append(q, key)
            }
        }
    }

    return num == 0
}