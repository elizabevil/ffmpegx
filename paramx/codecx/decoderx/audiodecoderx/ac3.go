package audiodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// AC3 5.1 ac3
type AC3 struct {
	DrcScale typex.UFloat32 `json:"drc_scale" flag:"-drc_scale,1"`
	/*drc_scale == 0
	DRC disabled. Produces full range audio.

	0 < drc_scale <= 1
	DRC enabled. Applies a fraction of the stream DRC value. Audio reproduction is between full range and full compression.

	drc_scale > 1
	DRC enabled. Applies drc_scale asymmetrically. Loud sounds are fully compressed. Soft sounds are enhanced.*/
}
