package transcoderx

import (
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
)

func verifyBin(bin string) error {
	stat, err := os.Stat(bin)
	if err != nil || stat.IsDir() {
		return paramx.ErrNotFountBin
	}
	return nil
}

func verifyArgs(args interfacex.IArg) error {
	if args == nil {
		return paramx.ErrNotArgs
	}
	return nil
}
func verifyUnmarshal(bin interfacex.Unmarshal) error {
	if bin == nil {
		return paramx.ErrNotUnmarshal
	}
	return nil
}

func verify(bin string, args interfacex.IArg) error {
	err := verifyBin(bin)
	if err != nil {
		return err
	}
	err = verifyArgs(args)
	if err != nil {
		return err
	}
	return nil
}

func mustVerifyBin(bin string) {
	err := verifyBin(bin)
	if err != nil {
		panic(err)
	}
}
