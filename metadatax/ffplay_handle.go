package metadatax

import (
	"bufio"
	"context"
	"io"
	"strings"
)

func (r DefaultProgress) MakePlayProgress(ctx context.Context, stream io.ReadCloser, out chan FFplay) {
	if ctx == nil || ctx.Err() != nil {
		return
	}
	r.ffplayNoChan(ctx, r.makeScanner(stream), out)
}
func (r DefaultProgress) ffplayNoChan(ctx context.Context, scanner *bufio.Scanner, out chan FFplay) {
	next := false
	for scanner.Scan() {
		line := scanner.Text()
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = strings.Contains(line, "fd=") && strings.Contains(line, "aq=")
		}
		if next {
			select {
			case <-ctx.Done():
				return
			default:
				pp := defaultFFplayPool.GetFFplay()
				makePlayProgress(line, &pp)
				out <- pp
				pp.Reset()
				defaultFFplayPool.PutFFplay(pp)
			}
		}
	}
}
