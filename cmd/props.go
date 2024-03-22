package cmd

import (
	"kvs/kvs"

	"github.com/sirupsen/logrus"
)

type Props struct {
	logger  *logrus.Logger
	storage kvs.Storage
}

func InitProps(props *Props, logger *logrus.Logger, storage kvs.Storage) {
	props.logger = logger
	props.storage = storage
}
