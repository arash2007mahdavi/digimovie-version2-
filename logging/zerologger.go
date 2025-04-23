package logging

import (
	"digimovie/src/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var zerologLevels = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info": zerolog.InfoLevel,
	"warn": zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

type zerologger struct {
	cfg *config.Config
	logger *zerolog.Logger
}

func NewZerologger(cfg *config.Config) *zerologger {
	l := &zerologger{}
	l.cfg = cfg
	l.Init()
	return l
}

func (l *zerologger) getLogLevel() zerolog.Level {
	level := l.cfg.Logger.Level
	zerolevel, ok := zerologLevels[level]
	if ok {
		return zerolevel
	}
	return zerolog.DebugLevel
}

func (l *zerologger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(l.getLogLevel())

	file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(file).With().Timestamp().Str("AppName", "DigiMovie").Str("LoggerType", "Zerologger").Logger()
	l.logger = &logger
}

func extraTOparasm(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}
	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}

func (l *zerologger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparasm(extra)
	l.logger.Debug().
		Str("Category", string(cat)).Str("SubCategory", string(sub)).
		Fields(params).Msg(msg)
}
func (l *zerologger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zerologger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparasm(extra)
	l.logger.Info().
		Str("Category", string(cat)).Str("SubCategory", string(sub)).
		Fields(params).Msg(msg)
}
func (l *zerologger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zerologger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparasm(extra)
	l.logger.Warn().
		Str("Category", string(cat)).Str("SubCategory", string(sub)).
		Fields(params).Msg(msg)
}
func (l *zerologger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zerologger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparasm(extra)
	l.logger.Error().
		Str("Category", string(cat)).Str("SubCategory", string(sub)).
		Fields(params).Msg(msg)
}
func (l *zerologger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zerologger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparasm(extra)
	l.logger.Fatal().
		Str("Category", string(cat)).Str("SubCategory", string(sub)).
		Fields(params).Msg(msg)
}
func (l *zerologger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}