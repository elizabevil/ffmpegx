package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.71 tee TEE
type TEE struct {

	//Muxer options can be specified for each slave by prepending them as a list of key=value pairs separated by ’:’, between square brackets. If the options values contain a special character or the ’:’ separator, they must be escaped; note that this is a second level escaping.

	//The following special options are also recognized:

	F typex.Format `json:"f" flag:"-f"`
	//Specify the format name. Required if it cannot be guessed from the output URL.

	Bsfs typex.Filter `json:"bsfs" flag:"-bsfs"`
	//Specify a list of bitstream filters to apply to the specified output.

	//It is possible to specify to which streams a given bitstream filter applies, by appending a stream specifier to the option separated by /. spec must be a stream specifier (see Format stream specifiers).

	//If the stream specifier is not specified, the bitstream filters will be applied to all streams in the output. This will cause that output operation to fail if the output contains streams to which the bitstream filter cannot be applied e.g. h264_mp4toannexb being applied to an output containing an audio stream.

	//Options for a bitstream filter must be specified in the form of opt=value.

	//Several bitstream filters can be specified, separated by ",".

	UseFifo typex.UBool `json:"use_fifo" flag:"-use_fifo"`
	//This allows to override tee muxer use_fifo option for individual slave muxer.

	FifoOptions typex.Dict `json:"fifo_options" flag:"-fifo_options"`
	//This allows to override tee muxer fifo_options for individual slave muxer. See fifo.

	Select typex.Param `json:"select" flag:"-select"`
	//Select the streams that should be mapped to the slave output, specified by a stream specifier. If not specified, this defaults to all the mapped streams. This will cause that output operation to fail if the output format does not accept all mapped streams.

	//You may use multiple stream specifiers separated by commas (,) e.g.: a:0,v

	Onfail Onfail `json:"onfail" flag:"-onfail"`
	//Specify behaviour on output failure. This can be set to either abort (which is default) or ignore. abort will cause whole process to fail in case of failure on this slave output. ignore will ignore failure on this output, so other outputs will continue without being affected.
}

type Onfail = typex.String

const (
	Onfail_abort   Onfail = "abort"
	Onfail_ignore  Onfail = "ignore"
	Onfail_default Onfail = "default"
)
