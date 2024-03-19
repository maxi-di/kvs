package main

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logger *logrus.Logger
)

func main() {
	rootCmd := &cobra.Command{
		Use: "kvs [COMMAND] [ARG]",
		Run: func(cmd *cobra.Command, args []string) {
			logger = initLogger()
			logger.Info("Hello")
			// logger.Infof("gosnmpserver revision: %s\n", revision)
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal(err)
	}
}

func initLogger() *logrus.Logger {

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.ReportCaller = true
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := " " + path.Base(f.File) + ":" + strconv.Itoa(f.Line)
			return function, fileName
		},
	}
	logger.Level = logrus.TraceLevel

	return logger
}
