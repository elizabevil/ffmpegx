package transcoderx

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var Debug bool = false // show cmd string
var DefaultWriter io.Writer = os.Stdout

func IsDebugging() bool {
	return Debug
}

var DebugPrintFunc func(format string, values ...any)

func DebugPrint(format string, values ...any) {
	if IsDebugging() {
		if DebugPrintFunc != nil {
			DebugPrintFunc(format, values...)
			return
		}
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(DefaultWriter, "[FFmpegx-Debug] "+format, values...)
	}
}
