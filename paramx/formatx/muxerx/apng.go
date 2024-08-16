package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// APNG 21.10 alp
// 4.14 apng
type APNG struct {
	FinalDelay typex.CentiSeconds `json:"final_delay" flag:"-final_delay"`
	//Force       a delay expressed in seconds after the last frame of each repetition.Default value is 0.0.

	Plays typex.UBool `json:"plays" flag:"-plays"`
	//specify how many times to play the content, 0 causes an infinte loop, with 1 there is no loop
}
