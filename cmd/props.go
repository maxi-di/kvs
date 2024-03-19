package cmd

import (
	"kvs/kvs"

	"github.com/sirupsen/logrus"
)

type Props struct {
	logger  *logrus.Logger
	storage kvs.KvStorage
}

func NewProps() *Props {
	return &Props{}
}

func InitProps(props *Props, logger *logrus.Logger, storage kvs.KvStorage) {
	props.logger = logger
	props.storage = storage
}
