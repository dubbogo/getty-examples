package main

import (
	"github.com/dubbogo/getty"
	"github.com/dubbogo/getty-examples/hello"

	"math/rand"
	"sync"
	"time"
)

const (
	WritePkgTimeout = 1e8
)

var (
	client   getty.Client
	sessions []getty.Session
	lock     sync.RWMutex
)

func main() {
	client = getty.NewTCPClient(
		getty.WithServerAddress("127.0.0.1:8090"),
		getty.WithConnectionNumber(2),
	)

	client.RunEventLoop(newSession)

	go test()

	hello.WaitCloseSignals(client)
}

func newSession(session getty.Session) (err error) {
	listener := hello.NewHelloMessageListener()
	listener.SessionOnOpen = func(session getty.Session) {
		sessions = append(sessions, session)
	}
	return hello.InitialSession(session, hello.NewHelloPackageHandler(), listener)
}

func test() {
	for {
		if selectSession() != nil {
			break
		}
		time.Sleep(time.Second)
	}
	echoTimes := 10

	counter := getty.CountWatch{}
	counter.Start()
	for i := 0; i < echoTimes; i++ {
		session := selectSession()
		err := session.WritePkg("hello", WritePkgTimeout)
		if err != nil {
			log.Infof("session.WritePkg(session{%s}, error{%v}", session.Stat(), err)
			session.Close()
			removeSession(session)
		}
	}
	cost := counter.Count()
	log.Infof("after loop %d times, echo cost %d ms", echoTimes, cost/1e6)
}

func selectSession() getty.Session {
	lock.RLock()
	defer lock.RUnlock()
	count := len(sessions)
	if count == 0 {
		log.Infof("client session array is nil...")
		return nil
	}

	return sessions[rand.Int31n(int32(count))]
}

func removeSession(session getty.Session) {
	if session == nil {
		return
	}
	lock.Lock()
	for i, s := range sessions {
		if s == session {
			sessions = append(sessions[:i], sessions[i+1:]...)
			log.Infof("delete session{%s}, its index{%d}", session.Stat(), i)
			break
		}
	}
	log.Infof("after remove session{%s}, left session number:%d", session.Stat(), len(sessions))
	lock.Unlock()
}
