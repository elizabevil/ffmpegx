package typex

type String string
type Number int
type UNumber uint

type Flt float32
type UI16 uint16
type I16 int16
type UI8 uint8
type I8 int8

type StreamSpecifier struct {
	K string
	V string
}

type Args []string

func (a Args) Args() Args {
	return a
}
