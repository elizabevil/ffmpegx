package transcoderx

import (
	"encoding/json"
	"github.com/elizabevil/ffmpegx/metadatax"
	"os/exec"
)

type CmdHandle = func(cmd *exec.Cmd)

type OutByteHandle = func(out []byte, err []byte)

type Unmarshal struct {
}

func (u Unmarshal) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}

var NewProgressMaker = metadatax.NewDefaultProgress

var JsonUnmarshal = Unmarshal{}
