package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // JPEG2000 9.5 jpeg2000
type JPEG2000 struct {
	Format *typex.UBool `json:"format" flag:"-format"`
	//Can be set to either j2k CODEC_J2K 0 or jp2 CODEC_JP2 1 (the default) that makes it possible to store non-rgb pix_fmts.

	TileWidth typex.Size `json:"tile_width" flag:"-tile_width"`
	//Sets tile width. Range is 1 to 1073741824.1<<30 Default is 256.

	TileHeight typex.Size `json:"tile_height" flag:"-tile_height"`
	//Sets tile height. Range is 1 to 1073741824. Default is 256.
	Pred     typex.UBool `json:"pred" flag:"-pred"`
	PredFlag Pred        `json:"pred_flag" flag:"-pred"`

	Sop typex.UBool `json:"sop" flag:"-sop"`
	//Enable this to add SOP marker at the start of each packet. Disabled by default.

	Eph typex.UBool `json:"eph" flag:"-eph"`
	//Enable this to add EPH marker at the end of each packet header. Disabled by default.

	Prog       Prog   `json:"prog" flag:"-prog"`
	LayerRates string `json:"layer_rates" flag:"-layer_rates"`
	//By default, when this option is not used, compression is done using the quality metric. This option allows for compression using compression ratio. The compression ratio for each level could be specified. The compression ratio of a layer l species the what ratio of total file size is contained in the first l layers.

}
type Pred string

const (
	dwt97int Pred = "dwt97int"
	dwt53    Pred = "dwt53"
)

type Prog string

const (
	lrcp Prog = "lrcp"
	rlcp Prog = "rlcp"
	rpcl Prog = "rpcl"
	pcrl Prog = "pcrl"
	cprl Prog = "cprl"
)

/*

Example usage:

ffmpeg -i input.bmp -c:v jpeg2000 -layer_rates "100,10,1" output.j2k
This would compress the image to contain 3 layers, where the data contained in the first layer would be compressed by 1000 times, compressed by 100 in the first two layers, and shall contain all data while using all 3 layers.
*/
