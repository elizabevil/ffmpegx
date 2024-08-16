package interfacex

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"reflect"
)

// IArg  impl IArg --> your custom method
type IArg interface {
	Args() typex.Args
}

type IArgs []IArg

func (a IArgs) Args() typex.Args {
	t := make(typex.Args, 0, len(a))
	for _, item := range a {
		if item == nil {
			continue
		}
		t = append(t, item.Args()...)
	}
	return t
}

type ParseArgs interface {
	ParamParse(input any) typex.Args
	ParamItemType(of reflect.Value) (string, bool)
}
