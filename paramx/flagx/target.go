package flagx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type OutputTarget = typex.OutputTarget

const (
	OutputTarget_vcd  OutputTarget = "vcd"
	OutputTarget_svcd OutputTarget = "svcd"
	OutputTarget_dvd  OutputTarget = "dvd"
	OutputTarget_dv   OutputTarget = "dv"
	OutputTarget_dv50 OutputTarget = "dv50"
)

// strict integer (decoding/encoding,audio,video)

type Strict = Flag

const (
	Strict_very Strict = "very"
	//strictly conform to an older more strict version of the spec or reference software

	Strict_strict Strict = "strict"
	//strictly conform to all the things in the spec no matter what consequences

	Strict_normal     Strict = "normal"
	Strict_unofficial Strict = "unofficial"
	//allow unofficial extensions

	Strict_experimental Strict = "experimental"
	//allow non standardized experimental things, experimental (unfinished/work in progress/not well tested) decoders and encoders. Note: experimental decoders can pose a security risk, do not use this for decoding untrusted input.
)
