package main

import (
	"github.com/jbowles/disfun/Godeps/_workspace/src/github.com/juju/loggo"
)

var first = loggo.GetLogger("first")

func FirstCritical(message string) {
	first.Criticalf(message)
}

func FirstError(message string) {
	first.Errorf(message)
}

func FirstWarning(message string) {
	first.Warningf(message)
}

func FirstInfo(message string) {
	first.Infof(message)
}

func FirstTrace(message string) {
	first.Tracef(message)
}
