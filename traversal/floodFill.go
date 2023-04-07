package traversal

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	return floodFill(image, sr, sc, color)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	dfsFloodFill(image, sr, sc, image[sr][sc], color)
	return image
}

func dfsFloodFill(image [][]int, x, y, orgColor, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[x]) || image[x][y] == newColor || image[x][y] != orgColor {
		return
	}
	image[x][y] = newColor
	dfsFloodFill(image, x+1, y, orgColor, newColor)
	dfsFloodFill(image, x-1, y, orgColor, newColor)
	dfsFloodFill(image, x, y+1, orgColor, newColor)
	dfsFloodFill(image, x, y-1, orgColor, newColor)

}
