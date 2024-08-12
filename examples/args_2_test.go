package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const FfmpegBin = "/usr/bin/ffmpeg"
const FfprobeBin = "/usr/bin/ffprobe"
const FfplayBin = "/usr/bin/ffplay"
const InputVedio = "./file/input.mp4"
const OutputVideo = "./file/output.mp4"
const OutputVideoHls = "./file/output.m3u8"

func hlsParams() interfacex.IArgs {
	input := optionx.Input{Input: InputVedio}
	output := optionx.Output{Output: OutputVideoHls}
	args := paramx.AnyArgs{"-c": "copy"}
	//argsx := typex.Args{"-c", "copy"}
	//var as typex.Size = 2
	hls := muxerx.HLS{
		HlsFlags: flagx.HlsFlags_delete_segments,
		//HlsListSize: &typex.ZeroUN,
		//HlsListSize: &as,
		//HlsTime:            typex.TimeDurationSecondI(61 * time.Second),
		//HlsTime:            typex.TimeDurationParseSecondI("00:01:01"),
		HlsTime:            typex.TimeDurationParseSecondI("0m2s"),
		Strftime:           typex.True,
		HlsDeleteThreshold: 1,
		HlsSegmentFilename: "%Y%m%d%H%M%S.ts",
	}
	argInterface := paramx.BuildIArgInterface(input, args, hls, output)
	return argInterface
}
func TestHlsParams(t *testing.T) {
	fmt.Println(hlsParams().Args())
	fmt.Println(makeParams().Args())
}
func makeParams() paramx.Param {
	param := paramx.Param{}
	param.GlobalHandle(func(input *optionx.Global) {
		input.Overwrite = true
	}).InputHandle(func(input *optionx.Input) {
		input.Re = true
		input.Inputs = []string{filepath.Join(pwd, InputVedio)}
	}).OutputHandle(func(output *optionx.Output) {
		output.Outputs = []string{filepath.Join(pwd, OutputVideo)}
		output.Vcodec = "libx264"
	}).CommonHandle(paramx.PositionOutput, optionx.Common{Acodec: "copy"})
	return param
}
func makeParamsToHls() interfacex.IArgs {
	getwd, _ := os.Getwd()
	argInterface := paramx.BuildIArgInterface(
		optionx.Expert{NoStdin: true},
		optionx.Generic{HideBanner: true},
		optionx.Global{Overwrite: true, StatsPeriod: 0.5},
		optionx.Input{
			Re:    true,
			Input: filepath.Join(getwd, InputVedio),
		},
		optionx.Common{
			Acodec: "copy",
			Vcodec: "copy",
			Scodec: "copy",
		},
		muxerx.HLS{
			HlsTime:            typex.TimeDurationSecondI(1000 * time.Millisecond),
			HlsDeleteThreshold: 1,
			Strftime:           1,
			HlsSegmentFilename: "%s.ts",
		},
		optionx.Output{Output: filepath.Base(OutputVideoHls)})
	return argInterface
}
