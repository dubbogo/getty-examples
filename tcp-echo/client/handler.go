/******************************************************
# DESC    : echo package handler
# AUTHOR  : Alex Stocks
# LICENCE : Apache License 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2016-09-04 13:08
# FILE    : handler.go
******************************************************/

package main

import (
	"errors"

	"time"
)

import (
	"github.com/dubbogo/getty"

)

var (
	errSessionNotExist = errors.New("session not exist!")
)

////////////////////////////////////////////
// EchoMessageHandler
////////////////////////////////////////////

type clientEchoSession struct {
	session getty.Session
	reqNum  int32
}

type EchoMessageHandler struct{}

func newEchoMessageHandler() *EchoMessageHandler {
	return &EchoMessageHandler{}
}

func (h *EchoMessageHandler) OnOpen(session getty.Session) error {
	client.addSession(session)

	return nil
}

func (h *EchoMessageHandler) OnError(session getty.Session, err error) {
	log.Infof("session{%s} got error{%v}, will be closed.", session.Stat(), err)
	client.removeSession(session)
}

func (h *EchoMessageHandler) OnClose(session getty.Session) {
	log.Infof("session{%s} is closing......", session.Stat())
	client.removeSession(session)
}

func (h *EchoMessageHandler) OnMessage(session getty.Session, pkg interface{}) {
	p, ok := pkg.(*EchoPackage)
	if !ok {
		log.Errorf("illegal packge{%#v}", pkg)
		return
	}

	log.Debugf("get echo package{%s}", p)
	client.updateSession(session)
}

func (h *EchoMessageHandler) OnCron(session getty.Session) {
	clientEchoSession, err := client.getClientEchoSession(session)
	if err != nil {
		log.Errorf("client.getClientSession(session{%s}) = error{%#v}", session.Stat(), err)
		return
	}
	if conf.sessionTimeout.Nanoseconds() < time.Since(session.GetActive()).Nanoseconds() {
		log.Warnf("session{%s} timeout{%s}, reqNum{%d}",
			session.Stat(), time.Since(session.GetActive()).String(), clientEchoSession.reqNum)
		client.removeSession(session)
		return
	}

	client.heartbeat(session)
}
