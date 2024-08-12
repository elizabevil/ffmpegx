package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// FLAC 21.37 flac
type FLAC struct {
	WriteHeader typex.UBool `json:"write_header" flag:"-write_header"` //write the file header if set to true, default is true
}

//This muxer accepts exactly one FLAC audio stream. Additionally, it is possible to add images with disposition ‘attached_pic’.
//ffmpeg -i INPUT -i pic1.png -i pic2.jpg -map 0:a -map 1 -map 2 -disposition:v attached_pic OUTPUT
