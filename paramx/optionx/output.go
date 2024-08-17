package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// https://ffmpeg.org/ffmpeg-all.html
type Output struct {
	Output  string   `json:"output" flag:""`
	Outputs []string `json:"outputs" flag:""`
	//5.4 Main options
	Fs           typex.LimitSize         `json:"fs" flag:"-fs"`
	Timestamp    typex.Date              `json:"timestamp" flag:"-timestamp"`
	Metadata     map[string]string       `json:"metadata" flag:"-metadata"`
	Disposition  []string                `json:"disposition" flag:"-disposition"`   //-disposition v
	DispositionX []typex.StreamSpecifier `json:"disposition_x" flag:"-disposition"` // -disposition:v:1
	Program      string                  `json:"program" flag:"-program"`           //[title=title:][program_num=program_num:]st=stream[:st=stream...]
	StreamGroup  string                  `json:"stream_group" flag:"-stream_group"` //[map=input_file_id=stream_group][type=type:]st=stream[:st=stream][:stg=stream_group][:id=stream_group_id...]
	Target       typex.OutputTarget      `json:"target" flag:"-target"`
	Dframes      typex.Frames            `json:"dframes" flag:"-dframes"` //This is an obsolete alias for -frames:d, which you should use instead.
	Attach       typex.FileName          `json:"attach" flag:"-attach"`
	//5.5 Video Options
	Vframes typex.Frames   `json:"vframes" flag:"-vframes"` //This is an obsolete alias for -frames:v
	Vcodec  typex.Codec    `json:"vcodec" flag:"-vcodec"`   //codec:This is an alias for codec:v.
	Pass    string         `json:"pass" flag:"-pass"`       //Select the pass number (1 or 2)
	Vf      map[string]any `json:"vf" flag:"-vf"`           //This is an alias for filter:v,
	Vtag    typex.Tag      `json:"vtag " flag:"-vtag "`     //Force video tag/fourcc. This is an alias for -tag:v.
	//5.7 Audio Options
	Aframes typex.Frames      `json:"aframes" flag:"-aframes"`
	Aq      string            `json:"aq" flag:"-aq"` //This is an alias for filter:a, see the filter option.
	Af      typex.FilterGraph `json:"af" flag:"-af"`
	//5.8 Advanced Audio options
	Atag typex.Tag `json:"atag" flag:"-atag"` //Force audio tag/fourcc. This is an alias for -tag:a.
	Map  []string  `json:"map" flag:"-map"`

	//5.11 Advanced options
	MapChapters         typex.InputFileIndex `json:"map_chapters" flag:"-map_chapters"` //input_file_index
	Shortest            bool                 `json:"shortest" flag:"-shortest"`
	ShortestBufDuration typex.Duration       `json:"shortest_buf_duration" flag:"-shortest_buf_duration"`
	Muxdelay            typex.SecondF        `json:"muxdelay" flag:"-muxdelay"`     //seconds
	Muxpreload          typex.SecondF        `json:"muxpreload" flag:"-muxpreload"` //seconds
	Streamid            []string             `json:"streamid" flag:"-streamid"`     //outputstreamindex:newvalue
	//19 Format Options
	Packetsize         typex.Size         `json:"packetsize" flag:"-packetsize"`
	AudioPreload       typex.MicrosecondI `json:"audio_preload" flag:"-audio_preload"`
	ChunkDuration      typex.MicrosecondI `json:"chunk_duration" flag:"-chunk_duration"`
	ChunkSize          typex.Size         `json:"chunk_size" flag:"-chunk_size"`
	MaxInterleaveDelta typex.MicrosecondI `json:"max_interleave_delta" flag:"-max_interleave_delta"`
	//Set maximum buffering duration for interleaving. The duration is expressed in microseconds, and defaults to 10000000 (10 seconds).
	AvoidNegativeTs AvoidNegativeTs `json:"avoid_negative_ts" flag:"-avoid_negative_ts"`
	FlushPackets    *typex.Bool     `json:"flush_packets" flag:"-flush_packets"`
	OutputTsOffset  typex.Offset    `json:"output_ts_offset" flag:"-output_ts_offset"`
	//==========
	Filter typex.String `json:"filter" flag:"-filter"` //typex.String, OPT_PERSTREAM | OPT_OUTPUT | OPT_HAS_ALT,
	////==========
	//=======
	Aspect typex.String `json:"aspect" flag:"-aspect"` //typex.String, OPT_VIDEO | OPT_PERSTREAM | OPT_OUTPUT,
	B      typex.String `json:"b" flag:"-b"`           //typex.String,   OPT_VIDEO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT,
	//=======
	Ab typex.String `json:"ab" flag:"-ab"` //typex.String, OPT_AUDIO | OPT_FUNC_ARG | OPT_PERFILE | OPT_OUTPUT,
	//=======
}
type AvoidNegativeTs = typex.String

const (
	AvoidNegativeTs_make_non_negative AvoidNegativeTs = "make_non_negative" //
	//Shift timestamps to make them non-negative. Also note that this affects only leading negative timestamps, and not non-monotonic negative timestamps.
	AvoidNegativeTs_make_zero AvoidNegativeTs = "make_zero" //
	//Shift timestamps so that the first timestamp is 0.
	AvoidNegativeTs_auto AvoidNegativeTs = "auto" // (default)
	//Enables shifting when required by the target format.
	AvoidNegativeTs_disabled AvoidNegativeTs = "disabled" //
	//Disables shifting of timestamp.
)
