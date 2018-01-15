//
// reqrep.go
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
	_ nano.Req = &reqsock{}
)

type reqsock struct {
	s
	reqSock *nanomsg.ReqSocket
}

func NewReqSocket() (nano.Req, error) {
	reqSock, err := nanomsg.NewReqSocket()
	if err != nil {
		return nil, err
	}

	err = reqSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &reqsock{s: s{sock: reqSock.Socket}, reqSock: reqSock}, nil
}

func (s *reqsock) SetResendInterval(interval time.Duration) error {
	return s.reqSock.SetResendInterval(interval)
}


// ensure we implement interfaces correctly
var (
	_ nano.Rep = &repsock{}
)

type repsock struct {
	s
	repSock *nanomsg.RepSocket
}

func NewRepSocket() (nano.Rep, error) {
	repSock, err := nanomsg.NewRepSocket()
	if err != nil {
		return nil, err
	}

	err = repSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &repsock{s: s{sock: repSock.Socket}, repSock: repSock}, nil
}

func (s *repsock) Address() string {
	return s.s.addr
}
