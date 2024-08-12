package examples

import (
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
	"testing"
	"time"
)

// testing
func TestPlayParams(t *testing.T) {
	generic := optionx.Generic{HideBanner: true}
	ffplay := optionx.FFplay{
		X:            0,
		Y:            0,
		Fs:           true,
		An:           false,
		Vn:           false,
		Sn:           false,
		Ss:           0,
		T:            0,
		Bytes:        0,
		SeekInterval: 0,
		Nodisp:       false,
		Noborder:     true,
		Alwaysontop:  true,
		Volume:       99,
		F:            "",
		WindowTitle:  "Text",
		Left:         0,
		Top:          0,
		Loop:         0,
		Vf:           "",
		Af:           "",
		Showmode:     "",
		I:            InputVedio,
	}
	ffplaEy := optionx.FFplayExpert{
		Codec:           "",
		Stats:           false,
		Fast:            false,
		Sync:            "",
		Autorotate:      true,
		NoAutorotate:    false,
		Genpts:          false,
		Lowres:          0,
		Autoexit:        true,
		Exitonkeydown:   false,
		Exitonmousedown: false,
		Framedrop:       true,
		NoFramedrop:     false,
		Infbuf:          true,
		NoInfbuf:        false,
		EnableVulkan:    false,
	}
	argInterface := paramx.BuildIArgInterface(generic, ffplay, ffplaEy)
	fmt.Println(argInterface.Args())
	cmd, err := transcoderx.Cmd(FfplayBin, argInterface)
	if err != nil {
		return
	}
	fmt.Println(cmd.String())
	cmd.Run()

}
func makePlayParam() interfacex.IArgs {
	generic := optionx.Generic{HideBanner: true}
	ffplay := optionx.FFplay{
		X:            0,
		Y:            0,
		Fs:           true,
		An:           false,
		Vn:           false,
		Sn:           false,
		Ss:           0,
		T:            0,
		Bytes:        0,
		SeekInterval: 0,
		Nodisp:       false,
		Noborder:     false,
		Alwaysontop:  false,
		Volume:       99,
		F:            "",
		WindowTitle:  "Text",
		Left:         0,
		Top:          0,
		Loop:         0,
		Vf:           "",
		Af:           "",
		Showmode:     "",
		I:            InputVedio,
	}
	return paramx.BuildIArgInterface(generic, ffplay)
}
func TestPlayPipeLine(t *testing.T) {
	transcoderx.Debug = true
	duration := 3000 * time.Millisecond
	killDuration := duration * 100
	background, cc := context.WithTimeout(context.Background(), duration)
	defer cc()
	//=========
	fmt.Println("Start", time.Now().Format(time.DateTime))
	now := time.Now()
	defer func() {
		fmt.Println("Since", time.Since(now))
	}()
	runFunc, err := transcoderx.PipelinePlay(FfplayBin, transcoderx.NewProgressMaker(), makePlayParam(), func(cmd *exec.Cmd) {
	})
	if err != nil {
		return
	}
	err = runFunc(background, func(process *os.Process) {
		go func() {
			time.Sleep(killDuration)
			process.Kill()
		}()
	}, func(progress metadatax.FFplay) {
		fmt.Println("progress::", progress)
	})
	fmt.Println("transcoderx.Pipeline", err)
}
