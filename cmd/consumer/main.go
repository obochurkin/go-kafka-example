package main

import (
	"context"
	"go-kafka-example/config"
	"log"

	"github.com/IBM/sarama"
)

// Consumer implements the ConsumerGroupHandler interface
type Consumer struct{}

// Setup is run at the beginning of a new session
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes Kafka messages
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message claimed: topic=%s partition=%d offset=%d value=%s",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		session.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	consumerGroup := "message-handlers"
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	brokers := []string{cfg.KakaConnection}

	// Configure Sarama
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create a new consumer group
	consumer, err := sarama.NewConsumerGroup(brokers, consumerGroup, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama consumer group: %v", err)
	}
	defer consumer.Close()

	// Consume messages
	for {
		err := consumer.Consume(context.Background(), []string{cfg.Topic}, &Consumer{})
		if err != nil {
			log.Fatalf("Error during consuming: %v", err)
		}
	}
}
