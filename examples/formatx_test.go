package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/demuxerx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	mpegts := demuxerx.Mpegts{
		ScanAllPmts:    &typex.True,
		FixTeletextPts: &typex.False,
		TsPacketsize:   1,
		MaxPacketSize:  1,
	}
	fmt.Println("mpegts", mpegts.Args())
}
func TestHls(t *testing.T) {
	hls := muxerx.HLS{
		HlsFlags:           flagx.HlsFlags_delete_segments,
		HlsTime:            typex.TimeDurationSecondI(10001 * time.Millisecond),
		HlsDeleteThreshold: 1,
		Strftime:           typex.True,
		HlsListSize:        &typex.ZeroUN,
		HlsSegmentFilename: "HlsSegmentFilename",
	}
	fmt.Println(hls.Args())
	//====
	hls2 := demuxerx.HLS{
		HttpSeekable: &typex.False,
	}
	fmt.Println(hls2.Args())
	//====

}
func TestSegment(t *testing.T) {
	generic := optionx.Generic{
		HideBanner: true,
		Loglevel:   flagx.Loglevel_warning,
	}
	expert := optionx.Expert{
		NoStdin: false,
	}
	formatInput := optionx.FormatInput{
		ErrDetect:       flagx.ErrDetect_ignore_err,
		Probesize:       512000,
		Analyzeduration: 512000,
	}
	format := optionx.Format{
		Fflags: flagx.Fflags_genpts,
	}
	global := optionx.Global{
		Overwrite: true,
	}
	input := optionx.Input{
		Input: "$from",
	}

	resampler := optionx.Resampler{
		Async: 1,
	}
	segment := muxerx.Segment{
		SegmentTime:          5,
		SegmentFormat:        "mpegts",
		SegmentFormatOptions: "mpegts_flags=initial_discontinuity:mpegts_copyts=1",
		SegmentListType:      muxerx.SegmentListType_m3u8,
		SegmentListFlags:     "+" + muxerx.SegmentListFlags_live,
		SegmentWrap:          5, //loop
		//SegmentListSize:      6,
	}
	common := optionx.Common{
		//Codec: "copy",
		Vcodec: "copy",
		Acodec: "copy",
		Scodec: "copy",
		//F: "segment",
		//F: "hls",
	}
	line, _ := transcoderx.CommandLine("ffmpeg", paramx.BuildIArgInterface(global, expert, generic, format, resampler, formatInput, input,
		common, segment, typex.Args{"-segment_list", "out.m3u8", "%d.ts"},
	))
	fmt.Println(line)
}
