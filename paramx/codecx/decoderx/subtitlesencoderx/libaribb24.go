package subtitledecoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBARIBB24 6.1 libaribb24
type LIBARIBB24 struct {
	Aribb24BasePath typex.Filename `json:"aribb24-base-path" flag:"-aribb24-base-path"`
	//Sets the base path for the libaribb24 library. This is utilized for reading of configuration files (for custom unicode conversions), and for dumping of non-text symbols as images under that location.

	//Unset by default.

	Aribb24SkipRubyText typex.Bool `json:"aribb24-skip-ruby-text" flag:"-aribb24-skip-ruby-text"`
	//Tells the decoder wrapper to skip text blocks that contain half-height ruby text.

	//Enabled by default.
}
