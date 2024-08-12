package flagx

type Hash = Flag

const (
	Hash_MD5        Hash = "MD5"        //        16
	Hash_MURMUR3    Hash = "murmur3"    //    16
	Hash_RIPEMD128  Hash = "RIPEMD128"  //  16
	Hash_RIPEMD160  Hash = "RIPEMD160"  //  20
	Hash_RIPEMD256  Hash = "RIPEMD256"  //  32
	Hash_RIPEMD320  Hash = "RIPEMD320"  //  40
	Hash_SHA160     Hash = "SHA160"     //     20
	Hash_SHA224     Hash = "SHA224"     //     28
	Hash_SHA256     Hash = "SHA256"     //     32
	Hash_SHA512_224 Hash = "SHA512/224" // 28
	Hash_SHA512_256 Hash = "SHA512/256" // 32
	Hash_SHA384     Hash = "SHA384"     //     48
	Hash_SHA512     Hash = "SHA512"     //     64
	Hash_CRC32      Hash = "CRC32"      //       4
	Hash_ADLER32    Hash = "adler32"    //     4
)
