package metadatax

import (
	"regexp"
	"strings"
)

// Progress  fftools\ffmpeg.c
type Progress struct {
	Frame string `json:"frame" toml:"frame"`
	Fps   string `json:"fps" toml:"fps"`
	Q     string `json:"q" toml:"q"`
	Size  string `json:"size" toml:"size"`
	Time  string `json:"time" toml:"time"`

	Bitrate   string `json:"bitrate" toml:"bitrate"`
	TotalSize string `json:"total_size" toml:"total_size"`

	Speed      string `json:"speed" toml:"speed"`
	OutTimeUs  string `json:"out_time_us" toml:"out_time_us"`
	OutTimeMs  string `json:"out_time_ms" toml:"out_time_ms"`
	OutTime    string `json:"out_time" toml:"out_time"`
	Dup        string `json:"dup" toml:"dup"`
	DropFrames string `json:"drop_frames" toml:"drop_frames"`

	Progress string `json:"progress" toml:"progress"`
}

func (p *Progress) Reset() {
	p.Frame = emptysString.Value()
	p.Fps = emptysString.Value()
	p.Q = emptysString.Value()
	p.Size = emptysString.Value()
	p.Time = emptysString.Value()

	p.Bitrate = emptysString.Value()
	p.TotalSize = emptysString.Value()

	p.Speed = emptysString.Value()
	p.OutTimeUs = emptysString.Value()
	p.OutTimeMs = emptysString.Value()
	p.OutTime = emptysString.Value()
	p.Dup = emptysString.Value()
	p.DropFrames = emptysString.Value()

	p.Progress = emptysString.Value()
}

var defaultFilterFunc = func(line string) bool {
	return strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") && strings.Contains(line, "speed=")
}
var progressRegexp = regexp.MustCompile(`=\s+`)

func makeProgress(line string, pp *Progress) {
	st := progressRegexp.ReplaceAllString(line, `=`)
	f := strings.Fields(st)
	for j := 0; j < len(f); j++ {
		field := f[j]
		fieldSplit := strings.Split(field, "=")
		if len(fieldSplit) > 1 {
			fieldname := strings.Split(field, "=")[0]
			fieldvalue := strings.Split(field, "=")[1]
			switch fieldname {
			case "frame":
				pp.Frame = fieldvalue
			case "fps":
				pp.Fps = fieldvalue
			case "q":
				pp.Q = fieldvalue
			case "size":
				pp.Size = fieldvalue
			case "time":
				pp.Time = fieldvalue
			case "bitrate":
				pp.Bitrate = fieldvalue
			case "total_size":
				pp.TotalSize = fieldvalue
			case "speed":
				pp.Speed = fieldvalue
			case "out_time_us":
				pp.OutTimeUs = fieldvalue
			case "out_time_ms":
				pp.OutTimeMs = fieldvalue
			case "out_time":
				pp.OutTime = fieldvalue
			case "dup":
				pp.Dup = fieldvalue
			case "drop_frames":
				pp.DropFrames = fieldvalue
			case "progress":
				pp.Progress = fieldvalue
			}
		}
	}
}
