//go:build windows

package transcoderx

import (
	"github.com/elizabevil/ffmpegx/paramx"
	"path/filepath"
)

func NewConfig(handles ...func(*Config)) Config {
	co := Config{}
	for _, handle := range handles {
		handle(&co)
	}
	return co
}

func NewConfigWithDir(dir string) (Config, error) {
	config := Config{
		filepath.Join(dir, "ffmpeg.exe"),
		filepath.Join(dir, "ffprobe.exe"),
		filepath.Join(dir, "ffplay.exe"),
	}
	if verifyBin(config.FFprobeBin) != nil || verifyBin(config.FFprobeBin) != nil {
		return config, paramx.ErrNotFountBin
	}
	return config, nil
}
