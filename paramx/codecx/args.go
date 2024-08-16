package codecx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

func (r Audio) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r AV) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r DE) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Decoding) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r DEV) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r EAV) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r EV) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Subtitle) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
