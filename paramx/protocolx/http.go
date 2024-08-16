package protocolx

import (
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
)

// 24.15 http
// HTTP (Hyper Text Transfer Protocol).
type HTTP struct {
	Seekable typex.UBool `json:"seekable" flag:"-seekable,-1"`
	//Control seekability of connection. If set to 1 the resource is supposed to be seekable, if set to 0 it is assumed not to be seekable, if set to -1 it will try to autodetect if it is seekable. Default value is -1.

	ChunkedPost typex.UBool `json:"chunked_post" flag:"-chunked_post"`
	//If set to 1 use chunked Transfer-Encoding for posts, default is 1.

	HttpProxy string `json:"http_proxy" flag:"-http_proxy"`
	//set HTTP proxy to tunnel through e.g. http://example.com:1234

	Headers string `json:"headers" flag:"-headers"`
	//Set custom HTTP headers, can override built in default headers. The value must be a string encoding the headers.

	ContentType string `json:"content_type" flag:"-content_type"`
	//Set a specific content type for the POST messages or for listen mode.

	UserAgent string `json:"user_agent" flag:"-user_agent"`
	//Override the User-Agent header. If not specified the protocol will use a string describing the libavformat build. ("Lavf/<version>")

	Referer string `json:"referer" flag:"-referer"`
	//Set the Referer header. Include ’Referer: URL’ header in HTTP request.

	MultipleRequests typex.UBool `json:"multiple_requests" flag:"-multiple_requests"`
	//Use persistent connections if set to 1, default is 0.

	PostData string `json:"post_data" flag:"-post_data"`
	//Set custom HTTP post data.

	MimeType string `json:"mime_type" flag:"-mime_type"`
	//Export the MIME type.

	HttpVersion bool `json:"http_version" flag:"-http_version"`
	//Exports the HTTP response version number. Usually "1.0" or "1.1".

	Cookies string `json:"cookies" flag:"-cookies"`
	//Set the cookies to be sent in future requests. The format of each cookie is the same as the value of a Set-Cookie HTTP response field. Multiple cookies can be delimited by a newline character.

	Icy typex.UBool `json:"icy" flag:"-icy,1"`
	//If set to 1 request ICY (SHOUTcast) metadata from the server. If the server supports this, the metadata has to be retrieved by the application by reading the icy_metadata_headers and icy_metadata_packet options. The default is 1.

	IcyMetadataHeaders string `json:"icy_metadata_headers" flag:"-icy_metadata_headers"`
	//If the server supports ICY metadata, this contains the ICY-specific HTTP reply headers, separated by newline characters.

	IcyMetadataPacket string `json:"icy_metadata_packet" flag:"-icy_metadata_packet"`
	//If the server supports ICY metadata, and icy was set to 1, this contains the last non-empty metadata packet sent by the server. It should be polled in regular intervals by applications interested in mid-stream metadata updates.

	Metadata string `json:"metadata" flag:"-metadata"`
	//Set an exported dictionary containing Icecast metadata from the bitstream, if present. Only useful with the C API.

	AuthType flagx.AuthType `json:"auth_type" flag:"-auth_type"`
	//Set HTTP authentication type. No option for Digest, since this method requires getting nonce parameters from the server first and can’t be used straight away like Basic.

	//Basic authentication sends a Base64-encoded string that contains a user name and password for the client. Base64 is not a form of encryption and should be considered the same as sending the user name and password in clear text (Base64 is a reversible encoding). If a resource needs to be protected, strongly consider using an authentication scheme other than basic authentication. HTTPS/TLS should be used with basic authentication. Without these additional security enhancements, basic authentication should not be used to protect sensitive or valuable information.

	SendExpect100 *typex.UBool `json:"send_expect_100" flag:"-send_expect_100,-1"`
	//Send an Expect: 100-continue header for POST. If set to 1 it will send, if set to 0 it won’t, if set to -1 it will try to send if it is applicable. Default value is -1.

	Location string `json:"location" flag:"-location"`
	//An exported dictionary containing the content location. Only useful with the C API.

	Offset typex.Offset `json:"offset" flag:"-offset"`
	//Set initial byte offset.

	EndOffset typex.U16Offset `json:"end_offset" flag:"-end_offset"`
	//Try to limit the request to bytes preceding this offset.

	Method flagx.HttpMethod `json:"method" flag:"-method"`
	//When used as a client option it sets the HTTP method for the request.

	//When used as a server option it sets the HTTP method that is going to be expected from the client(s). If the expected and the received HTTP method do not match the client will be given a Bad Request response. When unset the HTTP method is not checked for now. This will be replaced by autodetection in the future.

	Reconnect typex.UBool `json:"reconnect" flag:"-reconnect"`
	//Reconnect automatically when disconnected before EOF is hit.

	ReconnectAtEof typex.UBool `json:"reconnect_at_eof" flag:"-reconnect_at_eof"`
	//If set then eof is treated like an error and causes reconnection, this is useful for live / endless streams.

	ReconnectOnNetworkError typex.UBool `json:"reconnect_on_network_error" flag:"-reconnect_on_network_error"`
	//Reconnect automatically in case of TCP/TLS errors during connect.

	ReconnectOnHttpError string `json:"reconnect_on_http_error" flag:"-reconnect_on_http_error"`
	//A comma separated list of HTTP status codes to reconnect on. The list can include specific status codes (e.g. ’503’) or the strings ’4xx’ / ’5xx’.

	ReconnectStreamed typex.UBool `json:"reconnect_streamed" flag:"-reconnect_streamed"`
	//If set then even streamed/non seekable streams will be reconnected on errors.

	ReconnectDelayMax bool `json:"reconnect_delay_max" flag:"-reconnect_delay_max"`
	//Set the maximum delay in seconds after which to give up reconnecting.

	ReconnectMaxRetries bool `json:"reconnect_max_retries" flag:"-reconnect_max_retries"`
	//Set the maximum number of times to retry a connection. Default unset.

	ReconnectDelayTotalMax bool `json:"reconnect_delay_total_max" flag:"-reconnect_delay_total_max"`
	//Set the maximum total delay in seconds after which to give up reconnecting.

	RespectRetryAfter bool `json:"respect_retry_after" flag:"-respect_retry_after"`
	//If enabled, and a Retry-After header is encountered, its requested reconnection delay will be honored, rather than using exponential backoff. Useful for 429 and 503 errors. Default enabled.

	Listen typex.UBool `json:"listen" flag:"-listen"`
	//If set to 1 enables experimental HTTP server. This can be used to send data when used as an output option, or read data from a client with HTTP POST when used as an input option. If set to 2 enables experimental multi-client HTTP server. This is not yet implemented in ffmpeg.c and thus must not be used as a command line option.

	Resource bool `json:"resource" flag:"-resource"`
	//The resource requested by a client, when the experimental HTTP server is in use.

	ReplyCode typex.Code `json:"reply_code" flag:"-reply_code"`
	//The HTTP code returned to the client, when the experimental HTTP server is in use.

	ShortSeekSize typex.Size `json:"short_seek_size" flag:"-short_seek_size"`
	//Set the threshold, in bytes, for when a readahead should be prefered over a seek and new HTTP request. This is useful, for example, to make sure the same connection is used for reading large video packets with small audio packets in between.
}

/*

# Server side (sending):
ffmpeg -i somefile.ogg -c copy -listen 1 -f ogg http://server:port

# Client side (receiving):
ffmpeg -i http://server:port -c copy somefile.ogg

# Client can also be done with wget:
wget http://server:port -O somefile.ogg

# Server side (receiving):
ffmpeg -listen 1 -i http://server:port -c copy somefile.ogg

# Client side (sending):
ffmpeg -i somefile.ogg -chunked_post 0 -c copy -f ogg http://server:port

# Client can also be done with wget:
wget --post-file=somefile.ogg http://server:port

*/
