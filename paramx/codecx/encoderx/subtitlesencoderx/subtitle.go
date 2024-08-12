package subtitlesencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Dvdsub struct {
	Palette typex.Palette `json:"palette" flag:"-palette"`

	EvenRowsFix typex.UBool `json:"even_rows_fix" flag:"-even_rows_fix"`
}
