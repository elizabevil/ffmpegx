package muxerx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 21.65 nut NUT
type NUT struct {
	Syncpoints flagx.Syncpoints `json:"syncpoints" flag:"-syncpoints"`
	//Change the syncpoint usage in nut:

	//default use the normal low-overhead seeking aids.
	//	none do not use the syncpoints at all, reducing the overhead but making the stream non-seekable
	//	Use of this option is not recommended, as the resulting files are very damage sensitive and seeking is not possible.Also in general the overhead from syncpoints is negligible.Note, -write_index 0 can be used to disable all growing data tables, allowing to mux endless streams with limited memory and without these disadvantages.

	//timestamped extend the syncpoint with a wallclock field.
	//The none and timestamped flags are experimental.

	WriteIndex typex.UBool `json:"write_index" flag:"-write_index"`
	//Write index at the end, the default is to write an index.
}
