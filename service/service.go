package service

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Consumer 代表 Sarama 消费者 消费组
type Consumer struct {
	ready chan bool
}

// 启动 服务
func StartService()  {
	log.Println(op)
	// 调起服务
	service()
}

func service() {

	// 日志
	if op.KafkaConsumer.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	// 解析 kafka 版本
	version, err := sarama.ParseKafkaVersion(op.KafkaConsumer.Version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	// 使用组管理时，客户端将用于在使用者实例之间分发分区所有权的分区分配策略的类名
	switch op.KafkaConsumer.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Printf("无法识别的使用者组分区分配程序: %s", assignor)
		os.Exit(0)
	}

	//如果以前未提交偏移量，则使用的初始偏移量。
	//应该是OffsetNewest或OffsetOldest。默认为OffsetLatest。
	if op.KafkaConsumer.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	// 设置一个新的 Sarama 消费组
	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(op.KafkaConsumer.Brokers, ","), op.KafkaConsumer.Group, config)
	if err != nil {
		log.Printf("创建使用者组客户端时出错: %v", err)
		os.Exit(0)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(op.KafkaConsumer.Topics, ","), &consumer); err != nil {
				log.Printf("来自消费者的错误: %v", err)
				os.Exit(0)
			}
			// 检查上下文是否已取消，表示使用者应停止
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // 等待消费者被设置好
	log.Println("Sarama 消费者 启动运行...")

	sigterm := make(chan os.Signal, 1)

	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}

}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// 将消费者标记为就绪
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		//log.Printf("Message claimed: offset = %d, key = %s, value = %s, timestamp = %v, topic = %s", message.Offset, string(message.Key), string(message.Value), message.Timestamp, message.Topic)
		//session.MarkMessage(message, "")
		execute(session, message)
	}
	return nil
}