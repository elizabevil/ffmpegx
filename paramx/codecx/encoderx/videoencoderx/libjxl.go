package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBJXL 9.9 libjxl
type LIBJXL struct {
	Distance typex.Distance `json:"distance" flag:"-distance"`
	//Set the target Butteraugli distance. This is a quality setting: lower distance yields higher quality, with distance=1.0 roughly comparable to libjpeg Quality 90 for photographic content. Setting distance=0.0 yields true lossless encoding. Valid values range between 0.0 and 15.0, and sane values rarely exceed 5.0. Setting distance=0.1 usually attains transparency for most input. The default is 1.0.

	Effort typex.Level `json:"effort" flag:"-effort"`
	//Set the encoding effort used. Higher effort values produce more consistent quality and usually produces a better quality/bpp curve, at the cost of more CPU time required. Valid values range from 1 to 9, and the default is 7.

	Modular typex.UBool `json:"modular" flag:"-modular"`
	//Force the encoder to use Modular mode instead of choosing automatically. The default is to use VarDCT for lossy encoding and Modular for lossless. VarDCT is generally superior to Modular for lossy encoding but does not support lossless encoding.

}
