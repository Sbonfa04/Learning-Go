package point

func Median(p1, p2 Point) (res Point) {
  res.X = (p1.X + p2.X) / 2
  res.Y = (p1.Y + p2.Y) / 2
  return
}
