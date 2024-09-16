func removeCoveredIntervals(intervals [][]int) int {
    // sort first element by increasing order
    // sort second element by decreasing order
    // this way we know that the first element of later elements
    // are always covered by first element of previous elements
    sort.Slice(intervals, func(i,j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    i, count := 0, 0
    for j := 1; j < len(intervals) ; j++ {
        if intervals[i][1] < intervals[j][1] {
            i = j
            continue
        }
        count += 1
    }

    return len(intervals) - count
    
}