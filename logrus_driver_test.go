package slf4go_logrus

import (
	slog "github.com/go-eden/slf4go"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	Init()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	slog.Debug("hello")
	slog.Info("what???")
	slog.Warnf("warnning: %v", "surrender")

	l := slog.GetLogger()
	l.BindFields(slog.Fields{
		"logger": "test",
	})
	l.Errorf("error!!! %v", 100)
}
