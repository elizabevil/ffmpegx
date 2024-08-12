package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBRAV1E 9.6 librav1e
type LIBRAV1E struct {
	Qmax typex.Quantizer `json:"qmax" flag:"-qmax"`
	//Sets the maximum quantizer to use when using bitrate mode.

	Qmin typex.Quantizer `json:"qmin" flag:"-qmin"`
	//Sets the minimum quantizer to use when using bitrate mode.

	Qp typex.Level `json:"qp" flag:"-qp"`
	//Uses quantizer mode to encode at the given quantizer (0-255).

	Speed typex.Speed `json:"speed" flag:"-speed"`
	//Selects the speed preset (0-10) to encode with.

	Tiles typex.TileCount `json:"tiles" flag:"-tiles"`
	//Selects how many tiles to encode with.

	TileRows typex.TileCount `json:"tile-rows" flag:"-tile-rows"`
	//Selects how many rows of tiles to encode with.

	TileColumns typex.TileCount `json:"tile-columns" flag:"-tile-columns"`
	//Selects how many columns of tiles to encode with.

	Rav1EParams string `json:"rav1e-params" flag:"-rav1e-params"`
	//Set rav1e options using a list of key=value pairs separated by ":". See rav1e --help for a list of options.

}

/*

For example to specify librav1e encoding options with -rav1e-params:

ffmpeg -i input -c:v librav1e -b:v 500K -rav1e-params speed=5:low_latency=true output.mp4

*/
