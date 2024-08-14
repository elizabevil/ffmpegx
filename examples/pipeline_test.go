package examples

import (
	"context"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"os"
	"os/exec"
	"testing"
	"time"
)

// ============

func TestPipeline(t *testing.T) {
	transcoderx.Debug = true
	duration := 1000 * time.Millisecond
	killDuration := duration / 100
	background, cc := context.WithTimeout(context.Background(), duration)
	defer cc()
	//=========
	fmt.Println("Start", time.Now().Format(time.DateTime))
	now := time.Now()
	defer func() {
		fmt.Println("Since", time.Since(now))
	}()
	runFunc, err := transcoderx.Pipeline(FfmpegBin, nil, makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(background, func(process *os.Process) {
		go func() {
			time.Sleep(killDuration)
			cc()
		}()
	}, func(progress metadatax.Progress) {
		fmt.Println("progress::", progress)
	})
	fmt.Println("transcoderx.Pipeline", err)
}
func TestPipelineCtx(t *testing.T) {
	transcoderx.Debug = true
	duration := 1000 * time.Millisecond
	killDuration := duration / 100
	background, cc := context.WithTimeout(context.Background(), duration)
	defer cc()
	//=========
	fmt.Println("Start", time.Now().Format(time.DateTime))
	now := time.Now()
	defer func() {
		fmt.Println("Since", time.Since(now))
	}()
	runFunc, err := transcoderx.PipelineCtx(background, FfmpegBin, nil, makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(func(process *os.Process) {
		go func() {
			time.Sleep(killDuration)
			process.Kill()
		}()
	}, func(progress metadatax.Progress) {
		fmt.Println("progress::", progress)
	})
	fmt.Println("transcoderx.Pipeline", err)
}

func TestPipelineX(t *testing.T) {
	transcoderx.Debug = true
	duration := 10000 * time.Millisecond
	killDuration := duration * 100
	timeout, cancelFunc := context.WithTimeout(context.Background(), duration) // run time dur
	defer cancelFunc()
	count := 0
	ctx, c := context.WithCancel(timeout)
	go func() {
		for count < 2 {
			fmt.Println("========---------=====")
			pipelineXChild(ctx, killDuration)
			count++
		}
		c()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx Exit")
	}
}
func pipelineXChild(ctx context.Context, kill time.Duration) {
	fmt.Println("Start", time.Now().Format(time.DateTime))
	defer fmt.Println("C pipelineXChild end")
	runFunc, err := transcoderx.Pipeline(FfmpegBin, transcoderx.NewProgressMaker(), makeParamsToHls(), func(cmd *exec.Cmd) {
		cmd.Dir = "/tmp"
	})
	if err != nil {
		return
	}
	err = runFunc(ctx, func(process *os.Process) {
		go func() {
			time.Sleep(kill)
			process.Kill()
		}()
	}, func(progress metadatax.Progress) {
		fmt.Println("progress::", progress)
	})
	if err != nil {
		fmt.Println("transcoderx.Pipeline", err)
		return
	}
}
