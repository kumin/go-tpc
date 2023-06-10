package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Producer struct {
	client sarama.SyncProducer
}

func NewProducer(optFns ...ProducerOptionFnc) (*Producer, error) {
	opt := DefaultProducerOpt
	for _, optFn := range optFns {
		opt = optFn(opt)
	}
	cfgs := sarama.NewConfig()
	cfgs.Producer.RequiredAcks = sarama.WaitForAll
	cfgs.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer(opt.Brokers, cfgs)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &Producer{
		client: client,
	}, nil
}

func (p *Producer) SendMessage(message *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	partition, offset, err = p.client.SendMessage(message)
	return
}

func (p *Producer) Close() {
	p.client.Close()
}
