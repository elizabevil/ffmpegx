package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Global struct {
	RwTimeout typex.MicrosecondI `json:"rw_timeout" flag:"-rw_timeout"`
}
