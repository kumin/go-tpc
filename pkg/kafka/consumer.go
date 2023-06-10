package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"golang.org/x/sync/errgroup"
)

type handlFunc func(ctx context.Context, value []byte) error

type Consumer struct {
	client   sarama.ConsumerGroup
	handlers map[string]handlFunc
	options  *ConsumerOption
}

func NewConsumer(optFns ...ConsumerOptFunc) (*Consumer, error) {
	opt := DefaultConsumerOption
	for _, optFn := range optFns {
		opt = optFn(opt)
	}
	saramaCfg := sarama.NewConfig()
	if opt.Oldest {
		saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
	saramaCfg.Consumer.Offsets.AutoCommit.Enable = opt.AutoCommit
	client, err := sarama.NewConsumerGroup(opt.Brokers, opt.Group, saramaCfg)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		client:   client,
		handlers: make(map[string]handlFunc),
		options:  opt,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	topics := make([]string, 0, len(c.handlers))
	for k := range c.handlers {
		topics = append(topics, k)
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		for {
			if err := c.client.Consume(ctx, topics, c); err != nil {
				return err
			}
			if ctx.Err() != nil {
				return ctx.Err()
			}
		}
	})
	return eg.Wait()
}

func (c *Consumer) RegisterHandler(topic string, handler handlFunc) {
	c.handlers[topic] = handler
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if h, ok := c.handlers[message.Topic]; ok {
				if err := h(session.Context(), message.Value); err != nil {
					return err
				}
			}
			if !c.options.AutoCommit {
				fmt.Println("commit offset")
				// session.MarkOffset(message)
				session.MarkMessage(message, "")
				session.Commit()
			}
		case <-session.Context().Done():
			return nil
		}
	}
}
