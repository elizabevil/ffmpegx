package demuxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

func (r Apng) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Asf) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r DVD) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r GIF) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Mpegts) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r MPEG4) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r HLS) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

//0 --> -1
//0 ->0
