package metadatax

import (
	"sync"
	"unique"
)

var emptysString = unique.Make("")

var defaultProgressPool = NewProgressPool()
var defaultFFplayPool = NewFFplayPool()

type ProgressPool struct {
	sync.Pool
}

func NewProgressPool() ProgressPool {
	return ProgressPool{
		sync.Pool{New: func() any {
			return Progress{}
		}},
	}
}

func (p *ProgressPool) GetProgress() Progress {
	return p.Get().(Progress)
}
func (p *ProgressPool) PutProgress(px Progress) {
	p.Put(px)
}

type FFplayPool struct {
	sync.Pool
}

func NewFFplayPool() FFplayPool {
	return FFplayPool{
		sync.Pool{New: func() any {
			return FFplay{}
		}},
	}
}

func (p *FFplayPool) GetFFplay() FFplay {
	return p.Get().(FFplay)
}
func (p *FFplayPool) PutFFplay(px FFplay) {
	p.Put(px)
}
