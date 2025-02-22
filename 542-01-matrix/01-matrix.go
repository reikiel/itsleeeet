func updateMatrix(mat [][]int) [][]int {
    q := [][2]int{}

    for i, row := range mat {

        for j, val := range row {
            if val != 0 {
                mat[i][j] = math.MaxInt64
            } else {
                // the 0s will be initialised to 0 alr
                // we want to start checking from 0s
                q = append(q, [2]int{i,j})
            }
        }
    }

    dir := [4][2]int{
        [2]int{1,0},
        [2]int{-1,0},
        [2]int{0,1},
        [2]int{0,-1},
    }

    for len(q) > 0 {
        // pop
        curr := q[0]
        q = q[1:]
        i,j := curr[0], curr[1]

        for _, d := range dir {
            d0, d1 := d[0], d[1]
            if i+d0 >= 0 && i+d0 < len(mat) && j+d1 >= 0 && j+d1 <len(mat[0]) && mat[i+d0][j+d1] > mat[i][j] + 1 {
                mat[i+d0][j+d1] = mat[i][j] + 1
                q = append(q, [2]int{i+d0,j+d1})
            }
            
        }
        
    }
    return mat

}