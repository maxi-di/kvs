package main

import (
	"kvs/cmd"
	"kvs/kvs"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	appName = "kvs"
)

var (
	logger   *logrus.Logger
	revision = "uknown"
	verbose  = true
	location = ""
)

func main() {
	props := cmd.NewProps()

	logger = initLogger(logrus.PanicLevel)

	rootCmd := &cobra.Command{
		Use:  appName,
		Long: "Key value storage",
		Run: func(c *cobra.Command, args []string) {
		},
		PersistentPreRun: func(c *cobra.Command, args []string) {
			if verbose {
				logger.SetLevel(logrus.TraceLevel)
			}
			if location == "" {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					logger.Fatal("can't get user home dir")
				}
				location = path.Join(homeDir, ".local", "share", "kvs")
			}
			storage, err := kvs.NewFileStorage(location, logger)
			if err != nil {
				logger.Fatal(err)
			}
			cmd.InitProps(props, logger, storage)
		},
	}
	rootCmd.Version = revision
	rootCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "location (path) of the storage")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose for logger output")

	rootCmd.AddCommand(cmd.NewListDBCmd(props))
	rootCmd.AddCommand(cmd.NewNewDBCmd(props))
	rootCmd.AddCommand(cmd.NewRemoveDBCmd(props))
	rootCmd.AddCommand(cmd.NewInsertCmd(props))
	rootCmd.AddCommand(cmd.NewGetCmd(props))
	rootCmd.AddCommand(cmd.NewRemoveCmd(props))

	rootCmd.Execute()
}

func initLogger(level logrus.Level) *logrus.Logger {

	logger := logrus.New()
	logger.SetOutput(os.Stderr)
	logger.ReportCaller = true
	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := " " + path.Base(f.File) + ":" + strconv.Itoa(f.Line)
			return function, fileName
		},
	}
	logger.Level = level

	return logger
}
