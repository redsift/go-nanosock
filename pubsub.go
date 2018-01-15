package nanosock

import (
	"github.com/redsift/go-mangosock/nano"
	"github.com/op/go-nanomsg"
)

// ensure we implement interfaces correctly
var (
	_ nano.Pub = &pubsock{}
)

type pubsock struct {
	s
	pubSock *nanomsg.PubSocket
}

func NewPubSocket() (nano.Pub, error) {
	pubSock, err := nanomsg.NewPubSocket()
	if err != nil {
		return nil, err
	}

	err = pubSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &pubsock{s: s{sock: pubSock.Socket}, pubSock: pubSock}, nil
}

func (s *pubsock) Publish(data []byte) (int, error) {
	return s.Send(data)
}

// ensure we implement interfaces correctly
var (
	_ nano.Sub = &subsock{}
)

type subsock struct {
	s
	subSock *nanomsg.SubSocket
}

func NewSubSocket() (nano.Sub, error) {
	subSock, err := nanomsg.NewSubSocket()
	if err != nil {
		return nil, err
	}

	err = subSock.SetRecvMaxSize(-1)
	if err != nil {
		return nil, err
	}

	return &subsock{s: s{sock: subSock.Socket}, subSock: subSock}, nil}

func (s *subsock) Subscribe(topic string) error {
	return s.subSock.Subscribe(topic)
}

func (s *subsock) Unsubscribe(topic string) error {
	return s.subSock.Unsubscribe(topic)
}
