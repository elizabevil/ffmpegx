package interfacex

import (
	"context"
	"github.com/elizabevil/ffmpegx/metadatax"
	"io"
	"net/url"
)

type ProtocolOption interface {
	Options() (string, error)
	Url(hostname string) (url.URL, error)
}

type ProgressHandle interface {
	MakeProgress(ctx context.Context, stream io.ReadCloser, out chan metadatax.Progress)

	MakePlayProgress(ctx context.Context, stream io.ReadCloser, out chan metadatax.FFplay)
}
