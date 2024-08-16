package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// 24.16 Icecast
// icecast://[username[:password]@]server:port/mountpoint
type Icecast struct {
	IceGenre string `json:"ice_genre" flag:"-ice_genre"`
	//Set the stream genre.

	IceName string `json:"ice_name" flag:"-ice_name"`
	//Set the stream name.

	IceDescription string `json:"ice_description" flag:"-ice_description"`
	//Set the stream description.

	IceUrl string `json:"ice_url" flag:"-ice_url"`
	//Set the stream website URL.

	IcePublic typex.UBool `json:"ice_public" flag:"-ice_public"`
	//Set if the stream should be public.The default is 0 (not public).

	UserAgent string `json:"user_agent" flag:"-user_agent"`
	//Override the User-Agent header.If not specified a string of the form "Lavf/<version>" will be used.

	Password string `json:"password" flag:"-password"`
	//Set the Icecast mountpoint password.

	ContentType string `json:"content_type" flag:"-content_type"`
	//Set the stream content type.This must be set if it is different from audio/mpeg.

	LegacyIcecast typex.UBool `json:"legacy_icecast" flag:"-legacy_icecast"`
	//This enables support for Icecast versions < 2.4.0, that do not support the HTTP PUT method but the SOURCE method.

	Tls typex.UBool `json:"tls" flag:"-tls"`
	//Establish a TLS (HTTPS) connection to Icecast.
}
