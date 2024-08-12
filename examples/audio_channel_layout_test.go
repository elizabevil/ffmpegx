package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"testing"
)

func TestChannelLayout(t *testing.T) {
	layoutName := metadatax.ChannelLayoutMap["stereo"]
	fmt.Println(layoutName)
	fmt.Println(metadatax.ChannelNames[layoutName.Layout.Mask])
}

func TestChannel(t *testing.T) {
	stream := metadatax.Stream{
		BaseStream: metadatax.BaseStream{},
		StreamAudioOnly: metadatax.StreamAudioOnly{
			ChannelLayout: "stereo",
		},
		StreamVideoOnly: metadatax.StreamVideoOnly{},
	}
	layout := stream.GetChannelLayout()
	fmt.Println(layout)
	fmt.Println(metadatax.ChannelNames[layout.Layout.Mask])
}
