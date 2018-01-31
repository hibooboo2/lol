package ddragon

import (
	log "log"
	"os"

	"github.com/comail/colog"
)

var (
	logger    *log.Logger
	logWriter *colog.CoLog
)

func init() {
	logWriter = colog.NewCoLog(os.Stdout, "lol-ddragon:", log.Lshortfile)
	logWriter.SetDefaultLevel(colog.LTrace)
	logWriter.SetMinLevel(colog.LDebug)
	logger = log.New(logWriter, "", 0)
	logger.Println("Logger initialized for lolapi.")
}

// SetLogLevel Set level for logger
func SetLogLevel(level colog.Level) {
	logWriter.SetMinLevel(level)
}
