package metadatax

import (
	"context"
	"fmt"
	"os"
	"regexp"
)

type FFPlayHandle = func(progress FFplay)
type FFplayCtxHandle = func(ctx context.Context, handle func(process *os.Process), handles ...FFPlayHandle) error

type FFplay struct {
	MasterClock float32 `json:"master_clock" toml:"master_clock"`
	Key         string  `json:"key" toml:"key"`
	Diff        float32 `json:"diff" toml:"diff"`
	Fd          int     `json:"fd" toml:"fd"`
	Aq          int     `json:"aq" toml:"aq"` //KB
	Vq          int     `json:"vq" toml:"vq"` //kb
	Sq          int     `json:"sq" toml:"sq"` //B
	F           string  `json:"f" toml:"f"`   //B
}

func (p *FFplay) Reset() {
	p.MasterClock = 0
	p.Key = emptysString.Value()
	p.Diff = 0
	p.Fd = 0
	p.Aq = 0
	p.Vq = 0
	p.Sq = 0
	p.F = emptysString.Value()
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
