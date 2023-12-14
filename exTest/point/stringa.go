package point

import (
  "fmt"
)
func Str(p Point) string {
  return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}
