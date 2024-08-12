package protocolx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // 24.24 rtmp
// // Real-Time Messaging Protocol.
// rtmp://[username:password@]server[:port][/app][/instance][/playpath]
type Rtmp struct {
	Username typex.Username `json:"username" flag:"-username"`
	//An optional username (mostly for publishing).

	Password typex.Password `json:"password" flag:"-password"`
	//An optional password (mostly for publishing).

	Server typex.Server `json:"server" flag:"-server"`
	//The address of the RTMP server.

	Port typex.Port `json:"port" flag:"-port"`
	//The number of the TCP port to use (by default is 1935).

	App string `json:"app" flag:"-app"`
	//It is the name of the application to access. It usually corresponds to the path where the application is installed on the RTMP server (e.g. /ondemand/, /flash/live/, etc.). You can override the value parsed from the URI through the rtmp_app option, too.

	Playpath string `json:"playpath" flag:"-playpath"`
	//It is the path or name of the resource to play with reference to the application specified in app, may be prefixed by "mp4:". You can override the value parsed from the URI through the rtmp_playpath option, too.

	Listen bool `json:"listen" flag:"-listen"`
	//Act as a server, listening for an incoming connection.

	Timeout typex.Timeout `json:"timeout" flag:"-timeout"`
	//Maximum time to wait for the incoming connection. Implies listen.
	//Additionally, the following parameters can be set via command line options (or in code via AVOptions):

	RtmpApp string `json:"rtmp_app" flag:"-rtmp_app"`
	//Name of application to connect on the RTMP server. This option overrides the parameter specified in the URI.

	RtmpBuffer typex.Size `json:"rtmp_buffer" flag:"-rtmp_buffer"`
	//Set the client buffer time in milliseconds. The default is 3000.

	RtmpConn string `json:"rtmp_conn" flag:"-rtmp_conn"`
	//.
	RtmpEnhancedCodecs string `json:"rtmp_enhanced_codecs" flag:"-rtmp_enhanced_codecs"`
	//Specify the list of codecs the client advertises to support in an enhanced RTMP stream. This option should be set to a comma separated list of fourcc values, like hvc1,av01,vp09 for multiple codecs or hvc1 for only one codec. The specified list will be presented in the "fourCcLive" property of the Connect Command Message.

	RtmpFlashver string `json:"rtmp_flashver" flag:"-rtmp_flashver"`
	//Version of the Flash plugin used to run the SWF player. The default is LNX 9,0,124,2. (When publishing, the default is FMLE/3.0 (compatible; <libavformat version>).)

	RtmpFlushInterval typex.Size `json:"rtmp_flush_interval" flag:"-rtmp_flush_interval"`
	//Number of packets flushed in the same request (RTMPT only). The default is 10.

	RtmpLive RtmpLive `json:"rtmp_live" flag:"-rtmp_live"`
	//Specify that the media is a live stream. No resuming or seeking in live streams is possible. The default value is any, which means the subscriber first tries to play the live stream specified in the playpath. If a live stream of that name is not found, it plays the recorded stream. The other possible values are live and recorded.

	RtmpPageurl typex.Url `json:"rtmp_pageurl" flag:"-rtmp_pageurl"`
	//URL of the web page in which the media was embedded. By default no value will be sent.

	RtmpPlaypath typex.Url `json:"rtmp_playpath" flag:"-rtmp_playpath"`
	//Stream identifier to play or to publish. This option overrides the parameter specified in the URI.

	RtmpSubscribe string `json:"rtmp_subscribe" flag:"-rtmp_subscribe"`
	//mask of live stream to subscribe to. By default no value will be sent. It is only sent if the option is specified or if rtmp_live is set to live.

	RtmpSwfhash bool `json:"rtmp_swfhash" flag:"-rtmp_swfhash"`
	//SHA256 hash of the decompressed SWF file (32 bytes).

	RtmpSwfsize typex.Size `json:"rtmp_swfsize" flag:"-rtmp_swfsize"`
	//Size of the decompressed SWF file, required for SWFVerification.

	RtmpSwfurl string `json:"rtmp_swfurl" flag:"-rtmp_swfurl"`
	//URL of the SWF player for the media. By default no value will be sent.

	RtmpSwfverify string `json:"rtmp_swfverify" flag:"-rtmp_swfverify"`
	//URL to player swf file, compute hash/size automatically.

	RtmpTcurl typex.Url `json:"rtmp_tcurl" flag:"-rtmp_tcurl"`
	//URL of the target stream. Defaults to proto://host[:port]/app.

	TcpNodelay typex.UBool `json:"tcp_nodelay" flag:"-tcp_nodelay"`
	//Set TCP_NODELAY to disable Nagle’s algorithm. Default value is 0.

	//Remark: Writing to the socket is currently not optimized to minimize system calls and reduces the efficiency / effect of TCP_NODELAY.
}

/*

ffplay rtmp://myserver/vod/sample
To publish to a password protected server, passing the playpath and app names separately:

ffmpeg -re -i <input> -f flv -rtmp_playpath some/long/path -rtmp_app long/app/name rtmp://username:password@myserver
*/

type RtmpLive string

const (
	RtmpLive_any      RtmpLive = "any"
	RtmpLive_live     RtmpLive = "live"
	RtmpLive_recorded RtmpLive = "recorded"
)
