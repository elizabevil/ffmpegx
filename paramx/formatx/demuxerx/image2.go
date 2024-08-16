package demuxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 20.13 image2  IMAGE2
type IMAGE2 struct {
	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the frame rate for the video stream. It defaults to 25.

	Loop typex.UBool `json:"loop" flag:"-loop"`
	//If set to 1, loop over the input. Default value is 0.

	PatternType PatternType `json:"pattern_type" flag:"-pattern_type"`
	//Select the pattern type used to interpret the provided filename.
	PixelFormat typex.UNumber `json:"pixel_format" flag:"-pixel_format"`
	//Set the pixel format of the images to read. If not specified the pixel format is guessed from the first image file in the sequence.

	StartNumber typex.UNumber `json:"start_number" flag:"-start_number"`
	//Set the index of the file matched by the image file pattern to start to read from. Default value is 0.

	StartNumberRange typex.UNumber `json:"start_number_range" flag:"-start_number_range"`
	//Set the index interval range to check when looking for the first image file in the sequence, starting from start_number. Default value is 5.

	TsFromFile typex.Level `json:"ts_from_file" flag:"-ts_from_file"`
	//If set to 1, will set frame timestamp to modification time of image file. Note that monotonity of timestamps is not provided: images go in the same order as without this option. Default value is 0. If set to 2, will set frame timestamp to the modification time of the image file in nanosecond precision.

	VideoSize typex.VideoSize `json:"video_size" flag:"-video_size"`
	//Set the video size of the images to read. If not specified the video size is guessed from the first image file in the sequence.

	ExportPathMetadata typex.UBool `json:"export_path_metadata" flag:"-export_path_metadata"`
}
type PatternType string

const (
	glob_sequence PatternType = "glob_sequence"
	glob          PatternType = "glob"
	sequence      PatternType = "sequence"
	none          PatternType = "none"
)

/*



	20.13.1 Examples
	Use ffmpeg for creating a video from the images in the file sequence img-001.jpeg, img-002.jpeg, ..., assuming an input frame rate of 10 frames per second:
	ffmpeg -framerate 10 -i 'img-%03d.jpeg' out.mkv
	As above, but start by reading from a file with index 100 in the sequence:*/
