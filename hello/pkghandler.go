package hello

import (
	"errors"
	"github.com/dubbogo/getty"
	"log"
)

// ------------------------------------------------------
// package handler
type PackageHandler struct {
}

func NewHelloPackageHandler() *PackageHandler {
	return &PackageHandler{}
}

func (h *PackageHandler) Read(ss getty.Session, data []byte) (interface{}, int, error) {
	s := string(data)
	return s, len(s), nil
}

func (h *PackageHandler) Write(ss getty.Session, pkg interface{}) error {
	s, ok := pkg.(string)
	if !ok {
		log.Printf("illegal pkg:%+v", pkg)
		return errors.New("invalid package")
	}
	return ss.WriteBytes([]byte(s))
}
