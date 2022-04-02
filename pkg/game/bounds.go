package game

func withinBounds(x, y int, left, top, right, bottom int) bool {
	return x >= left && x <= right && y >= top && y <= bottom
}

type rectangle struct {
	x      int
	y      int
	width  int
	height int
}

func (r rectangle) withinBounds(x int, y int) bool {
	return x >= r.x && x <= r.x+r.width && y >= r.y && y <= r.y+r.height
}
