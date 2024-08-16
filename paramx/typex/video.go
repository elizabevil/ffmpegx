package typex

import (
	"fmt"
)

type FrameRate = UNumber
type SampleRate = FrameRate
type SampleBits = UI16

type Depth = Number
type DepthF = Flt
type Cutoff = Number
type LastPred = Number
type Rotation = I16 //-360~360

type FrameSize = Size
type FrameNumber = Number
type SliceSize = Size
type DiaSize = Size

type Quality = Flt

type Quantizer = UI16 //1024
type RateControl = UI16

type Charset = String
type Antialias = String
type Algorithm = String
type Driver = String
type Frequency = String

type Iblock = Flt

type VideoSize = String

// NewVideoSize widthxheight
func NewVideoSize(width, height int) VideoSize {
	return VideoSize(fmt.Sprintf("%dx%d", height, width))
}

type Aspect = String //(4:3, 16:9 or 1.3333, 1.7777)
type Ratio = Aspect
type Rate = String
type RateI = UNumber

func NewRatio(width, height int) Ratio {
	return Ratio(fmt.Sprintf("%d:%d", height, width))
}

func NewRate(num, den int) Rate {
	return Rate(fmt.Sprintf("%d/%d", num, den))
}
