package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"src/utils/config"
)

func InitLogrus() {
	output, err := newLogWriter(config.Conf.Logrus.OutputFile)
	if err == nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(output)
	} else {
		logrus.Error(err)
	}
}

type LogWriter struct {
	file *os.File
}

func newLogWriter(logPath string) (*LogWriter, error) {
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &LogWriter{file: file}, nil
}

func (this *LogWriter) Write(b []byte) (n int, err error) {
	_, _ = os.Stdout.Write(b)
	return this.file.Write(b)
}
