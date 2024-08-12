package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"log"
	"os"
	"testing"
)

func init() {
	transcoderx.Debug = true
}

var pwd, _ = os.Getwd()

func makeTranscoder() transcoderx.Transcoder {
	transcode := transcoderx.NewTranscoder(transcoderx.NewConfig())
	return transcode
}

func TestTranscoderProcess(t *testing.T) {
	tran := makeTranscoder()
	tran.Args = makeParams()
	transcoderx.Debug = true
	process, err := tran.StartProcess(func(command *os.ProcAttr) {
	})
	if err != nil {
		fmt.Println("StartProcess", err)
		return
	}
	wait, err := process.Wait()
	if err != nil {
		fmt.Println("Wait", err)
		return
	}
	fmt.Println(wait.Exited(), err)
}

// ============

func TestCmd(t *testing.T) {
	cmd, _ := transcoderx.Cmd(FfmpegBin, createCutTimeParam())
	cmd.Stderr = os.Stdout
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Panicln(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(err)
}
func TestTranscoderCmd(t *testing.T) {
	tran := makeTranscoder()
	var cmd = func(tran transcoderx.Transcoder) {
		cmd, _ := tran.Cmd()
		cmd.Stderr = os.Stdout
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			fmt.Errorf("%w cmd", err)
			return
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Errorf("%w Wait", err)
			return
		}
		fmt.Println(err)
	}
	tran.Args = createCutTimeParam()
	cmd(tran)
}
