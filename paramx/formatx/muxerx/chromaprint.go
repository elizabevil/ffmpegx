package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.27 chromaprint CHROMAPRINT
type CHROMAPRINT struct {
	Algorithm *typex.UBool `json:"algorithm" flag:"-algorithm,1"`
	//Select version of algorithm to fingerprint with. Range is 0 to 4. Version 3 enables silence detection. Default is 1.

	FpFormat FpFormat `json:"fp_format" flag:"-fp_format"`
	//Format to output the fingerprint as. Accepts the following options:

	SilenceThreshold typex.I16 `json:"silence_threshold" flag:"-silence_threshold"`
	//Threshold for detecting silence. Range is from -1 to 32767, where -1 disables silence detection. Silence detection can only be used with version 3 of the algorithm.

	//Silence detection must be disabled for use with the AcoustID service. Default is -1.

}
type FpFormat = typex.String

const (
	FpFormat_base64 FpFormat = "base64"
	//Base64 compressed fingerprint (default)

	FpFormat_compressed FpFormat = "compressed"
	//Binary compressed fingerprint

	FpFormat_raw FpFormat = "raw"
	//Binary raw fingerprint
)
