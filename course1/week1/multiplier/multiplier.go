package multiplier

import (
	"regexp"
	"strconv"
	"strings"
)

func Multiply(n, m string) string {
	n1, m1 := normalizeZeroes(n, m)
	if len(n1) == 1 {
		return digitMultiply(n1, m1)
	}
	a, b := splitHalves(n1)
	c, d := splitHalves(m1)
	p := sumString(a, b)
	q := sumString(c, d)
	ac := Multiply(a, c)
	bd := Multiply(b, d)
	pq := Multiply(p, q)
	adbc := substractString(pq, sumString(ac, bd))
	bp1 := bypow10(ac, 2*len(a))
	bp2 := bypow10(adbc, len(a))
	sum1 := sumString(bp2, bd)
	sum2 := sumString(bp1, sum1)
	return removeZeroes(sum2)
}

func sumString(s1 string, s2 string) (res string) {
	s1, s2 = normalizeZeroes(s1, s2)
	var rem int
	for i := len(s1) - 1; i >= 0; i-- {
		n2 := sureAtoi(string([]byte(s2)[i]))
		n1 := sureAtoi(string([]byte(s1)[i]))
		sum := n1 + n2 + rem
		res = strconv.Itoa(sum%10) + res
		rem = sum / 10
	}
	if rem > 0 {
		res = strconv.Itoa(rem) + res
	}
	return
}

func substractString(s1 string, s2 string) (res string) {
	s1, s2 = normalizeZeroes(s1, s2)
	var rem int
	for i := len(s1) - 1; i >= 0; i-- {
		n2 := sureAtoi(string([]byte(s2)[i]))
		n1 := sureAtoi(string([]byte(s1)[i]))
		if n2+rem > n1 {
			n1 += 10
		}
		substract := n1 - (n2 + rem)
		res = strconv.Itoa(substract%10) + res
		rem = n1 / 10
	}
	res = removeZeroes(res)
	return
}

func removeZeroes(s string) string {
	r := regexp.MustCompile(`0*(\d+)`)
	sm := r.FindStringSubmatch(s)
	return sm[1]
}
func bypow10(s1 string, n int) string {
	return s1 + strings.Repeat("0", n)
}
func digitMultiply(n string, m string) string {
	return strconv.Itoa(sureAtoi(n) * sureAtoi(m))
}
func sureAtoi(n string) (res int) {
	res, _ = strconv.Atoi(n)
	return
}
func splitHalves(s string) (s1, s2 string) {
	s = evenStringLength(s)
	n := len(s)
	b := []byte(s)
	hn := n / 2
	return string(b[:hn]), string(b[hn:])
}

func evenStringLength(s string) string {
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return s
}
func normalizeZeroes(s1, s2 string) (r1, r2 string) {
	r1, r2 = s1, s2
	n1, n2 := len(r1), len(r2)
	if n1 > n2 {
		r2 = strings.Repeat("0", n1-n2) + r2
	} else if n1 < n2 {
		r1 = strings.Repeat("0", n2-n1) + r1
	}
	return
}
