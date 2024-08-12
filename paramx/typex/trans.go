package typex

import "strconv"

func (r Flt) Number() Number {
	return Number(r)
}
func (r Flt) UNumber() UNumber {
	return UNumber(r)
}
func (r Flt) String() String {
	return String(strconv.FormatFloat(float64(r), 'f', -1, 32))
}
