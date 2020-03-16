package common

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var Logger *logrus.Logger

func InitLog() *logrus.Logger {
	logger := logrus.New()
	logger.AddHook(FsHook())

	Logger = logger

	Logger.WithFields(logrus.Fields{
		"hello":"ok",
	}).Info("hello")

	return Logger
}

func GetLog() *logrus.Logger {
	return Logger

}

func FsHook() logrus.Hook {
	fileName := viper.GetString(`log.path`) + "/" + viper.GetString(`log.filename`)
	fmt.Println(fileName)
	writer, err := rotatelogs.New(
		fileName+".%Y%m%d",
		rotatelogs.WithLinkName(viper.GetString(`log.path`)+"/"+viper.GetString(`log.filename`)),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(100),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	logrus.SetLevel(logrus.DebugLevel)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
	return lfsHook
}
