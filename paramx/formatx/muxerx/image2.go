package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 21.50 image2, image2pipe IMAGE2,
type IMAGE2 struct {
	FramePts typex.UBool `json:"frame_pts" flag:"-frame_pts"`
	//If set to 1, expand the filename with the packet PTS (presentation time stamp). Default value is 0.

	StartNumber typex.UNumber `json:"start_number" flag:"-start_number"`
	//Start the sequence from the specified number. Default value is 1.

	Update typex.UBool `json:"update" flag:"-update"`
	//If set to 1, the filename will always be interpreted as just a filename, not a pattern, and the corresponding file will be continuously overwritten with new images. Default value is 0.

	Strftime typex.UBool `json:"strftime" flag:"-strftime"`
	//If set to 1, expand the filename with date and time information from strftime(). Default value is 0.

	AtomicWriting typex.UBool `json:"atomic_writing" flag:"-atomic_writing"`
	//Write output to a temporary file, which is renamed to target filename once writing is completed. Default is disabled.

	ProtocolOpts typex.Param `json:"protocol_opts" flag:"-protocol_opts"`
	//Set protocol options as a :-separated list of key=value parameters. Values containing the : special character must be escaped.
}
