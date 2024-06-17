package server

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {

	logFilePath := "./logs/api/"

	log := logrus.New()

	writeMap := createLogWriter(logFilePath)

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//New hook
	//log.Hooks.Add(lfHook)
	log.AddHook(lfHook)

	return func(c *gin.Context) {
		//Start time
		startTime := time.Now()

		//Process request
		c.Next()

		//End time
		endTime := time.Now()

		//Execution time
		latencyTime := endTime.Sub(startTime).Microseconds()

		//Request method
		reqMethod := c.Request.Method

		//Request routing
		reqUri := c.Request.URL.Path

		// status code
		statusCode := c.Writer.Status()

		// request IP
		clientIP := c.ClientIP()

		//Log format
		log.WithFields(logrus.Fields{
			"status_code":      statusCode,
			"latency_time(ms)": latencyTime,
			"client_ip":        clientIP,
			"req_method":       reqMethod,
			"req_uri":          reqUri,
		}).Info()

	}
}

func createLogWriter(fileName string) lfshook.WriterMap {
	//Set rotatelogs
	logWriter, _ := rotatelogs.New(
		//Split file name
		fileName+"%Y-%m-%d.log",

		//Generate soft chain, point to the latest log file

		//Set maximum save time (7 days)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//Set log cutting interval (1 day)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	return writeMap
}
