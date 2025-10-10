package logging

import (
	"os"

	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zeroSyncLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zerologLevelMap = map[string]zerolog.Level{
	"debug":   zerolog.DebugLevel,
	"info":    zerolog.InfoLevel,
	"warning": zerolog.WarnLevel,
	"error":   zerolog.ErrorLevel,
	"fatal":   zerolog.FatalLevel,
}

func newZeroLogLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (zl *zeroLogger) getLogLevel() zerolog.Level {
	level, exist := zerologLevelMap[zl.cfg.Logger.Level]
	if !exist {
		return zerolog.DebugLevel
	}
	return level
}

func (zl *zeroLogger) Init() {
	once.Do(func() {

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		file, err := os.OpenFile(zl.cfg.Logger.Filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			panic("could not open log file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "car-app").
			Str("LoggerName", "zerolog").
			Logger()
		zerolog.SetGlobalLevel(zl.getLogLevel())

		zeroSyncLogger = &logger

	})
	zl.logger = zeroSyncLogger
}

func (zl *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	zl.logger.Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)

}

func (zl *zeroLogger) Debugf(template string, args ...interface{}) {
	zl.logger.Debug().Msgf(template, args...)
}

func (zl *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	zl.logger.Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)

}

func (zl *zeroLogger) Infof(template string, args ...interface{}) {
	zl.logger.Info().Msgf(template, args...)
}

func (zl *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	zl.logger.Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)

}

func (zl *zeroLogger) Warnf(template string, args ...interface{}) {
	zl.logger.Warn().Msgf(template, args...)
}

func (zl *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	zl.logger.Error().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)

}

func (zl *zeroLogger) Errorf(template string, args ...interface{}) {
	zl.logger.Error().Msgf(template, args...)
}

func (zl *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	zl.logger.Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(mapToZeroParams(extra)).
		Msg(msg)

}

func (zl *zeroLogger) Fatalf(template string, args ...interface{}) {
	zl.logger.Fatal().Msgf(template, args...)
}
