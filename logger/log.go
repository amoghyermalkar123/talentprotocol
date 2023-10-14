package logger

import (
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	logFilePath := "/var/log/tpbackend/service.log"
	logFile, _ := rotatelogs.New(logFilePath + ".%Y%m%d%H%M%S")
	writer := zerolog.MultiLevelWriter(os.Stdout, logFile)

	log := zerolog.New(writer).With().Timestamp().Logger()

	return log
}
