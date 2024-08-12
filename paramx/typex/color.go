package typex

import (
	"fmt"
)

//https://ffmpeg.org/ffmpeg-all.html#Color

/*
It can be the name of a color as defined below (case insensitive match) or a [0x|#]RRGGBB[AA] sequence, possibly followed by @ and a string representing the alpha component.
The alpha component may be a string composed by "0x" followed by an hexadecimal number or a decimal number between 0.0 and 1.0, which represents the opacity value (‘0x00’ or ‘0.0’ means completely transparent, ‘0xff’ or ‘1.0’ completely opaque). If the alpha component is not specified then ‘0xff’ is assumed.
The string ‘random’ will result in a random color.
*/

// [0x|#]RRGGBB[AA]
type Color = String
type ColorUI8 = UI8

func NewRgbColor(R ColorUI8, G ColorUI8, B ColorUI8, A ColorUI8) Color {
	return Color(fmt.Sprintf("#%X%X%X%X", R, G, B, A))
}

type Graph = Color
type Palette = Color
