package logging

import (
	"sync"

	"github.com/mhmojtaba/golang-car-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var zapSyncLogger *zap.SugaredLogger

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var zapLogLevelMap = map[string]zapcore.Level{
	"debug":   zapcore.DebugLevel,
	"info":    zapcore.InfoLevel,
	"warning": zapcore.WarnLevel,
	"error":   zapcore.ErrorLevel,
	"fatal":   zapcore.FatalLevel,
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (zl *zapLogger) getLogLevel() zapcore.Level {
	level, exist := zapLogLevelMap[zl.cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (zl *zapLogger) Init() {
	once.Do(func() {

		writeSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   zl.cfg.Logger.Filepath,
			MaxSize:    5, // megabytes
			MaxBackups: 15,
			MaxAge:     5, //days
			LocalTime:  true,
			Compress:   true, // disabled by default
		})

		config := zap.NewProductionEncoderConfig()

		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(zapcore.NewJSONEncoder(config), writeSyncer, zl.getLogLevel())

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		zapSyncLogger = logger.With("AppName", "car-app", "LoggerName", "zerolog")

	})

	zl.logger = zapSyncLogger
}

func (zl *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getLogKey(extra, cat, sub)

	zl.logger.Debugw(msg, params...)
}

func (zl *zapLogger) Debugf(template string, args ...interface{}) {
	zl.Debugf(template, args...)
}

func (zl *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getLogKey(extra, cat, sub)

	zl.logger.Infow(msg, params...)
}
func (zl *zapLogger) Infof(template string, args ...interface{}) {
	zl.Infof(template, args...)
}

func (zl *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getLogKey(extra, cat, sub)

	zl.logger.Warnw(msg, params...)
}
func (zl *zapLogger) Warnf(template string, args ...interface{}) {
	zl.Warnf(template, args...)
}

func (zl *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getLogKey(extra, cat, sub)

	zl.logger.Errorw(msg, params...)
}

func (zl *zapLogger) Errorf(template string, args ...interface{}) {
	zl.Errorf(template, args...)
}

func (zl *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := getLogKey(extra, cat, sub)

	zl.logger.Fatalw(msg, params...)
}
func (zl *zapLogger) Fatalf(template string, args ...interface{}) {
	zl.Fatalf(template, args...)
}

func getLogKey(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	params := mapToZapParams(extra)
	extra["Category"] = cat
	extra["SubCategory"] = sub
	return params
}
