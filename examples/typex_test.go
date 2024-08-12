package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx/utilx"
	"testing"
	"time"
)

func TestBasicType(t *testing.T) {
	fmt.Println(typex.Flt(1.12).Number())
	fmt.Println(typex.Flt(1.5).UNumber())
	fmt.Println(typex.Flt(1.543543101).String()) //1.5435431
}

// LightPink #FFB6C1	255,182,193
func TestColor(t *testing.T) {
	fmt.Println(typex.NewRgbColor(255, 182, 193, 25))
}
func TestDuration(t *testing.T) {
	fmt.Println(typex.TimeZero)
	duration := 3*time.Second + 500*time.Millisecond //3.5s
	fmt.Println(typex.TimeDuration(duration, time.Millisecond))
	fmt.Println(typex.TimeDurationSecond(duration))
	fmt.Println(typex.TimeDurationSecondI(duration))
	fmt.Println(typex.TimeDurationParseSecondF("200ms"))
	fmt.Println(typex.TimeDurationParseSecondF("200000us"))
	fmt.Println(typex.TimeDurationParseSecondF("00:01:01"))
	fmt.Println(typex.TimeDurationParseSecondF("61s"))
}
func TestUtilTime(t *testing.T) {
	//TimeParseFormat := "15:04:05.01"
	s := "00:00:36.13"
	parse, err := time.Parse("15:04:05.00", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("parse", parse.Sub(typex.TimeZero))
	fmt.Println("parse", utilx.DurToSec(s))
}
func TestVideoParams(t *testing.T) {
	fmt.Println(typex.NewVideoSize(1024, 720))
	fmt.Println(typex.NewRatio(1024, 720))
	fmt.Println(typex.NewRate(25, 1))
}
