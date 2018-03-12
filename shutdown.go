package main

import (
	"os"
	"strings"

	"github.com/golang/glog"
)

// log a message and shut down the server with the given exit status.
func shutdown(status int, t string, args ...interface{}) {
	if !strings.HasSuffix(t, "\n") {
		t += "\n"
	}
	if status == 0 {
		glog.Infof(t, args...)
	} else {
		glog.Errorf(t, args...)
	}

	// shutting down the main context will propogate to all subcontexts,
	// allowing shutdown handlers to fire on registered context handlers
	ctxShutdown()

	os.Exit(status)
}
