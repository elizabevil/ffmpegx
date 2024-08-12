package metadatax

import (
	"context"
	"os"
)

type ProgressHandle = func(progress Progress)
type ProcessCtxHandle = func(ctx context.Context, handle func(process *os.Process), progressHandles ...ProgressHandle) error

func NewDefaultProgress() DefaultProgress { return DefaultProgress{} }
