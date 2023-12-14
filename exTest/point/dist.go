package point

import (
  "math"
)

func Dist(p1, p2 Point) float64 {
  dx := p1.X - p2.X
  dy := p1.Y - p2.Y
  return math.Sqrt(dx * dx + dy * dy)
}
