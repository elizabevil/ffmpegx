package metadatax

import (
	"bufio"
	"bytes"
	"context"
	"io"
)

type DefaultProgress struct {
	Filter func(str string) bool
}

func (r DefaultProgress) makeScanner(stream io.ReadCloser) *bufio.Scanner {
	split := func(data []byte, atEOF bool) (advance int, token []byte, spliterror error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, '\n'); i >= 0 {
			// We have a full newline-terminated line.
			return i + 1, data[0:i], nil
		}
		if i := bytes.IndexByte(data, '\r'); i >= 0 {
			// We have a cr terminated line
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
	scanner := bufio.NewScanner(stream)
	scanner.Split(split)
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	return scanner
}

func (r DefaultProgress) MakeProgress(ctx context.Context, stream io.ReadCloser, out chan Progress) {
	if ctx == nil || ctx.Err() != nil {
		return
	}
	go r.progressHandle(ctx, r.makeScanner(stream), out)
	select {
	case <-ctx.Done():
		return
	}
}
func (r DefaultProgress) progressHandle(ctx context.Context, scanner *bufio.Scanner, out chan Progress) {
	next := false
	for scanner.Scan() {
		line := scanner.Text()
		if r.Filter != nil {
			next = r.Filter(line)
		} else {
			next = defaultFilterFunc(line)
		}
		if next && ctx.Err() == nil {
			pp := defaultProgressPool.GetProgress()
			makeProgress(line, &pp)
			out <- pp
			pp.Reset()
			defaultProgressPool.PutProgress(pp)

		}
	}
}
