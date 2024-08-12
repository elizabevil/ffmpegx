package muxerx

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	sas := []string{
		"21.1 Raw muxers",
		"21.2 Raw PCM muxers",
		"21.3 MPEG-1/MPEG-2 program stream muxers",
		"21.4 MOV/MPEG-4/ISOMBFF muxers",
		"21.4.1 Fragmentation",
		"21.5 a64",
		"21.6 ac4",
		"21.7 adts",
		"21.8 aea",
		"21.9 aiff",
		"21.10 alp",
		"21.11 amr",
		"21.12 amv",
		"21.13 apm",
		"21.14 apng",
		"21.15 argo_asf",
		"21.16 argo_cvg",
		"21.17 asf, asf_stream",
		"21.18 ass",
		"21.19 ast",
		"21.20 au",
		"21.21 avi",
		"21.22 avif",
		"21.23 avm2",
		"21.24 bit",
		"21.25 caf",
		"21.26 codec2",
		"21.27 chromaprint",
		"21.28 crc",
		"21.29 dash",
		"21.30 daud",
		"21.31 dv",
		"21.32 ffmetadata",
		"21.33 fifo",
		"21.34 film_cpk",
		"21.35 filmstrip",
		"21.36 fits",
		"21.37 flac",
		"21.38 flv",
		"21.39 framecrc",
		"21.40 framehash",
		"21.41 framemd5",
		"21.42 gif",
		"21.43 gxf",
		"21.44 hash",
		"21.45 hds",
		"21.46 hls",
		"21.47 iamf",
		"21.48 ico",
		"21.49 ilbc",
		"21.50 image2, image2pipe",
		"21.51 ircam",
		"21.52 ivf",
		"21.53 jacosub",
		"21.54 kvag",
		"21.55 lc3",
		"21.56 lrc",
		"21.56.1 Metadata",
		"21.57 matroska",
		"21.57.1 Metadata",
		"21.58 md5",
		"21.59 microdvd",
		"21.60 mmf",
		"21.61 mp3",
		"21.62 mpegts",
		"21.63 mxf, mxf_d10, mxf_opatom",
		"21.64 null",
		"21.65 nut",
		"21.66 ogg",
		"21.67 rcwt",
		"21.68 segment, stream_segment, ssegment",
		"21.69 smoothstreaming",
		"21.70 streamhash",
		"21.71 tee",
		"21.72 webm_chunk",
		"21.73 webm_dash_manifest",
	}
	for _, sa := range sas {
		split := strings.Split(sa, " ")
		s2 := strings.ToUpper(split[1])
		s := split[1] + ".go"
		_, err := os.Stat(s)
		if err == nil && !errors.Is(err, os.ErrNotExist) {
			continue
		}

		os.WriteFile(s, []byte(fmt.Sprintf("package muxerx\n\t// %s %s\ntype %s struct { \n }", sa, s2, s2)), 0644)
	}
}
