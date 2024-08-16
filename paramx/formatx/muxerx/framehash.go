package muxerx

// 21.40 framehash FRAMEHASH
type FRAMEHASH struct {
	Hash string `json:"hash" flag:"-hash"`
	//Use the cryptographic hash function specified by the string algorithm. Supported values include MD5, murmur3, RIPEMD128, RIPEMD160, RIPEMD256, RIPEMD320, SHA160, SHA224, SHA256 (default), SHA512/224, SHA512/256, SHA384, SHA512, CRC32 and adler32.
}
