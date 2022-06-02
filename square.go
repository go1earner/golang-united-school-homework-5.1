package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sqr Square) End() Point {
	side := int(sqr.a)
	return Point{
		x: sqr.start.x + side,
		y: sqr.start.y + side,
	}
}

func (sqr Square) Area() uint {
	return sqr.a * sqr.a
}

func (sqr Square) Perimeter() uint {
	return sqr.a * 4
}
