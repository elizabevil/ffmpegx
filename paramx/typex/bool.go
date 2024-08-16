package typex

// Bool if default !=0 but 0 is one of options --> *Bool
// other --> Bool
type Bool = I8    // -1 ~ 2(
type UBool = Bool // > 0

// most rules only read
var (
	Auto  Bool = -1
	False Bool = 0 // *False or False
	True  Bool = 1
	Multi Bool = 2
)

type BoolStr = String
type BoolMode = String

// only read
var (
	BoolTrue  BoolStr = "true"
	BoolFalse BoolStr = "false"
)
var (
	BoolOn   BoolMode = "on"
	BoolOff  BoolMode = "off"
	BoolAuto BoolMode = "auto"
)
