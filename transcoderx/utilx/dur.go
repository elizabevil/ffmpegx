package utilx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"time"
)

const (
	DurParse = "15:04:05.00"
)

//00:00:01.12 -->1.12s
func DurToSec(dur string) (sec float64) {
	parse, err := time.Parse(DurParse, dur)
	if err != nil {
		return 0
	}
	return parse.Sub(typex.TimeZero).Seconds()
}
