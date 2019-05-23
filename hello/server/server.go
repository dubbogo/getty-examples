package main

import (
	"github.com/dubbogo/getty"
	"github.com/dubbogo/getty-examples/hello"
)

func main() {
	server := getty.NewTCPServer(
		getty.WithLocalAddress(":8090"),
	)

	go server.RunEventLoop(newSession)

	log.Info("start hello server")
	hello.WaitCloseSignals(server)
}

func newSession(session getty.Session) (err error) {
	return hello.InitialSession(session, hello.NewHelloPackageHandler(), hello.NewHelloMessageListener())
}
