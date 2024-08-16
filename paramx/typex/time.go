package typex

import (
	"strings"
	"time"
)

// https://ffmpeg.org/ffmpeg-all.html#Time-duration
type Time = Number
type UTime = UNumber
type Delay = UI16

type TTL = Time
type Timeout = Time
type Delta = Time

type SecondF = Flt
type SecondI = UTime
type TimeShift = UTime

type MillisecondI = SecondI
type MicrosecondI = SecondI

type MillisecondF = SecondF
type MicrosecondF = SecondF
type CentiSeconds = UI16

type U16Offset = UI16 //65535
type Coordinate = UNumber
type TimeBase = String
type Date = String

// 00:02:00 or 120.5(s)
type Duration = Flt    //second
type DurationI = UTime //second
type Position = Duration

const (
	TimeParseFormat     = "15:04:05"
	TimeParseZeroFormat = "00:00:00"
)

var (
	//TimeZero, _ = time.Parse(TimeParseFormat, TimeParseZeroFormat)
	TimeZero = time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
)

// TimeDuration go time.Duration->spec time.Duration
func TimeDuration(duration time.Duration, spec time.Duration) Duration {
	return Duration(duration) / Duration(spec)
}

// TimeDurationParse string to  time.Duration
func TimeDurationParse(dur string) time.Duration {
	if strings.Contains(dur, ":") {
		parse, err := time.Parse(TimeParseFormat, dur)
		if err != nil {
			return 0
		}
		return parse.Sub(TimeZero)
	}
	duration, err := time.ParseDuration(dur)
	if err != nil {
		return 0
	}
	return duration
}

// TimeDurationParseSecondF 55  must end with s ==> 55s
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func TimeDurationParseSecondF(dur string) SecondF {
	return TimeDurationSecond(TimeDurationParse(dur))
}

// TimeDurationParseSecondI string to  SecondI
func TimeDurationParseSecondI(dur string) SecondI {
	return TimeDurationSecondI(TimeDurationParse(dur))
}

// TimeDurationSecond go time.Duration->SecondF
func TimeDurationSecond(duration time.Duration) SecondF {
	return TimeDuration(duration, time.Second)
}

// TimeDurationSecondI go time.Duration->SecondI
func TimeDurationSecondI(duration time.Duration) SecondI {
	return SecondI(TimeDuration(duration, time.Second))
}

//7.3 Time duration
//There are two accepted syntaxes for expressing time duration.
/*
HH expresses the number of hours,
MM the number of minutes for a maximum of 2 digits,
and SS the number of seconds for a maximum of 2 digits.
The m at the end expresses decimal value for SS.
	[-][HH:]MM:SS[.m...]
*/

/*
S expresses the number of seconds, with the optional decimal part m.
The optional literal suffixes ‘s’, ‘ms’ or ‘us’ indicate to interpret the value as seconds, milliseconds or microseconds, respectively.
	[-]S+[.m...][s|ms|us]
*/
/*
 Examples
‘55’
55 seconds ‘23.189’ ‘0.2’ ‘200ms’ ‘200000us’

‘12:03:45’
12 hours, 03 minutes and 45 seconds
*/

/*
The accepted syntax is:

[(YYYY-MM-DD|YYYYMMDD)[T|t| ]]((HH:MM:SS[.m...]]])|(HHMMSS[.m...]]]))[Z]
now

*/
