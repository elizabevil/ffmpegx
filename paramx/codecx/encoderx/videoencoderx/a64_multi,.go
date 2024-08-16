package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// A64_MULTI, 9.1 a64_multi, a64_multi5
type A64_MULTI struct {
	G typex.Size `json:"g" flag:"-g"` // (encoding, video)

	Q                    typex.Quality `json:"q" flag:"-q"`
	MaxExtraCbIterations typex.UNumber `json:"max_extra_cb_iterations" flag:"-max_extra_cb_iterations"`
	//Max extra codebook recalculation passes, more is better and slower.

	SkipEmptyCb typex.UBool `json:"skip_empty_cb" flag:"-skip_empty_cb"`
	//Avoid wasting bytes, ignore vintage MacOS decoder.

	MaxStrips typex.Strips `json:"max_strips" flag:"-max_strips"`
	MinStrips typex.Strips `json:"min_strips" flag:"-min_strips"`
	//The minimum and maximum number of strips to use. Wider range sometimes improves quality. More strips is generally better quality but costs more bits. Fewer strips tend to yield more keyframes. Vintage compatible is 1..3.

	StripNumberAdaptivity typex.Strips `json:"strip_number_adaptivity" flag:"-strip_number_adaptivity"`
	//How much number of strips is allowed to change between frames. Higher is better but slower.
}
