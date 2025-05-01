package logging

import (
	"digimovie/src/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zaplogLevels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zaplogger struct {
	cfg *config.Config
	logger *zap.SugaredLogger
}

func NewZaplogger(cfg *config.Config) *zaplogger {
	l := &zaplogger{}
	l.cfg = cfg
	l.Init()
	return l
}

func (l *zaplogger) getLogLevel() zapcore.Level {
	level := l.cfg.Logger.Level
	Zaplevel, ok := zaplogLevels[level]
	if ok {
		return Zaplevel
	}
	return zapcore.DebugLevel
}

func (l *zaplogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: l.cfg.Logger.FilePath,
		MaxSize: 5,
		MaxAge: 5,
		MaxBackups: 100,
		Compress: true,
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	logger = logger.With("AppName", "DigiMovie", "LoggerType", "Zaplogger")

	l.logger = logger
}

func extraTOparams(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	var params []interface{}
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}

func (l *zaplogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparams(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}
func (l *zaplogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zaplogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparams(extra, cat, sub)
	l.logger.Infow(msg, params...)
}
func (l *zaplogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zaplogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparams(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}
func (l *zaplogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zaplogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparams(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}
func (l *zaplogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zaplogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extraTOparams(extra, cat, sub)
	l.logger.Fatalw(msg, params...)
}
func (l *zaplogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}