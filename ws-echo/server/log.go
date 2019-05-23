package main

import (
	"github.com/dubbogo/getty"
	"github.com/dubbogo/getty-examples/logger"
)

var (
	log getty.Logger
)

func init() {
	log = logger.ZapLogger()
}
