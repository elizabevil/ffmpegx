package inputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// VIDEO4LINUX2 , 3.19 video4linux2, v4l2
type VIDEO4LINUX2 struct {
	Standard string `json:"standard" flag:"-standard"`
	//Set the standard. Must be the name of a supported standard. To get a list of the supported standards, use the list_standards option.

	Channel typex.Channels `json:"channel" flag:"-channel,-1"`
	//Set the input channel number. Default to -1, which means using the previously selected channel.

	VideoSize flagx.VideoSize `json:"video_size" flag:"-video_size"`

	//Set the video frame size. The argument must be a string in the form WIDTHxHEIGHT or a valid size abbreviation.

	PixelFormat flagx.PixelFormat `json:"pixel_format" flag:"-pixel_format"`
	//Select the pixel format (only valid for raw video input).

	InputFormat typex.Format `json:"input_format" flag:"-input_format"`
	//Set the preferred pixel format (for raw video) or a codec name. This option allows one to select the input format, when several are available.

	Framerate typex.FrameRate `json:"framerate" flag:"-framerate"`
	//Set the preferred video frame rate.
	FramerateFlag flagx.VideoRate `json:"framerate_flag" flag:"-framerate"`

	ListFormats ListFormats `json:"list_formats" flag:"-list_formats"`
	//List available formats (supported pixel formats, codecs, and frame sizes) and exit.

	//‘all’ 	Show all supported standards.
	ListStandards typex.UBool `json:"list_standards" flag:"-list_standards"`
	Timestamps    typex.Level `json:"timestamps" flag:"-timestamps"`
	Ts            typex.Level `json:"ts" flag:"-ts"`
	//Set type of timestamps for grabbed frames.

	UseLibv4L2 typex.UBool `json:"use_libv4l2" flag:"-use_libv4l2"`
	//Use libv4l2 (v4l-utils) conversion functions. Default is 0.
}

type ListFormats = typex.String

const (
	ListFormats_all ListFormats = "all"
	//Show all available (compressed and non-compressed) formats.

	ListFormats_raw ListFormats = "raw"
	//Show only raw video (non-compressed) formats.

	ListFormats_compressed ListFormats = "compressed"
	//Show only compressed formats.
)

type Ts = typex.String

const (
	Ts_default Ts = "default"
	//Show all available (compressed and non-compressed) formats.

	Ts_abs Ts = "abs"
	//Show only raw video (non-compressed) formats.

	Ts_mono2abs Ts = "mono2abs"
	//Show only compressed formats.
)
