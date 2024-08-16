package transcoderx

import (
	"github.com/elizabevil/ffmpegx/paramx"
	"os"
)

type Config struct {
	FFmpegBin  string `json:"ffmpeg_bin"`
	FFprobeBin string `json:"ffprobe_bin"`
	FFplayBin  string `json:"ffplay_bin"`
}

func (c Config) Verify() error {
	stat, err := os.Stat(c.FFmpegBin)
	if err != nil || stat.IsDir() {
		return paramx.ErrNotFountBin
	}
	stat, err = os.Stat(c.FFprobeBin)
	if err != nil || stat.IsDir() {
		return paramx.ErrNotFountBin
	}
	return nil
}
func (c Config) MustVerify() {
	mustVerifyBin(c.FFmpegBin)
	mustVerifyBin(c.FFprobeBin)
}
