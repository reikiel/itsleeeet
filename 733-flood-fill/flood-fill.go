func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color { return image }
    ogCol := image[sr][sc]
    image[sr][sc] = color

    // up
    if sr > 0 && image[sr-1][sc] == ogCol {
        floodFill(image, sr-1, sc, color)
    }

    // down
    if sr < len(image)-1 && image[sr+1][sc] == ogCol {
        floodFill(image, sr+1, sc, color)
    }

    // left 
    if sc > 0 && image[sr][sc-1] == ogCol {
        floodFill(image, sr, sc-1, color)
    }

    // right
    if sc < len(image[0])-1 && image[sr][sc+1] == ogCol {
        floodFill(image, sr, sc+1, color)
    }


    return image
}