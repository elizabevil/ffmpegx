package transcoderx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
)

// Metadata File `s Metadata With Default args
func Metadata(ffprobeBin string, input string, handles ...OutByteHandle) (metadatax.Metadata, error) {
	var metadata metadatax.Metadata
	err := verifyBin(ffprobeBin)
	if err != nil {
		return metadata, err
	}
	if len(input) == 0 {
		return metadata, paramx.ErrNotInputs
	}
	args := []string{"-hide_banner", "-i", input, "-print_format", "json", "-show_format", "-show_streams", "-show_error"}
	cmd := exec.Command(ffprobeBin, args...)
	DebugPrint("%s", cmd.String())
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("executing (%s)| error: %s | message: %s %s", cmd.String(), err, outb.String(), errb.String())
	}
	outData := outb.Bytes()
	for _, handle := range handles {
		handle(outData, errb.Bytes())
	}
	if err := json.Unmarshal(outData, &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// MetadataWithArgs File `s Metadata With  args
func MetadataWithArgs(ffprobeBin string, args interfacex.IArg, unmarshal interfacex.Unmarshal, handles ...OutByteHandle) (metadatax.Metadata, error) {
	var metadata metadatax.Metadata
	err := verify(ffprobeBin, args)
	if err != nil {
		return metadata, err
	}
	err = verifyUnmarshal(unmarshal)
	if err != nil {
		return metadata, err
	}
	cmd := exec.Command(ffprobeBin, args.Args()...)
	DebugPrint("%s", cmd.String())
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return metadata, fmt.Errorf("executing (%s)| error: %s | message: %s %s", cmd.String(), err, outb.String(), errb.String())
	}
	outData := outb.Bytes()
	for _, handle := range handles {
		handle(outData, errb.Bytes())
	}
	if err := unmarshal.Unmarshal(outData, &metadata); err != nil {
		return metadata, err
	}
	return metadata, nil
}

// CommandLine Make command
func CommandLine(ffmpegBin string, args interfacex.IArg) (string, string) {
	if verifyArgs(args) != nil {
		return "", ""
	}
	command := exec.Command(ffmpegBin, args.Args()...)
	defer func() { command = nil }()
	DebugPrint("%s", command.String())
	return command.Path, command.String()
}

// Cmd Use command
func Cmd(bin string, args interfacex.IArg, handles ...func(command *exec.Cmd)) (*exec.Cmd, error) {
	err := verify(bin, args)
	if err != nil {
		return nil, err
	}
	command := exec.Command(bin, args.Args()...)
	for _, handle := range handles {
		handle(command)
	}
	DebugPrint("%s", command.String())
	return command, nil
}

// StartProcess Use os.StartProcess
func StartProcess(ffmpegBin string, args interfacex.IArg, handles ...func(command *os.ProcAttr)) (*os.Process, error) {
	err := verify(ffmpegBin, args)
	if err != nil {
		return nil, err
	}
	pa := os.ProcAttr{}
	for _, handle := range handles {
		handle(&pa)
	}
	bin, _ := CommandLine(ffmpegBin, args)
	process, err := os.StartProcess(bin, args.Args(), &pa)
	if err != nil {
		return nil, err
	}
	return process, nil
}

// Pipeline bin:ffmpeg/ffplay Overload Realtime: Chan ProgressHandle
func Pipeline(bin string, ph interfacex.ProgressHandle, args interfacex.IArg, handles ...CmdHandle) (metadatax.ProcessCtxHandle, error) {
	err := verify(bin, args)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(bin, args.Args()...)
	for _, handle := range handles {
		handle(cmd)
	}
	DebugPrint("%s", cmd.String())
	return func(ctx context.Context, handle func(process *os.Process), progressHandle ...metadatax.ProgressHandle) error {
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("pipe %w", err)
		}
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("start %w", err)
		}
		if handle != nil {
			handle(cmd.Process)
		}
		if ctx == nil {
			ctx = context.Background()
		}
		cancelCtx, cancelFunc := context.WithCancel(ctx)
		out := make(chan metadatax.Progress)
		go func() {
			if ph == nil {
				ph = metadatax.DefaultProgress{}
			}
			ph.MakeProgress(cancelCtx, stderrIn, out)
		}()
		go func() {
			err = cmd.Wait()
			cancelFunc()
		}()
		hasFunc := progressHandle != nil
		for {
			select {
			case <-cancelCtx.Done():
				return err
			case <-ctx.Done():
				cmd.Process.Kill()
				return ctx.Err()
			case data, ok := <-out:
				if ok && hasFunc {
					for _, item := range progressHandle {
						item(data)
					}
				}
			}
		}
	}, nil
}

// PipelineCtx bin:ffmpeg/ffplay Overload Realtime: Chan ProgressHandle
func PipelineCtx(ctx context.Context, bin string, ph interfacex.ProgressHandle, args interfacex.IArg, handles ...CmdHandle) (metadatax.ProcessHandle, error) {
	err := verify(bin, args)
	if err != nil {
		return nil, err
	}
	cmd := exec.CommandContext(ctx, bin, args.Args()...)
	for _, handle := range handles {
		handle(cmd)
	}
	DebugPrint("%s", cmd.String())
	return func(handle func(process *os.Process), progressHandle ...metadatax.ProgressHandle) error {
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("pipe %w", err)
		}
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("start %w", err)
		}
		if handle != nil {
			handle(cmd.Process)
		}
		cancelCtx, cancelFunc := context.WithCancel(ctx)
		out := make(chan metadatax.Progress)
		go func() {
			if ph == nil {
				ph = metadatax.DefaultProgress{}
			}
			ph.MakeProgress(cancelCtx, stderrIn, out)
		}()
		go func() {
			err = cmd.Wait()
			cancelFunc()
		}()
		hasFunc := progressHandle != nil
		for {
			select {
			case <-cancelCtx.Done():
				return err
			case data, ok := <-out:
				if ok && hasFunc {
					for _, item := range progressHandle {
						item(data)
					}
				}
			}
		}
	}, nil
}

// PipelinePlay bin:ffplay Overload Realtime: Chan ProgressHandle
func PipelinePlay(bin string, ph interfacex.ProgressHandle, args interfacex.IArg, handles ...CmdHandle) (metadatax.FFplayCtxHandle, error) {
	err := verify(bin, args)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(bin, args.Args()...)
	for _, handle := range handles {
		handle(cmd)
	}
	DebugPrint("%s", cmd.String())
	return func(ctx context.Context, handle func(process *os.Process), progressHandle ...metadatax.PlayHandle) error {
		stderrIn, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("pipe %w", err)
		}
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("start %w", err)
		}
		if handle != nil {
			handle(cmd.Process)
		}
		if ctx == nil {
			ctx = context.Background()
		}
		cancelCtx, cancelFunc := context.WithCancel(ctx)
		out := make(chan metadatax.FFplay)
		go func() {
			if ph == nil {
				ph = metadatax.DefaultProgress{}
			}
			ph.MakePlayProgress(cancelCtx, stderrIn, out)

		}()
		go func() {
			err = cmd.Wait()
			cancelFunc()
		}()
		hasFunc := progressHandle != nil
		for {
			select {
			case <-cancelCtx.Done():
				return err
			case data, ok := <-out:
				if ok && hasFunc {
					for _, item := range progressHandle {
						item(data)
					}
				}
			}
		}
	}, nil
}
