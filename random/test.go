package main

import (
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
)

func main() {
	// Define the log file path and rotation settings
	logFilePath := "/path/to/your/logfile.log"
	logFile, _ := rotatelogs.New(logFilePath + ".%Y%m%d%H%M%S")
	writer := zerolog.MultiLevelWriter(os.Stdout, logFile)

	// Create a zerolog logger
	log := zerolog.New(writer).With().Timestamp().Logger()

	// Now, you can use the 'log' instance to log messages to the specified log file
	log.Info().Msg("This is an info message")
}
