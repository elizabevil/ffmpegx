package outputdevicex

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// // DECKLINK 4.4 decklink
type DECKLINK struct {
	ListDevices typex.BoolStr `json:"list_devices" flag:"-list_devices"`
	//If set to true, print a list of devices and exit. Defaults to false. This option is deprecated, please use the -sinks option of ffmpeg to list the available output devices.

	ListFormats typex.Bool `json:"list_formats" flag:"-list_formats"`
	//If set to true, print a list of supported formats and exit. Defaults to false.

	Preroll typex.SecondI `json:"preroll" flag:"-preroll"`
	//Amount of time to preroll video in seconds. Defaults to 0.5.

	DuplexMode   flagx.DuplexMode `json:"duplex_mode" flag:"-duplex_mode"`
	TimingOffset string           `json:"timing_offset" flag:"-timing_offset"`
	//Sets the genlock timing pixel offset on the used output. Defaults to ‘unset’.

	Link Link `json:"link" flag:"-link"`
	//Sets the SDI video link configuration on the used output. Must be ‘unset’, ‘single’ link SDI, ‘dual’ link SDI or ‘quad’ link SDI. Defaults to ‘unset’.

	Sqd typex.BoolStr `json:"sqd" flag:"-sqd"`
	//Enable Square Division Quad Split mode for Quad-link SDI output. Must be ‘unset’, ‘true’ or ‘false’. Defaults to unset.

	LevelA typex.BoolStr `json:"level_a" flag:"-level_a"`
	//Enable SMPTE Level A mode on the used output. Must be ‘unset’, ‘true’ or ‘false’. Defaults to unset.

	VancQueueSize typex.Size `json:"vanc_queue_size" flag:"-vanc_queue_size"`
	//Sets maximum output buffer size in bytes for VANC data. If the buffering reaches this value, outgoing VANC data will be dropped. Defaults to ‘1048576’.
}
type Link string

const (
	Link_unset  Link = "unset"
	Link_single Link = "single"
	Link_dual   Link = "dual"
	Link_quad   Link = "quad"
)

/*

List output devices:
ffmpeg -sinks decklink
List supported formats:
ffmpeg -i test.avi -f decklink -list_formats 1 'DeckLink Mini Monitor'
Play video clip:
ffmpeg -i test.avi -f decklink -pix_fmt uyvy422 'DeckLink Mini Monitor'
Play video clip with non-standard framerate or video size:
ffmpeg -i test.avi -f decklink -pix_fmt uyvy422 -s 720x486 -r 24000/1001 'DeckLink Mini Monitor'
4.5 fbdev
*/
