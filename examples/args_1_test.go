package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"testing"
)

func TestParamEmpty(t *testing.T) {
	param := paramx.Param{}
	fmt.Println(param.Args()) //EMPTY
}
func TestStreamFilter(t *testing.T) {
	stream := optionx.InputStream{
		Itsscale:     0,
		ReinitFilter: 12,
	}
	fmt.Println(parsex.DefaultParser.ParamParse(stream))
	common := optionx.Common{
		ArStreamSpecifier: typex.StreamSpecifier{
			K: "as",
			V: "as",
		},
	}
	fmt.Println(parsex.DefaultParser.ParamParse(common))
	outS := optionx.OutputStream{
		Disposition: "attached_pic",
	}
	fmt.Println(parsex.DefaultParser.ParamParse(outS))
}
func TestPointArgs(t *testing.T) {
	hls := muxerx.HLS{
		HlsInitTime:        1,
		HlsTime:            10,
		HlsDeleteThreshold: 2,
		HlsListSize:        &typex.ZeroUN, //default 5 --> force 0
		HlsSegmentFilename: "%Y%m%d%H%M%S.ts",
	}
	fmt.Println(hls.Args())
	fmt.Println(parsex.DefaultParser.ParamParse(hls))
}
func createCutTimeParam() paramx.Param {
	param := paramx.Param{}
	param.GlobalHandle(func(input *optionx.Global) {
		input.Overwrite = true
	}).InputHandle(func(input *optionx.Input) {
		input.Inputs = []string{InputVedio}
	}).OutputHandle(func(output *optionx.Output) {
		output.Outputs = []string{OutputVideo}
	}).CommonHandle(paramx.PositionInput, optionx.Common{
		An: true,
	}).CommonHandle(paramx.PositionOutput, optionx.Common{
		F: "hls",
		CodecSpecifier: []typex.StreamSpecifier{
			{"v", "h264x"},
		},
	}).CommonHandle(paramx.PositionOutput, optionx.Common{ //overwrite
		F: "flv",
		T: 4,
	})
	return param
}
func TestCutTimeArgs(t *testing.T) {
	param := createCutTimeParam()
	fmt.Println(param.Args())
}
