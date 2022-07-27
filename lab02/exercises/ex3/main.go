package ex3

type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) getName() string {
	return "rectangle"
}

type Circle struct {
	Radius int
}

func (c *Circle) getName() string {
	return "circle"
}

type Square struct {
	Side int
}

func (s *Square) getName() string {
	return "square"
}

func (s Shape) visitForRectangle(r *Rectangle) {

}

type ShapeVisitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Shape interface {
	getName() string
	accept(ShapeVisitor)
}
