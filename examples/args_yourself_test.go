package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"reflect"
	"testing"
)

type YourHls muxerx.HLS

func (r YourHls) Args() typex.Args {
	return YourParse{}.ParamParse(r)
}

type YourParse struct {
}

func (y YourParse) ParamParse(input any) typex.Args {
	return []string{"-c", "copy"}
}

func (y YourParse) ParamItemType(of reflect.Value) (string, bool) {
	return "c", false
}

func TestArgsYourself(t *testing.T) {
	hls := YourHls{}
	fmt.Println(hls.Args())
}
