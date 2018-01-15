//
// socket.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package nanosock

import (
	"time"

	"github.com/op/go-nanomsg"
	"github.com/redsift/go-mangosock/nano"
)

// ensure we implement interfaces correctly
var (
	_ nano.Socket = &s{}
)

type s struct {
	sock *nanomsg.Socket
	addr string
}

func (s *s) Bind(addr string) error {
	s.addr = addr
	_, err := s.sock.Bind(addr)
	return err
}

func (s *s) Connect(addr string) error {
	s.addr = addr
	_, err := s.sock.Connect(addr)
	return err
}

func (s *s) SetSendTimeout(timeout time.Duration) error {
	return s.sock.SetSendTimeout(timeout)
}

func (s *s) SetRecvTimeout(timeout time.Duration) error {
	return s.sock.SetRecvTimeout(timeout)
}

func (s *s) SetRecvMaxSize(size int64) error {
	return s.sock.SetRecvMaxSize(size)
}

func (s *s) Send(data []byte) (int, error) {
	// deepak: go -> cgo bridge doesn't like go pointers and go seems to optimize string vars.
	// We copy the byte array to avoid any go pointer issues in cgo land.
	// fix "panic: runtime error: cgo argument has Go pointer to Go pointer" (github.com/op/go-nanomsg/nanomsg.go:144)
	if len(data) < 65 {
		dst := make([]byte, len(data), len(data))
		copy(dst, data)
		data = dst
	}
	return s.sock.Send(data, 0)
}

func (s *s) Recv() ([]byte, error) {
	return s.sock.Recv(0)
}

func (s *s) Close() error {
	return s.sock.Close()
}
