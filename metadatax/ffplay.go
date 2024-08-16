package metadatax

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type PlayHandle = func(progress FFplay)
type FFplayCtxHandle = func(ctx context.Context, handle func(process *os.Process), handles ...PlayHandle) error
type FFplay struct {
	MasterClock float32 `json:"master_clock"`
	Key         string  `json:"key"`
	Diff        float32 `json:"diff"`
	Fd          int     `json:"fd"`
	Aq          int     `json:"aq"` //KB
	Vq          int     `json:"vq"` //kb
	Sq          int     `json:"sq"` //B
	F           string  `json:"f"`  //B
}

var playRegexp = regexp.MustCompile(`\\s+|:`)

func makePlayProgress(line string, ff *FFplay) {
	allString := playRegexp.ReplaceAllString(line, " ")
	_, err := fmt.Sscanf(
		allString,
		"%f %s %f fd= %d aq= %dKB vq= %dKB sq= %dB f=%s",
		&ff.MasterClock, &ff.Key, &ff.Diff, &ff.Fd, &ff.Aq, &ff.Vq, &ff.Sq, &ff.F)
	if err != nil {
		return
	}
}

func (r DefaultProgress) MakePlayProgress(ctx context.Context, stream io.ReadCloser, out chan FFplay) {
	if ctx == nil || ctx.Err() != nil {
		return
	}
	r.ffplayNoChan(ctx, r.makeScanner(stream), out)
}
func (r DefaultProgress) ffplayNoChan(ctx context.Context, scanner *bufio.Scanner, out chan FFplay) {
	pp := FFplay{}
	next := false
	for scanner.Scan() {
		line := scanner.Text()
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "fd=") && strings.Contains(line, "aq=")
		}
		if next {
			select {
			case <-ctx.Done():
				return
			default:
				makePlayProgress(line, &pp)
				out <- pp
			}
		}
	}
}
