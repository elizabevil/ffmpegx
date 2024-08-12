package videodecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBDAV1D 4.3 libdav1d
type LIBDAV1D struct {
	//framethreads
	//Set amount of frame threads to use during decoding. The default value is 0 (autodetect). This option is deprecated for libdav1d >= 1.0 and will be removed in the future. Use the option max_frame_delay and the global option threads instead.

	//tilethreads
	//Set amount of tile threads to use during decoding. The default value is 0 (autodetect). This option is deprecated for libdav1d >= 1.0 and will be removed in the future. Use the global option threads instead.

	MaxFrameDelay typex.Delay `json:"max_frame_delay" flag:"-max_frame_delay"`
	//Set max amount of frames the decoder may buffer internally. The default value is 0 (autodetect).

	//Filmgrain typex.Delay `json:"filmgrain" flag:"-filmgrain"`
	//Apply film grain to the decoded video if present in the bitstream. Defaults to the internal default of the library. This option is deprecated and will be removed in the future. See the global option export_side_data to export Film Grain parameters instead of applying it.

	Oppoint typex.Delay `json:"oppoint" flag:"-oppoint"`
	//Select an operating point of a scalable AV1 bitstream (0 - 31). Defaults to the internal default of the library.

	Alllayers typex.UBool `json:"alllayers" flag:"-alllayers"`
	//Output all spatial layers of a scalable AV1 bitstream. The default value is false.

}
