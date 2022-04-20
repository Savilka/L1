package point

import "math"

// Point структура описывающая координаты точки
type Point struct {
	x, y float64
}

func (p Point) getX() float64 {
	return p.x
}

func (p Point) getY() float64 {
	return p.y
}

// NewPoint возвращает новую точку с координатами x, y
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance находит расстояние между двумя точками
func Distance(a, b Point) float64 {
	ax := a.getX()
	ay := a.getY()
	bx := b.getX()
	by := b.getY()
	return math.Sqrt((ax-bx)*(ax-bx) + (ay-by)*(ay-by))
}
