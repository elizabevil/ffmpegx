package optionx

import "github.com/elizabevil/ffmpegx/paramx/flagx"

// Generic These options are shared amongst the ff* tools.
type Generic struct {
	Version bool `json:"version" flag:"-version"`
	//Show version.

	Buildconf bool `json:"buildconf" flag:"-buildconf"`
	//Show the build configuration, one option per line.

	Formats bool `json:"formats" flag:"-formats"`
	//Show available formats (including devices).

	Demuxers bool `json:"demuxers" flag:"-demuxers"`
	//Show available demuxers.

	Muxers bool `json:"muxers" flag:"-muxers"`
	//Show available muxers.

	Devices bool `json:"devices" flag:"-devices"`
	//Show available devices.

	Codecs bool `json:"codecs" flag:"-codecs"`
	//Show all codecs known to libavcodec.
	//Note that the term ’codec’ is used throughout this documentation as a shortcut for what is more correctly called a media bitstream format.

	Decoders bool `json:"decoders" flag:"-decoders"`
	//Show available decoders.

	Encoders bool `json:"encoders" flag:"-encoders"`
	//Show all available encoders.

	Bsfs bool `json:"bsfs" flag:"-bsfs"`
	//Show available bitstream filters.

	Protocols bool `json:"protocols" flag:"-protocols"`
	//Show available protocols.

	Filters bool `json:"filters" flag:"-filters"`
	//Show available libavfilter filters.

	PixFmts bool `json:"pix_fmts" flag:"-pix_fmts"`
	//Show available pixel formats.

	SampleFmts bool `json:"sample_fmts" flag:"-sample_fmts"`
	//Show available sample formats.

	Layouts bool `json:"layouts" flag:"-layouts"`
	//Show channel names and standard channel layouts.

	Dispositions bool `json:"dispositions" flag:"-dispositions"`
	//Show stream dispositions.

	Colors bool `json:"colors" flag:"-colors"`
	//Show recognized color names.

	//ffmpeg -loglevel repeat+level+verbose -i input output
	Loglevel   flagx.Loglevel `json:"loglevel" flag:"-loglevel"`
	Report     bool           `json:"report" flag:"-report"`
	HideBanner bool           `json:"hide_banner" flag:"-hide_banner"`
}
