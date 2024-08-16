package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

type Resampler struct {
	Uchl         string `json:"uchl" flag:"-uchl"`
	UsedChlayout string `json:"used_chlayout" flag:"-used_chlayout"`
	//Set used input channel layout. Default is unset. This option is only used for special remapping.

	Isr          typex.SampleRate `json:"isr" flag:"-isr"`
	InSampleRate typex.SampleRate `json:"in_sample_rate" flag:"-in_sample_rate"`
	//Set the input sample rate. Default value is 0.

	Osr           typex.SampleRate `json:"osr" flag:"-osr"`
	OutSampleRate typex.SampleRate `json:"out_sample_rate" flag:"-out_sample_rate"`
	//Set the output sample rate. Default value is 0.

	Isf         typex.SampleFmt `json:"isf" flag:"-isf"`
	InSampleFmt typex.SampleFmt `json:"in_sample_fmt" flag:"-in_sample_fmt"`
	//Specify the input sample format. It is set by default to none.

	Osf          typex.SampleFmt `json:"osf" flag:"-osf"`
	OutSampleFmt typex.SampleFmt `json:"out_sample_fmt" flag:"-out_sample_fmt"`
	//Specify the output sample format. It is set by default to none.

	Tsf               typex.SampleFmt `json:"tsf" flag:"-tsf"`
	InternalSampleFmt typex.SampleFmt `json:"internal_sample_fmt" flag:"-internal_sample_fmt"`
	//Set the internal sample format. Default value is none. This will automatically be chosen when it is not explicitly set.

	Ichl        flagx.ChannelLayout `json:"ichl" flag:"-ichl"`
	InChlayout  flagx.ChannelLayout `json:"in_chlayout" flag:"-in_chlayout"`
	Ochl        flagx.ChannelLayout `json:"ochl" flag:"-ochl"`
	OutChlayout flagx.ChannelLayout `json:"out_chlayout" flag:"-out_chlayout"`
	//Set the input/output channel layout.

	//See (ffmpeg-utils)the Channel Layout section in the ffmpeg-utils(1) manual for the required syntax.

	Clev           typex.Flt `json:"clev" flag:"-clev"`
	CenterMixLevel typex.Flt `json:"center_mix_level" flag:"-center_mix_level"`
	//Set the center mix level. It is a value expressed in deciBe ,and must be in the interval [-32,32].

	Slev             typex.Flt `json:"slev" flag:"-slev"`
	SurroundMixLevel typex.Flt `json:"surround_mix_level" flag:"-surround_mix_level"`
	//Set the surround mix level. It is a value expressed in deciBel, and must be in the interval [-32,32].

	LfeMixLevel typex.Flt `json:"lfe_mix_level" flag:"-lfe_mix_level"`
	//Set LFE mix into non LFE level. It is used when there is a LFE input but no LFE output. It is a value expressed in deciBel, and must be in the interval [-32,32].

	Rmvol          typex.Flt `json:"rmvol" flag:"-rmvol"`
	RematrixVolume typex.Flt `json:"rematrix_volume" flag:"-rematrix_volume"`
	//Set rematrix volume. Default value is 1.0.

	RematrixMaxval typex.UFlt `json:"rematrix_maxval" flag:"-rematrix_maxval"`
	//Set maximum output value for rematrixing. This can be used to prevent clipping vs. preventing volume reduction. A value of 1.0 prevents clipping.

	Flags    typex.Flags `json:"flags" flag:"-flags"`
	SwrFlags typex.Flags `json:"swr_flags" flag:"-swr_flags"`
	//Set flags used by the converter. Default value is 0.

	//"res"`
	//force resampling, this flag forces resampling to be used even when the input and output sample rates match.

	DitherScale typex.ScaleI `json:"dither_scale" flag:"-dither_scale"`
	//Set the dither scale. Default value is 1.

	DitherMethod DitherMethod `json:"dither_method" flag:"-dither_method"`
	//Set dither method. Default value is 0.
	Resampler ResamplerEngine `json:"resampler" flag:"-resampler"`
	//Set resampling engine. Default value is swr.

	FilterSize typex.Size `json:"filter_size" flag:"-filter_size"`
	//For swr only, set resampling filter size, default value is 32.

	PhaseShift typex.UI8 `json:"phase_shift" flag:"-phase_shift"`
	//For swr only, set resampling phase shift, default value is 10, and must be in the interval [0,30].

	LinearInterp typex.UBool `json:"linear_interp" flag:"-linear_interp"`
	//Use linear interpolation when enabled (the default). Disable it if you want to preserve speed instead of quality when exact_rational fails.

	ExactRational typex.UBool `json:"exact_rational" flag:"-exact_rational"`
	//For swr only, when enabled, try to use exact phase_count based on input and output sample rate. However, if it is larger than 1 << phase_shift, the phase_count will be 1 << phase_shift as fallback. Default is enabled.

	Cutoff typex.UBool `json:"cutoff" flag:"-cutoff"`
	//Set cutoff frequency (swr: 6dB point; soxr: 0dB point) ratio; must be a float value between 0 and 1. Default value is 0.97 with swr, and 0.91 with soxr (which, with a sample-rate of 44100, preserves the entire audio band to 20kHz).

	Precision typex.Flt `json:"precision" flag:"-precision"`
	//For soxr only, the precision in bits to which the resampled signal will be calculated. The default value of 20 (which, with suitable dithering, is appropriate for a destination bit-depth of 16) gives SoX’s ’High Quality’; a value of 28 gives SoX’s ’Very High Quality’.

	Cheby typex.UBool `json:"cheby" flag:"-cheby"`
	//For soxr only, selects passband rolloff none (Chebyshev) & higher-precision approximation for ’irrational’ ratios. Default value is 0.

	Async typex.Number `json:"async" flag:"-async"`
	//For swr only, simple 1 parameter audio sync to timestamps using stretching, squeezing, filling and trimming. Setting this to 1 will enable filling and trimming, larger values represent the maximum amount in samples that the data may be stretched or squeezed for each second. Default value is 0, thus no compensation is applied to make the samples match the audio timestamps.

	FirstPts typex.Number `json:"first_pts" flag:"-first_pts"`
	//For swr only, assume the first pts should be this value. The time unit is 1 / sample rate. This allows for padding/trimming at the start of stream. By default, no assumption is made about the first frame’s expected pts, so no padding or trimming is done. For example, this could be set to 0 to pad the beginning with silence if an audio stream starts after the video stream or to trim any samples with a negative pts due to encoder delay.

	MinComp typex.SecondF `json:"min_comp" flag:"-min_comp"`
	//For swr only, set the minimum difference between timestamps and audio data (in seconds) to trigger stretching/squeezing/filling or trimming of the data to make it match the timestamps. The default is that stretching/squeezing/filling and trimming is disabled (min_comp = FLT_MAX).

	MinHardComp typex.SecondF `json:"min_hard_comp" flag:"-min_hard_comp"`
	//For swr only, set the minimum difference between timestamps and audio data (in seconds) to trigger adding/dropping samples to make it match the timestamps. This option effectively is a threshold to select between hard (trim/fill) and soft (squeeze/stretch) compensation. Note that all compensation is by default disabled through min_comp. The default is 0.1.

	CompDuration typex.SecondF `json:"comp_duration" flag:"-comp_duration"`
	//For swr only, set duration (in seconds) over which data is stretched/squeezed to make it match the timestamps. Must be a non-negative double float value, default value is 1.0.

	MaxSoftComp typex.UFlt `json:"max_soft_comp" flag:"-max_soft_comp"`
	//For swr only, set maximum factor by which data is stretched/squeezed to make it match the timestamps. Must be a non-negative double float value, default value is 0.

	MatrixEncoding MatrixEncoding `json:"matrix_encoding" flag:"-matrix_encoding"`
	//Select matrixed stereo encoding.
	FilterType FilterType `json:"filter_type" flag:"-filter_type"`
	//For swr only, select resampling filter type. This only affects resampling operations.
	KaiserBeta typex.Flt `json:"kaiser_beta" flag:"-kaiser_beta"`
	//For swr only, set Kaiser window beta value. Must be a double float value in the interval [2,16], default value is 9.

	OutputSampleBits typex.SampleBits `json:"output_sample_bits" flag:"-output_sample_bits"`
	//For swr only, set number of used output sample bits for dithering. Must be an integer in the interval [0,64], default value is 0, which means it’s not used.
}

type DitherMethod = typex.Flags

const (
	DitherMethod_rectangular DitherMethod = "rectangular"
	//select rectangular dither

	DitherMethod_triangular DitherMethod = "triangular"
	//select triangular dither

	DitherMethod_triangular_hp DitherMethod = "triangular_hp"
	//select triangular dither with high pass

	DitherMethod_lipshitz DitherMethod = "lipshitz"
	//select Lipshitz noise shaping dither.

	DitherMethod_shibata DitherMethod = "shibata"
	//select Shibata noise shaping dither.

	DitherMethod_low_shibata DitherMethod = "low_shibata"
	//select low Shibata noise shaping dither.

	DitherMethod_high_shibata DitherMethod = "high_shibata"
	//select high Shibata noise shaping dither.

	DitherMethod_f_weighted DitherMethod = "f_weighted"
	//select f-weighted noise shaping dither

	DitherMethod_modified_e_weighted DitherMethod = "modified_e_weighted"
	//select modified-e-weighted noise shaping dither

	DitherMethod_improved_e_weighted DitherMethod = "improved_e_weighted"
	//select improved-e-weighted noise shaping dither
)

// Set resampling engine. Default value is swr.
type ResamplerEngine = typex.String

const (
	Resampler_swr ResamplerEngine = "swr"
	//select the native SW Resampler; filter options precision and cheby are not applicable in this case.

	Resampler_soxr ResamplerEngine = "soxr"
	//select the SoX Resampler (where available); compensation, and filter options filter_size, phase_shift, exact_rational, filter_type & kaiser_beta, are not applicable in this case.
)

type MatrixEncoding = typex.String
type FilterType = typex.String

const (
	MatrixEncoding_none MatrixEncoding = "none"
	//select none

	MatrixEncoding_dolby MatrixEncoding = "dolby"
	//select Dolby

	MatrixEncoding_dplii MatrixEncoding = "dplii"
	//select Dolby Pro Logic II
)
const (
	FilterType_cubic FilterType = "cubic"
	//select cubic

	FilterType_blackman_nuttall FilterType = "blackman_nuttall"
	//select Blackman Nuttall windowed sinc

	FilterType_kaiser FilterType = "kaiser"
	//select Kaiser windowed sinc
)
