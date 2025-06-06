package initialize

import (
	"github.com/robfig/cron/v3"
	"go-blog/global"
	"go-blog/task"
	"go.uber.org/zap"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keysAndValues", keysAndValues))
}

func (l *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Error(err), zap.Any("keysAndValues", keysAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{
		logger: global.Log,
	}
}

func InitCron() {
	c := cron.New(cron.WithLogger(NewZapLogger()))

	err := task.RegisterScheduledTasks(c)

	if err != nil {
		global.Log.Error("error: scheduling cron job", zap.Error(err))
		os.Exit(1)
	}
}
