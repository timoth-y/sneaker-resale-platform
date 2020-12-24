package events

import (
	"github.com/golang/glog"
	"github.com/streadway/amqp"

	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/util"
)

type Broker struct {
	Channel *amqp.Channel
	Exchange string
	core.Serializer
}

func NewEventsBroker(config config.ConnectionConfig, exchange string, serializer core.Serializer) *Broker {
	conn, err := amqp.DialTLS(config.URL, util.NewTLSConfig(config.TLS)); if err != nil {
		glog.Fatal(err)
	}

	ch, err := conn.Channel(); if err != nil {
		glog.Fatal(err)
	}

	return &Broker{
		Channel: ch,
		Exchange: exchange,
		Serializer: serializer,
	}
}