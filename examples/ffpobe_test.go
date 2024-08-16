package examples

import (
	"encoding/xml"
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"strconv"
	"strings"
	"testing"
)

// /usr/bin/ffprobe -hide_banner -i ./file/input.mp4 -print_format json -show_format -show_streams -show_error
func TestProbe(t *testing.T) {
	transcoderx.Debug = true
	transcode := transcoderx.NewTranscoder(transcoderx.NewConfig(func(config *transcoderx.Config) {
		config.FFprobeBin = FfprobeBin
	}))
	metadata, err := transcode.Metadata(InputVedio, func(out []byte, err []byte) {

	})
	if err != nil {
		panic(err)
	}
	fmt.Println(metadata.Format)
	for _, stream := range metadata.Streams {
		fmt.Println(stream.Index, stream.CodecType, stream.CodecName, stream.CodecLongName, stream.StreamVideoOnly.Width)
	}
}

func TestProbeParams(t *testing.T) {
	generic := optionx.Generic{HideBanner: true}
	fprobe := optionx.FFprobe{
		I:                InputVedio,
		PrintFormat:      optionx.OutputFormat_json,
		ShowData:         false, //packets_and_frames data
		ShowError:        true,
		ShowFormat:       true,
		ShowFrames:       false,
		ShowPackets:      false,
		ShowPrograms:     true,
		ShowStreamGroups: false,
		ShowStreams:      true,
		ShowChapters:     true,
		ShowVersions:     false, //all version
		CountFrames:      true,
		CountPackets:     true,
		//ShowDataHash:        flagx.Hash_MD5,
		//ShowEntries: "stream=codec_type:stream_disposition=default", //    stream>>disposition>>default
		ShowLog: 0,
	}
	transcoderx.Debug = true
	metadata, err := transcoderx.MetadataWithArgs(FfprobeBin, paramx.BuildIArgInterface(generic, fprobe), transcoderx.JsonUnmarshal, func(out []byte, err []byte) {
	})
	//metadata, err := transcoderx.MetadataWithArgs(FfprobeBin, paramx.BuildIArgInterface(generic, fprobe), XmlUnmarshal{})
	if err != nil {
		fmt.Println("MetadataX", err)
		return
	}
	fmt.Println(metadata.Format)
	fmt.Println(metadata.StreamWithType(typex.AVMEDIA_TYPE_VIDEO))
	fmt.Println(metadata.StreamCountWithType(typex.AVMEDIA_TYPE_VIDEO))
}

type XmlUnmarshal struct {
}

func (x XmlUnmarshal) Unmarshal(bytes []byte, data any) error {
	return xml.Unmarshal(bytes, data)
}

func TestProgressSpeed(t *testing.T) {
	float, err := strconv.ParseFloat(strings.TrimRight("5.73e+03x", "x"), 10)
	fmt.Println(err, float)
}
