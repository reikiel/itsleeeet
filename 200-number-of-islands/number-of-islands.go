func numIslands(grid [][]byte) int {
    res := 0
    for i, row := range grid {
        for j, val := range row {
            if val == '1' {
                res++
                markVisit(grid,i,j)
            }
        }
    }
    return res
}

func markVisit(grid [][]byte, i int, j int) {
    grid[i][j] = '0'

    if i > 0 && grid[i-1][j] == '1' {
        markVisit(grid, i-1, j)
    }

    if i < len(grid)-1 && grid[i+1][j] == '1' {
        markVisit(grid, i+1, j)
    }

    if j > 0 && grid[i][j-1] == '1' {
        markVisit(grid, i, j-1)
    }

    if j < len(grid[0])-1 && grid[i][j+1] == '1' {
        markVisit(grid, i, j+1)
    }
}