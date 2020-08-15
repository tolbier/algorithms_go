package strassen

type Matrix []Row
type Row []int

func (x Row) Len() int {
	return len(x)
}
func (x Matrix) Len() int {
	return len(x)
}
func (x Matrix) Multiply(y Matrix) (res Matrix) {
	if x.Len() == 1 && y.Len() == 1 {
		return Matrix{{x[0][0] * y[0][0]}}
	}
	a, b, c, d := x.Split()
	e, f, g, h := y.Split()
	p1 := a.Multiply(f.Subtract(h))
	p2 := a.Sum(b).Multiply(h)
	p3 := c.Sum(d).Multiply(e)
	p4 := d.Multiply(g.Subtract(e))
	p5 := a.Sum(d).Multiply(e.Sum(h))
	p6 := b.Subtract(d).Multiply(g.Sum(h))
	p7 := a.Subtract(c).Multiply(e.Sum(f))

	return Join(
		p5.Sum(p4).Subtract(p2).Sum(p6),
		p1.Sum(p2),
		p3.Sum(p4),
		p1.Sum(p5).Subtract(p3).Subtract(p7),
	)
}
func (x Matrix) Sum(y Matrix) (res Matrix) {
	res = make(Matrix, x.Len())
	for j, row := range x {
		res[j] = make(Row, row.Len())
		for i, _ := range row {
			res[j][i] = x[j][i] + y[j][i]
		}
	}
	return
}
func (x Matrix) Subtract(y Matrix) (res Matrix) {
	res = make(Matrix, len(x))
	for j, row := range x {
		res[j] = make(Row, row.Len())
		for i := 0; i < row.Len(); i++ {
			res[j][i] = x[j][i] - y[j][i]
		}
	}
	return

}
func (x Matrix) Split() (a, b, c, d Matrix) {
	n := x.Len()
	m := x[0].Len()
	a = x.SubMatrix(0, m/2, 0, n/2)
	b = x.SubMatrix(m/2, m, 0, n/2)
	c = x.SubMatrix(0, m/2, n/2, n)
	d = x.SubMatrix(m/2, m, n/2, n)
	return
}

func Join(a, b, c, d Matrix) Matrix {
	ab := JoinHoriz(a, b)
	cd := JoinHoriz(c, d)
	return JoinVert(ab, cd)
}

func JoinVert(ab Matrix, cd Matrix) (res Matrix) {
	res = append(res, ab...)
	res = append(res, cd...)
	return
}

func JoinHoriz(a Matrix, b Matrix) (res Matrix) {
	res = make(Matrix, a.Len())
	for j, row := range a {
		res[j] = append(res[j], row...)
		res[j] = append(res[j], b[j]...)
	}
	return res
}
func (x Matrix) SubMatrix(startI, endI, startJ, endJ int) (res Matrix) {
	sJ := x[startJ:endJ]
	res = make(Matrix, sJ.Len())
	for j, _ := range sJ {
		res[j] = sJ[j][startI:endI]
	}
	return res
}
