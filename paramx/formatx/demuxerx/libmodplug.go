package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 20.15 libmodplug  LIBMODPLUG
type LIBMODPLUG struct {
	NoiseReduction typex.UBool `json:"noise_reduction" flag:"-noise_reduction"`
	//Apply a simple low-pass filter. Can be 1 (on) or 0 (off). Default is 0.

	ReverbDepth typex.Level `json:"reverb_depth" flag:"-reverb_depth"`
	//Set amount of reverb. Range 0-100. Default is 0.

	ReverbDelay typex.Level `json:"reverb_delay" flag:"-reverb_delay"`
	//Set delay in ms, clamped to 40-250 ms. Default is 0.

	BassAmount typex.Level `json:"bass_amount" flag:"-bass_amount"`
	//Apply bass expansion a.k.a. XBass or megabass. Range is 0 (quiet) to 100 (loud). Default is 0.

	BassRange typex.Level `json:"bass_range" flag:"-bass_range"`
	//Set cutoff i.e. upper-bound for bass frequencies. Range is 10-100 Hz. Default is 0.

	SurroundDepth typex.Level `json:"surround_depth" flag:"-surround_depth"`
	//Apply a Dolby Pro-Logic surround effect. Range is 0 (quiet) to 100 (heavy). Default is 0.

	SurroundDelay typex.Level `json:"surround_delay" flag:"-surround_delay"`
	//Set surround delay in ms, clamped to 5-40 ms. Default is 0.

	MaxSize typex.UNumber `json:"max_size" flag:"-max_size"`
	//The demuxer buffers the entire file into memory. Adjust this value to set the maximum buffer size, which in turn, acts as a ceiling for the size of files that can be read. Range is 0 to 100 MiB. 0 removes buffer size limit (not recommended). Default is 5 MiB.

	VideoStreamExpr typex.String `json:"video_stream_expr" flag:"-video_stream_expr"`
	//String which is evaluated using the eval API to assign colors to the generated video stream. Variables which can be used are x, y, w, h, t, speed, tempo, order, pattern and row.

	VideoStream typex.UBool `json:"video_stream" flag:"-video_stream"`
	//Generate video stream. Can be 1 (on) or 0 (off). Default is 0.

	VideoStreamW typex.LevelUI16 `json:"video_stream_w" flag:"-video_stream_w"`
	//Set video frame width in ’chars’ where one char indicates 8 pixels. Range is 20-512. Default is 30.

	VideoStreamH typex.LevelUI16 `json:"video_stream_h" flag:"-video_stream_h"`
	//Set video frame height in ’chars’ where one char indicates 8 pixels. Range is 20-512. Default is 30.

	VideoStreamPtxt *typex.UBool `json:"video_stream_ptxt" flag:"-video_stream_ptxt,1"`
	//Print metadata on video stream. Includes speed, tempo, order, pattern, row and ts (time in ms). Can be 1 (on) or 0 (off). Default is 1.

}
