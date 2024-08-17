package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 24.11 ftp
// FTP (File Transfer Protocol).
// ftp://[user[:password]@]server[:port]/path/to/remote/resource.mpeg
type FTP struct {
	Timeout typex.MicrosecondI `json:"timeout" flag:"-timeout,-1"`
	//Set timeout in microseconds of socket I/O operations used by the underlying low level operation. By default it is set to -1, which means that the timeout is not specified.

	FtpUser typex.Username `json:"ftp-user" flag:"-ftp-user"`
	//Set a user to be used for authenticating to the FTP server. This is overridden by the user in the FTP URL.

	FtpPassword typex.Password `json:"ftp-password" flag:"-ftp-password"`
	//Set a password to be used for authenticating to the FTP server. This is overridden by the password in the FTP URL, or by ftp-anonymous-password if no user is set.

	FtpAnonymousPassword typex.Password `json:"ftp-anonymous-password" flag:"-ftp-anonymous-password"`
	//Password used when login as anonymous user. Typically an e-mail address should be used.

	FtpWriteSeekable typex.UBool `json:"ftp-write-seekable" flag:"-ftp-write-seekable"`
	//Control seekability of connection during encoding. If set to 1 the resource is supposed to be seekable, if set to 0 it is assumed not to be seekable. Default value is 0.
}
