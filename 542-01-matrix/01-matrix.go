func updateMatrix(mat [][]int) [][]int {
    // res := make([][]int, len(mat))
    // for i:= range res {
    //     res[i] = make([]int, len(mat[0]))
    // }

    rows, cols := len(mat), len(mat[0])

    q := [][2]int{}

    for i, row := range mat {
        for j, val := range row {
            if val == 0 {
                q = append(q, [2]int{i,j})
            } else {
                mat[i][j] = math.MaxInt32
            }
        }
    }

    for len(q) > 0 {
        // pop from queue
        curr := q[0]
        q = q[1:]

        i, j := curr[0], curr[1]

        // see for each direction
        // if curr + 1 is < than comparing, then assign comparing to curr+1 and add to queue
        // if not means there was already a smaller distance found
        if i > 0 && mat[i-1][j] > mat[i][j] + 1 {
            mat[i-1][j] = mat[i][j] + 1
            q = append(q, [2]int{i-1,j})
        }

        if i < rows-1 && mat[i+1][j] > mat[i][j] + 1 {
            mat[i+1][j] = mat[i][j] + 1
            q = append(q, [2]int{i+1,j})
        }

        if j > 0 && mat[i][j-1] > mat[i][j] + 1 {
            mat[i][j-1] = mat[i][j] + 1
            q = append(q, [2]int{i,j-1})
        }

        if j < cols-1 && mat[i][j+1] > mat[i][j] + 1 {
            mat[i][j+1] = mat[i][j] + 1
            q = append(q, [2]int{i,j+1})
        }
    }
    return mat
}