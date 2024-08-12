package inputdevicex

import "github.com/elizabevil/ffmpegx/paramx/typex"

// LIBCDIO 3.13 libcdio
type LIBCDIO struct {
	Speed typex.Speed `json:"speed" flag:"-speed"`
	//Set drive reading speed. Default value is 0.

	//The speed is specified CD-ROM speed units. The speed is set through the libcdio cdio_cddap_speed_set function. On many CD-ROM drives, specifying a value too large will result in using the fastest speed.

	ParanoiaMode ParanoiaMode `json:"paranoia_mode" flag:"-paranoia_mode"`
	//Set paranoia recovery mode flags. It accepts one of the following values:

	//Default value is ‘disable’.
	//For more information about the available recovery modes, consult the paranoia project documentation.

}
type ParanoiaMode string

const (
	ParanoiaMode_disable   ParanoiaMode = "disable"
	ParanoiaMode_verify    ParanoiaMode = "verify"
	ParanoiaMode_overlap   ParanoiaMode = "overlap"
	ParanoiaMode_neverskip ParanoiaMode = "neverskip"
	ParanoiaMode_full      ParanoiaMode = "full"
)
