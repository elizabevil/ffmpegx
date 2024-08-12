package transcoderx

import (
	"github.com/elizabevil/ffmpegx/metadatax"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"os"
	"os/exec"
)

type Transcoder struct {
	Args   interfacex.IArg
	Ph     interfacex.ProgressHandle
	Config Config `json:"config"`
}

func NewTranscoder(config Config) Transcoder {
	config.MustVerify()
	return Transcoder{Config: config, Ph: NewProgressMaker()}
}

func (r Transcoder) Metadata(input string, handles ...OutByteHandle) (metadatax.Metadata, error) {
	return Metadata(r.Config.FFprobeBin, input, handles...)
}

func (r Transcoder) MetadataWithArgs(args interfacex.IArg, unmarshal interfacex.Unmarshal, handles ...OutByteHandle) (metadatax.Metadata, error) {
	return MetadataWithArgs(r.Config.FFprobeBin, args, unmarshal, handles...)
}

func (r Transcoder) CommandLine(args interfacex.IArg) (string, string) {
	return CommandLine(r.Config.FFmpegBin, args)
}

func (r Transcoder) Cmd(handles ...CmdHandle) (*exec.Cmd, error) {
	return Cmd(r.Config.FFmpegBin, r.Args, handles...)
}
func (r Transcoder) StartProcess(handles ...func(command *os.ProcAttr)) (*os.Process, error) {
	return StartProcess(r.Config.FFmpegBin, r.Args, handles...)
}

func (r Transcoder) Pipeline(cmdHandles ...CmdHandle) (metadatax.ProcessCtxHandle, error) {
	return Pipeline(r.Config.FFmpegBin, r.Ph, r.Args, cmdHandles...)
}

func (r Transcoder) PipelinePlay(cmdHandles ...CmdHandle) (metadatax.FFplayCtxHandle, error) {
	return PipelinePlay(r.Config.FFplayBin, r.Ph, r.Args, cmdHandles...)
}
