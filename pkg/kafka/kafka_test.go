package kafka

import (
	"context"
	"log"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/kumin/go-tpc/pkg/envx"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestProducer(t *testing.T) {
	ast := assert.New(t)
	producer, err := NewProducer(
		WithProducerBrokers(envx.GetArray("KAFKA_BROKERS", "localhost:9092")),
	)
	ast.Nil(err)
	defer producer.Close()
	message := &sarama.ProducerMessage{
		Topic: "test_topic",
	}
	message.Value = sarama.StringEncoder("this is test message")
	_, _, err = producer.SendMessage(message)
	ast.Nil(err)
}

func TestConsumer(t *testing.T) {
	ast := assert.New(t)
	consumer, err := NewConsumer(
		WithConsumerBrokers(envx.GetArray("KAFKA_BROKERS", "localhost:9092")),
		WithConsumerOldest(true),
		WithConsumerGroup("test_group1"),
		WithConsumerAutoCommit(true),
	)
	ast.Nil(err)
	consumer.RegisterHandler("test_topic", func(ctx context.Context, value []byte) error {
		mess := string(value)
		log.Print(mess)
		ast.Equal("this is test message", mess)
		return nil
	})
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return consumer.Start(ctx)
	})
	err = eg.Wait()
	ast.Nil(err)
}
