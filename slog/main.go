package main

import (
	"log"
	"log/slog"
	"os"
)

var logOptions = slog.HandlerOptions{Level: slog.LevelDebug}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix | log.LUTC | log.Lshortfile)
	log.SetPrefix("[slog-demo] ")

	// simplest example; default log level is Info, and no lower level will be recorded
	slog.Info("hello world!", "field1", 42) // 2023/10/08 07:36:06.049905 main.go:15: [slog-demo] INFO hello world! field1=42

	// output log in text format, enable DEBUG level
	textTandler := slog.NewTextHandler(os.Stdout, &logOptions)
	textLogger := slog.New(textTandler)
	slog.SetDefault(textLogger)
	slog.Debug("hello world!", "field1", 42) // time=2023-10-08T09:38:08.844+02:00 level=DEBUG msg="hello world!" field1=42

	// output log in json format, enable DEBUG level
	jsonHandler := slog.NewJSONHandler(os.Stdout, &logOptions)
	jsonLogger := slog.New(jsonHandler)
	slog.SetDefault(jsonLogger)
	slog.Debug("hello world!", "field1", 42) // {"time":"2023-10-08T09:42:43.184688767+02:00","level":"DEBUG","msg":"hello world!","field1":42}
}
