package cmq

import (
	"fmt"
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/log"
	"github.com/nats-io/nats.go"
	"sync"
)

type Streaming interface {
	CreateStream(jetStream nats.JetStreamContext) error
	Consume(js nats.JetStreamContext, topic string) error
	Publish(js nats.JetStreamContext, topic string, message any) error
}

type JetStreamWrapper struct {
	sync.Once //lazy initializing for create stream
	client    nats.JetStreamContext
}

func NewJetStreamWrapper() (*JetStreamWrapper, error) {
	nc, err := nats.Connect(config.C.NatsServer.Address)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	return &JetStreamWrapper{client: js}, nil
}

func (nc *JetStreamWrapper) CreateStream(jetStream nats.JetStreamContext) error {
	stream, err := jetStream.StreamInfo(config.C.NatsServer.StreamName)
	if stream == nil {
		log.Logger.Info(fmt.Sprintf("Creating stream: %s\n", config.C.NatsServer.StreamName))
		_, err = jetStream.AddStream(&nats.StreamConfig{
			Name:     config.C.NatsServer.StreamName,
			Subjects: []string{config.C.NatsServer.StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (nc *JetStreamWrapper) Consume(js nats.JetStreamContext, topic string) error {
	_, err := js.Subscribe(config.C.NatsServer.SubjectNameMessage, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Logger.Info(fmt.Sprintf("Unable to Ack:%s", err))
			return
		}
	})

	if err != nil {
		log.Logger.Info("Subscribe failed")
		return err
	}

	return nil
}

func (nc *JetStreamWrapper) Publish(js nats.JetStreamContext, topic string, message any) error {
	return nil
}
