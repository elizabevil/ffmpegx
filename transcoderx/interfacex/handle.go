package interfacex

import (
	"github.com/elizabevil/ffmpegx/metadatax"
	"io"
	"net/url"
)

type ProtocolOption interface {
	Options() (string, error)
	Url(hostname string) (url.URL, error)
}

type ProgressHandle interface {
	MakeProgress(stream io.ReadCloser, out chan metadatax.Progress)
}
