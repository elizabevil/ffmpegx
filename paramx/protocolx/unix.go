package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

//24.44 unix
//Unix local socket
//unix://filepath

type Unix struct {
	Timeout typex.MillisecondI `json:"timeout" flag:"-timeout"`
	//Timeout in ms.

	Listen typex.UBool `json:"listen" flag:"-listen"`
	//Create the Unix socket in listening mode.
}
