package demuxerx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	sas := []string{
		"20.1 aa",
		"20.2 aac",
		"20.3 apng",
		"20.4 asf",
		"20.5 concat",
		"20.5.1 Syntax",
		"20.6 dash",
		"20.7 dvdvideo",
		"20.8 ea",
		"20.9 imf",
		"20.10 flv, live_flv, kux",
		"20.11 gif",
		"20.12 hls",
		"20.13 image2",
		"20.14 libgme",
		"20.15 libmodplug",
		"20.16 libopenmpt",
		"20.17 mov/mp4/3gp",
		"20.18 mpegts",
		"20.19 mpjpeg",
		"20.20 rawvideo",
		"20.21 rcwt",
		"20.22 sbg",
		"20.23 tedcaptions",
		"20.24 vapoursynth",
		"20.25 w64",
		"20.26 wav"}
	for _, sa := range sas {
		split := strings.Split(sa, " ")
		s2 := strings.ToUpper(split[1])
		s := split[1] + ".go"
		_, err := os.Stat(s)
		if err == nil && !errors.Is(err, os.ErrNotExist) {
			continue
		}

		os.WriteFile(s, []byte(fmt.Sprintf("package demuxerx\n\t// %s %s\ntype %s struct { \n }", sa, s2, s2)), 0644)
	}
}
