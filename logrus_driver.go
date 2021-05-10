package slf4go_logrus

import (
	slog "github.com/go-eden/slf4go"
	log "github.com/sirupsen/logrus"
)

// Init setup slf4go's driver atomically
func Init() {
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
	case slog.TraceLevel:
		if l.Format == nil {
			entry.Trace(l.Args...)
		} else {
			entry.Tracef(*l.Format, l.Args...)
		}
	case slog.DebugLevel:
		if l.Format == nil {
			entry.Debug(l.Args...)
		} else {
			entry.Debugf(*l.Format, l.Args...)
		}
	case slog.InfoLevel:
		if l.Format == nil {
			entry.Info(l.Args...)
		} else {
			entry.Infof(*l.Format, l.Args...)
		}
	case slog.WarnLevel:
		if l.Format == nil {
			entry.Warn(l.Args...)
		} else {
			entry.Warnf(*l.Format, l.Args...)
		}
	case slog.ErrorLevel:
		if l.Format == nil {
			entry.Error(l.Args...)
		} else {
			entry.Errorf(*l.Format, l.Args...)
		}
	case slog.PanicLevel:
		if l.Format == nil {
			entry.Panic(l.Args...)
		} else {
			entry.Panicf(*l.Format, l.Args...)
		}
	case slog.FatalLevel:
		if l.Format == nil {
			entry.Fatal(l.Args...)
		} else {
			entry.Fatalf(*l.Format, l.Args...)
		}
	}
}

func (d *LogrusDriver) GetLevel(logger string) (sl slog.Level) {
	switch log.GetLevel() {
	case log.TraceLevel:
		sl = slog.TraceLevel
	case log.DebugLevel:
		sl = slog.DebugLevel
	case log.InfoLevel:
		sl = slog.InfoLevel
	case log.WarnLevel:
		sl = slog.WarnLevel
	case log.ErrorLevel:
		sl = slog.ErrorLevel
	case log.PanicLevel:
		sl = slog.PanicLevel
	case log.FatalLevel:
		sl = slog.FatalLevel
	}
	return slog.TraceLevel // don't know
}
