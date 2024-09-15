func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    if image[sr][sc] == color { return image }
    startCol := image[sr][sc]
    image[sr][sc] = color

    // up
    if sr > 0 && image[sr-1][sc] == startCol {
        floodFill(image, sr-1, sc, color)
    }

    if sr < len(image)-1 && image[sr+1][sc] == startCol {
        floodFill(image, sr+1, sc, color)
    }

    if sc > 0 && image[sr][sc-1] == startCol {
        floodFill(image, sr, sc-1, color)
    }

    if sc < len(image[0])-1 && image[sr][sc+1] == startCol {
        floodFill(image, sr, sc+1, color)
    }

    return image
}