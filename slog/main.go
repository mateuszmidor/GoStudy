package main

import (
	"log"
	"log/slog"
	"os"
)

var logOptions = slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true} // AddSource to include file and line number

func main() {
	// simplest example; default log level is Info, and no lower level will be recorded
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix | log.LUTC | log.Lshortfile)
	log.SetPrefix("[slog-demo] ")
	slog.Info("hello world!", "field1", 42) // 2023/10/08 07:36:06.049905 main.go:15: [slog-demo] INFO hello world! field1=42

	// output log in text format, enable DEBUG level
	textHandler := slog.NewTextHandler(os.Stdout, &logOptions)
	textLogger := slog.New(textHandler)
	slog.SetDefault(textLogger)
	slog.Debug("hello world!", "field1", 42) // time=2023-10-08T10:18:08.170+02:00 level=DEBUG source=/home/user/SoftwareDevelopment/GoStudy/slog/main.go:22 msg="hello world!" field1=42

	// output log in json format, enable DEBUG level
	jsonHandler := slog.NewJSONHandler(os.Stdout, &logOptions)
	jsonLogger := slog.New(jsonHandler)
	slog.SetDefault(jsonLogger)
	slog.Debug("hello world!", slog.Int("field1", 42)) // {"time":"2023-10-08T10:18:08.170959432+02:00","level":"DEBUG","source":{"function":"main.main","file":"/home/user/SoftwareDevelopment/GoStudy/slog/main.go","line":28},"msg":"hello world!","field1":42}
}
