package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/protocolx"
	"reflect"
	"testing"
)

// options contains a list of &-separated options of the form key=val.
func TestProtocolx(t *testing.T) {
	rtp := protocolx.RTP{
		Ttl:           8080,
		RtcpPort:      8080,
		LocalRtpPort:  8080,
		LocalRtcpPort: 8080,
		Sources:       "",
	}
	fmt.Println(rtp.Options())
	url, err := rtp.Url("localhost:80")
	if err != nil {
		return
	}
	fmt.Println(url.String())
	asd := map[string]string{"as": "as"}
	fmt.Println(parsex.DefaultParser.ParamItemType(reflect.ValueOf(asd)))
}
