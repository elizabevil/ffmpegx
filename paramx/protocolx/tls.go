package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/parsex"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"net/url"
)

// 24.42 tls
// Transport Layer Security (TLS) / Secure Sockets Layer (SSL)
// tls://hostname:port[?options]
type TLS struct {
	CaFile bool `json:"cafile" flag:"-cafile"` // filename
	//A file containing certificate authority (CA) root certificates to treat as trusted.If the linked TLS library contains a default this might not need to be specified for verification to work, but not all libraries and setups have defaults built in.The file must be in OpenSSL PEM format.

	TlsVerify typex.Bool `json:"tls_verify" flag:"-tls_verify"` // 1|0
	//If enabled, try to verify the peer that we are communicating with.Note, if using OpenSSL, this currently only makes sure that the peer certificate is signed by one of the root certificates in the CA database, but it does not validate that the certificate actually matches the host name we are trying to connect to.(With other backends, the host name is validated as well.)

	//This is disabled by default since it requires a CA database to be provided by the caller in many cases.

	CertFile string `json:"cert" flag:"-cert"` // filename
	//A file containing a certificate to use in the handshake with the peer.(When operating as server, in listen mode, this is more often required by the peer, while client certificates only are mandated in certain setups.)

	KeyFile string `json:"key" flag:"-key"` // filename
	//A file containing the private key for the certificate.

	Listen bool `json:"listen" flag:"-listen"` // 1|0
	//If enabled, listen for connections on the provided port, and assume the server role in the handshake instead of the client role.

	HttpProxy bool `json:"http_proxy" flag:"-http_proxy"`
	//The HTTP proxy to tunnel through, e.g.http: //example.com:1234. The proxy must support the CONNECT method.

}

/*
To create a TLS/SSL server that serves an input stream.
	ffmpeg -i input -f format tls://hostname:port?listen&cert=server.crt&key=server.key
To play back a stream from the TLS/SSL server using ffplay:
	ffplay tls://hostname:port
*/

func (r TLS) Options() (string, error) {
	return parsex.Options(r)
}
func (r TLS) Url(destination string) (url.URL, error) {
	//sap://destination[:port][?options]
	//ffmpeg -re -i input -f sap sap://224.0.0.255?same_port=1
	options, err := r.Options()
	if err != nil {
		return url.URL{}, err
	}
	u := url.URL{
		Scheme:   "sap",
		Host:     destination,
		RawQuery: options,
	}
	return u, nil
}
