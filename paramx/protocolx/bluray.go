package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

type Bluray struct {
	Angle typex.UI8 `json:"angle" flag:"-angle"`
	//BluRay angle

	Chapter typex.UI16 `json:"chapter" flag:"-chapter"`
	//Start chapter (1...N)

	Playlist typex.UI16 `json:"playlist" flag:"-playlist"`
	//Playlist to read (BDMV/PLAYLIST/?????.mpls)
}

/*
Read longest playlist from BluRay mounted to /mnt/bluray:

bluray:/mnt/bluray
Read angle 2 of playlist 4 from BluRay mounted to /mnt/bluray, start from chapter 2:

-playlist 4 -angle 2 -chapter 2 bluray:/mnt/bluray

*/

// 24.4 cache
//Cache the input stream to temporary file. It brings seeking capability to live streams.

type Cache struct {
	ReadAheadLimit typex.Limit `json:"read_ahead_limit" flag:"-read_ahead_limit"`
	//Amount in bytes that may be read ahead when seeking isn’t supported. Range is -1 to INT_MAX. -1 for unlimited. Default is 65536.
}

// 24.5 concat
type Concat struct {
}
