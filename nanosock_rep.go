//
// hilt/nanosock/nanosock_rep.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package nanosock

import (
	"github.com/op/go-nanomsg"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &NanoRepSock{}
)

type NanoRepSock struct {
	NanoSock
	repSock *nanomsg.RepSocket
}

func NewRepSocket() (socket.Socket, error) {
	repSock, err := nanomsg.NewRepSocket()
	if err != nil {
		return nil, err
	}

	err = repSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &NanoRepSock{NanoSock: NanoSock{sock: repSock.Socket}, repSock: repSock}, nil
}
