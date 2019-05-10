package hello

import (
	"github.com/dubbogo/getty"
	"log"
	"time"
)

// ------------------------------------------------------
// message handler
type MessageHandler struct {
	SessionOnOpen func(session getty.Session)
}

func NewHelloMessageListener() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) OnOpen(session getty.Session) error {
	log.Printf("OnOpen session{%s} open", session.Stat())
	if h.SessionOnOpen != nil {
		h.SessionOnOpen(session)
	}
	return nil
}

func (h *MessageHandler) OnError(session getty.Session, err error) {
	log.Printf("OnError session{%s} got error{%v}, will be closed.", session.Stat(), err)
}

func (h *MessageHandler) OnClose(session getty.Session) {
	log.Printf("OnClose session{%s} is closing......", session.Stat())
}

func (h *MessageHandler) OnMessage(session getty.Session, pkg interface{}) {
	s, ok := pkg.(string)
	if !ok {
		log.Printf("illegal packge{%#v}", pkg)
		return
	}

	log.Printf("OnMessage: %s", s)
}

func (h *MessageHandler) OnCron(session getty.Session) {
	active := session.GetActive()
	if CronPeriod.Nanoseconds() < time.Since(active).Nanoseconds() {
		log.Printf("OnCorn session{%s} timeout{%s}", session.Stat(), time.Since(active).String())
		session.Close()
	}
}
