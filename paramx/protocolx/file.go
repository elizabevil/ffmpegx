package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 24.10 file
type File struct {
	Truncate typex.UBool `json:"truncate" flag:"-truncate,1"`
	//Truncate existing files on write, if set to 1. A value of 0 prevents truncating. Default value is 1.

	Blocksize typex.Size `json:"blocksize" flag:"-blocksize"`
	//Set I/O operation maximum block size, in bytes. Default value is INT_MAX, which results in not limiting the requested block size. Setting this value reasonably low improves user termination request reaction time, which is valuable for files on slow medium.

	Follow typex.UBool `json:"follow" flag:"-follow"`
	//If set to 1, the protocol will retry reading at the end of the file, allowing reading files that still are being written. In order for this to terminate, you either need to use the rw_timeout option, or use the interrupt callback (for API users).

	Seekable typex.Bool `json:"seekable" flag:"-seekable"`
	//Controls if seekability is advertised on the file. 0 means non-seekable, -1 means auto (seekable for normal files, non-seekable for named pipes).

	//Many demuxers handle seekable and non-seekable resources differently, overriding this might speed up opening certain files at the cost of losing some features (e.g. accurate seeking).
}
