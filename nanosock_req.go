//
// hilt/nanosock/nanosock_req.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package nanosock

import (
	"time"

	"github.com/op/go-nanomsg"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &NanoReqSock{}
)

type NanoReqSock struct {
	NanoSock
	reqSock *nanomsg.ReqSocket
}

func NewReqSocket() (socket.Socket, error) {
	reqSock, err := nanomsg.NewReqSocket()
	if err != nil {
		return nil, err
	}

	err = reqSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &NanoReqSock{NanoSock: NanoSock{sock: reqSock.Socket}, reqSock: reqSock}, nil
}

func (s *NanoReqSock) SetResendInterval(interval time.Duration) error {
	return s.reqSock.SetResendInterval(interval)
}
