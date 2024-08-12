package optionx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

func (r Global) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Generic) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r IO) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Output) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Input) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Audio) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Expert) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r ExpertInput) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r ExpertOutput) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r AdvPerFile) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Encoding) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Decoding) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Coding) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Format) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r FormatInput) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r FormatOutput) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Resampler) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Protocol) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}

func (r Scaler) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r Subtitle) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r FFprobe) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r WriterOption) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r FFplay) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
func (r FFplayExpert) Args() typex.Args {
	return parsex.DefaultParser.ParamParse(r)
}
