package point

func NewPoint(x, y float64) Point {
  var p Point
  p.X, p.Y = x, y
  return p
}
