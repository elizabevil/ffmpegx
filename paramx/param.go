package paramx

import (
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"slices"
)

//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...

type Position bool

const (
	PositionInput  Position = false
	PositionOutput Position = true
)

type InputHandel = func(opt *optionx.Input)
type OutPutHandel = func(opt *optionx.Output)
type ParamHandel = func(opt *optionx.IO)
type GlobalHandel = func(opt *optionx.Global)

type Param struct {
	Input        optionx.Input   `json:"input"`
	Output       optionx.Output  `json:"output"`
	Global       optionx.Global  `json:"global"`
	CommonInput  interfacex.IArg `json:"common_input"`
	CommonOutput interfacex.IArg `json:"common_output"`
}

func (r *Param) InputHandle(handles ...InputHandel) *Param {
	for _, handle := range handles {
		handle(&r.Input)
	}
	return r
}
func (r *Param) OutputHandle(handles ...OutPutHandel) *Param {
	for _, handle := range handles {
		handle(&r.Output)
	}
	return r

}
func (r *Param) GlobalHandle(handles ...GlobalHandel) *Param {
	for _, handle := range handles {
		handle(&r.Global)
	}
	return r

}

//========

func (r *Param) SetInputs(inputs ...string) *Param {
	r.Input.Inputs = inputs
	return r
}
func (r *Param) SetOutputs(output ...string) *Param {
	r.Output.Outputs = output
	return r
}
func (r *Param) CommonHandle(isInput Position, args interfacex.IArg) *Param {
	if isInput == PositionInput {
		r.CommonInput = args
	} else {
		r.CommonOutput = args
	}
	return r
}

func (r Param) Args() typex.Args {
	//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...
	return slices.Concat(r.Global.Args(), BuildIArgs(r.CommonInput), r.Input.Args(), BuildIArgs(r.CommonOutput), r.Output.Args())
}
