package paramx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"reflect"
	"slices"
)

// ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...

func BuildArgs(args ...typex.Args) typex.Args {
	//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...
	return slices.Concat(args...)
}
func BuildIArgs(args ...interfacex.IArg) typex.Args {
	return interfacex.IArgs(slices.Concat(args)).Args()
}
func BuildIArgInterface(args ...interfacex.IArg) interfacex.IArgs {
	return slices.Concat(args)
}

type AnyArgs map[string]any

func (r AnyArgs) Args() typex.Args {
	args := make(typex.Args, 0, len(r))
	for key, value := range r {
		itemType, b := parsex.DefaultParser.ParamItemType(reflect.ValueOf(value))
		if b {
			if len(key) > 0 {
				args = append(args, key, itemType)
			} else {
				// outputs
				args = append(args, itemType)
			}
		} else {
			args = append(args, key+itemType)
		}
	}
	return args
}
