package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.16 argo_cvg ARGO_CVG
type ArgoCvg struct {
	SkipRateCheck typex.UBool `json:"skip_rate_check" flag:"-skip_rate_check"`
	//skip sample rate check (default is false)

	Loop typex.Count `json:"loop" flag:"-loop"`
	//set loop flag (default is false)

	Reverb *typex.UBool `json:"reverb" flag:"-reverb,1"`
	//set reverb flag (default is true)
}
