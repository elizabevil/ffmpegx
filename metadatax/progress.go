package metadatax

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"regexp"
	"strings"
)

// Progress  fftools\ffmpeg.c
type Progress struct {
	Frame string `json:"frame"`
	Fps   string `json:"fps"`
	Q     string `json:"q"`
	Size  string `json:"size"`
	Time  string `json:"time"`

	Bitrate   string `json:"bitrate"`
	TotalSize string `json:"total_size"`

	Speed      string `json:"speed"`
	OutTimeUs  string `json:"out_time_us"`
	OutTimeMs  string `json:"out_time_ms"`
	OutTime    string `json:"out_time"`
	Dup        string `json:"dup"`
	DropFrames string `json:"drop_frames"`

	Progress string `json:"progress"`
}

type DefaultProgress struct {
	Filter func(str string) bool
}

var re = regexp.MustCompile(`=\s+`)

func makeProgress(line string, pp *Progress) {
	st := re.ReplaceAllString(line, `=`)
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

func (r DefaultProgress) makeScanner(stream io.ReadCloser) *bufio.Scanner {
	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner := bufio.NewScanner(stream)
	scanner.Split(split)
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	return scanner
}

func (r DefaultProgress) MakeProgress(ctx context.Context, stream io.ReadCloser, out chan Progress) {
	if ctx == nil || ctx.Err() != nil {
		return
	}
	go r.progressHandle(ctx, r.makeScanner(stream), out)
	select {
	case <-ctx.Done():
		close(out)
		return
	}
}
func (r DefaultProgress) progressHandle(ctx context.Context, scanner *bufio.Scanner, out chan Progress) {
	pp := Progress{}
	next := false
	for scanner.Scan() {
		line := scanner.Text()
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "time=") && strings.Contains(line, "bitrate=") && strings.Contains(line, "speed=")
		}
		if next && ctx.Err() == nil {
			makeProgress(line, &pp)
			out <- pp
		}
	}
}

func (r DefaultProgress) MakeProgressX(ctx context.Context, stream io.ReadCloser, handle ProgressHandle) {
	if ctx == nil || ctx.Err() != nil {
		return
	}
	r.progressNoChan(ctx, r.makeScanner(stream), handle)
}
func (r DefaultProgress) progressNoChan(ctx context.Context, scanner *bufio.Scanner, handle ProgressHandle) {
	pp := Progress{}
	for scanner.Scan() {
		line := scanner.Text()
		next := false
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "time=") && strings.Contains(line, "size=") && strings.Contains(line, "bitrate=") && strings.Contains(line, "speed=")
		}
		if next {
			select {
			case <-ctx.Done():
				return
			default:
				makeProgress(line, &pp)
				handle(pp)
			}
		}
	}
}
