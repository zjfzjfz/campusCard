package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Log *logrus.Logger
var File *os.File

func init() {
    Log = logrus.New()
    Log.SetLevel(logrus.InfoLevel)

    File, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        Log.Out = File
    } else {
        Log.Info("Failed to log to file, using default stderr")
    }
}

func FileClose(){
	File.Close()
}