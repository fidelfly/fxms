package mskit

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fidelfly/fxgo/logx"
	"github.com/micro/go-micro/util/log"
)

const dateFormat = "2006-01-02 15:04:05"

func SetupLog(level string, name string) {
	logx.SetFormatter(logx.JSONFormatter(dateFormat))
	if level, err := logx.ParseLevel(level); err != nil {
		logx.SetLevel(logx.WarnLevel)
	} else {
		logx.SetLevel(level)
	}

	rotateLog := logx.RotateLog(fmt.Sprintf("./logs/%s.log", strings.ReplaceAll(name, ".", "_")), 1, 7, 7, true)

	logx.SetOutput(io.MultiWriter(os.Stdout, rotateLog))

	log.SetLogger(logx.WithFields(logx.Fields{
		"service": name,
	}))
}
