package videoencoderx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// VC2 9.30 vc2
type VC2 struct {
	B typex.BitrateInt `json:"b" flag:"-b"`

	//Sets target video bitrate. Usually that’s around 1:6 of the uncompressed video bitrate (e.g. for 1920x1080 50fps yuv422p10 that’s around 400Mbps). Higher values (close to the uncompressed bitrate) turn on lossless compression mode.

	FieldOrder typex.Level `json:"field_order" flag:"-field_order"`
	//Enables field coding when set (e.g. to tt - top field first) for interlaced inputs. Should increase compression with interlaced content as it splits the fields and encodes each separately.
	FieldOrderFlag flagx.FieldOrder `json:"field_order_flag" flag:"-field_order"` // (video) 	//Set/override the field order of the video. Possible values:

	WaveletDepth typex.Level `json:"wavelet_depth" flag:"-wavelet_depth"`
	//Sets the total amount of wavelet transforms to apply, between 1 and 5 (default). Lower values reduce compression and quality. Less capable decoders may not be able to handle values of wavelet_depth over 3.

	WaveletType WaveletType `json:"wavelet_type" flag:"-wavelet_type"`
	//Sets the transform type. Currently only 5_3 (LeGall) and 9_7 (Deslauriers-Dubuc) are implemented, with 9_7 being the one with better compression and thus is the default.

	SliceWidth  typex.UI16 `json:"slice_width" flag:"-slice_width"`
	SliceHeight typex.UI16 `json:"slice_height" flag:"-slice_height"`
	//Sets the slice size for each slice. Larger values result in better compression. For compatibility with other more limited decoders use slice_width of 32 and slice_height of 8.

	Tolerance typex.PercentF `json:"tolerance" flag:"-tolerance"`
	//Sets the undershoot tolerance of the rate control system in percent. This is to prevent an expensive search from being run.

	Qm Qm `json:"qm" flag:"-qm"`
	//Sets the quantization matrix preset to use by default or when wavelet_depth is set to 5
}
type Qm string
type WaveletType string

const (
	WaveletType_9_7          WaveletType = "9_7"
	WaveletType_5_3          WaveletType = "5_3"
	WaveletType_haar         WaveletType = "haar"
	WaveletType_haar_noshift WaveletType = "haar_noshift"
)

const (
	Qm_default Qm = "default"
	Qm_color   Qm = "color"
	Qm_flat    Qm = "flat"
)
