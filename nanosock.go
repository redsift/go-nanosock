//
// hilt/nanosock/nanosock.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package nanosock

import (
	"sync"
	"time"

	"github.com/op/go-nanomsg"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &NanoSock{}
)

type NanoSock struct {
	sync.Mutex
	sock *nanomsg.Socket
}

func (s *NanoSock) Bind(addr string) error {
	_, err := s.sock.Bind(addr)
	return err
}

func (s *NanoSock) Connect(addr string) error {
	_, err := s.sock.Connect(addr)
	return err
}

func (s *NanoSock) SetSendTimeout(timeout time.Duration) error {
	return s.sock.SetSendTimeout(timeout)
}

func (s *NanoSock) SetRecvTimeout(timeout time.Duration) error {
	return s.sock.SetRecvTimeout(timeout)
}

func (s *NanoSock) SetResendInterval(timeout time.Duration) error {
	return nil
}

func (s *NanoSock) Send(data []byte) error {
	// deepak: go -> cgo bridge doesn't like go pointers and go seems to optimize string vars.
	// We copy the byte array to avoid any go pointer issues in cgo land.
	// TODO: Limit length of byte array
	dst := make([]byte, len(data), len(data))
	copy(dst, data)
	_, err := s.sock.Send(dst, 0)
	return err
}

func (s *NanoSock) Recv() ([]byte, error) {
	return s.sock.Recv(0)
}

func (s *NanoSock) Close() error {
	return s.sock.Close()
}
