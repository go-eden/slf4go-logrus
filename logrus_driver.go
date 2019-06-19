package slf4go_logrus

import (
	slog "github.com/go-eden/slf4go"
	log "github.com/sirupsen/logrus"
)

// init setup slf4go's driver atomically
func init() {
	slog.SetDriver(&LogrusDriver{})
}

type LogrusDriver struct {
}

func (d *LogrusDriver) Name() string {
	return "logrus"
}

func (d *LogrusDriver) Print(l *slog.Log) {
	var entry *log.Entry
	// preset fields
	if l.Fields != nil {
		entry = log.WithFields(log.Fields(l.Fields))
	} else {
		entry = log.WithFields(log.Fields{})
	}
	// execute logging
	switch l.Level {
	case slog.LEVEL_TRACE:
		if l.Format == nil {
			entry.Trace(l.Args...)
		} else {
			entry.Tracef(*l.Format, l.Args...)
		}
	case slog.LEVEL_DEBUG:
		if l.Format == nil {
			entry.Debug(l.Args...)
		} else {
			entry.Debugf(*l.Format, l.Args...)
		}
	case slog.LEVEL_INFO:
		if l.Format == nil {
			entry.Info(l.Args...)
		} else {
			entry.Infof(*l.Format, l.Args...)
		}
	case slog.LEVEL_WARN:
		if l.Format == nil {
			entry.Warn(l.Args...)
		} else {
			entry.Warnf(*l.Format, l.Args...)
		}
	case slog.LEVEL_ERROR:
		if l.Format == nil {
			entry.Error(l.Args...)
		} else {
			entry.Errorf(*l.Format, l.Args...)
		}
	case slog.LEVEL_PANIC:
		if l.Format == nil {
			entry.Panic(l.Args...)
		} else {
			entry.Panicf(*l.Format, l.Args...)
		}
	case slog.LEVEL_FATAL:
		if l.Format == nil {
			entry.Fatal(l.Args...)
		} else {
			entry.Fatalf(*l.Format, l.Args...)
		}
	}
}

func (d *LogrusDriver) GetLevel(logger string) slog.Level {
	return slog.LEVEL_TRACE // don't know
}
