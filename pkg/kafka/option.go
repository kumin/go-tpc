package kafka

type ConsumerOptFunc func(opt *ConsumerOption) *ConsumerOption

type ConsumerOption struct {
	Brokers    []string
	Group      string
	Topic      []string
	Oldest     bool
	AutoCommit bool
}

var DefaultConsumerOption = &ConsumerOption{
	Brokers:    []string{"localhost:9092"},
	Group:      "group_1",
	Topic:      []string{},
	Oldest:     true,
	AutoCommit: true,
}

func WithConsumerBrokers(brokers []string) ConsumerOptFunc {
	return func(opt *ConsumerOption) *ConsumerOption {
		opt.Brokers = brokers
		return opt
	}
}

func WithConsumerGroup(group string) ConsumerOptFunc {
	return func(opt *ConsumerOption) *ConsumerOption {
		opt.Group = group
		return opt
	}
}

func WithConsumerOldest(oldest bool) ConsumerOptFunc {
	return func(opt *ConsumerOption) *ConsumerOption {
		opt.Oldest = oldest
		return opt
	}
}

func WithConsumerAutoCommit(autoCommit bool) ConsumerOptFunc {
	return func(opt *ConsumerOption) *ConsumerOption {
		opt.AutoCommit = autoCommit
		return opt
	}
}

type ProducerOptionFnc func(opt *ProducerOption) *ProducerOption

type ProducerOption struct {
	Brokers      []string
	Topic        string
	NumProducers int
}

var DefaultProducerOpt = &ProducerOption{
	Brokers:      []string{"localhost:9092"},
	NumProducers: 1,
}

func WithProducerBrokers(brokers []string) ProducerOptionFnc {
	return func(opt *ProducerOption) *ProducerOption {
		opt.Brokers = brokers
		return opt
	}
}

func WithNumProducers(numProducers int) ProducerOptionFnc {
	return func(opt *ProducerOption) *ProducerOption {
		opt.NumProducers = numProducers
		return opt
	}
}

func WithProducerTopic(topic string) ProducerOptionFnc {
	return func(opt *ProducerOption) *ProducerOption {
		opt.Topic = topic
		return opt
	}
}
