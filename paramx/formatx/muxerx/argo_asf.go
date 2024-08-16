package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 21.15 argo_asf ArgoAsf
type ArgoAsf struct {
	VersionMajor typex.UI16 `json:"version_major" flag:"-version_major"`
	//override      file major version, specified as an integer, default value is 2

	VersionMinor typex.UI16 `json:"version_minor" flag:"-version_minor"`
	//override file minor version, specified as an integer, default value is 1

	Name typex.Name `json:"name" flag:"-name"`
	//Embed file name into file, if not specified use the output file name.The name is truncated to 8 characters.
}
