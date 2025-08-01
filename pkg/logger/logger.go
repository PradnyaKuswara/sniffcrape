package logger

import (
	"bytes"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

type CustomWriter struct {
	console io.Writer
	file    io.Writer
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	w.console.Write(p)

	clean := stripANSI(p)
	return w.file.Write(clean)
}

func stripANSI(input []byte) []byte {
	ansiEscape := []byte("\x1b")
	var output bytes.Buffer
	inEscape := false

	for i := 0; i < len(input); i++ {
		if input[i] == ansiEscape[0] {
			inEscape = true
			continue
		}
		if inEscape {
			if (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') {
				inEscape = false
			}
			continue
		}
		output.WriteByte(input[i])
	}
	return output.Bytes()
}

func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// Buat direktori jika belum ada
	logDir := "storage/logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		logrus.Fatalf("Failed to create log directory: %v", err)
	}

	// Buka file log
	logFile := logDir + "/app.log"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	// Set output ke console + file
	Log.SetOutput(&CustomWriter{
		console: os.Stdout,
		file:    file,
	})

	Log.SetLevel(logrus.DebugLevel)
}
