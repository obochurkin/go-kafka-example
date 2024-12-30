package handlers

import (
	"go-kafka-example/config"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
)

type MessageDTO struct {
	Message string `json:"message" validate:"required,max=255"`
}

type MessageController struct {
	Cfg      *config.Config
	Producer sarama.SyncProducer
}

func (c MessageController) SendMessage(ctx echo.Context) error {
	var msg MessageDTO

	if err := ctx.Bind(&msg); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := ctx.Validate(&msg); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	message := &sarama.ProducerMessage{
		Topic: c.Cfg.Topic,
		Value: sarama.StringEncoder([]byte(msg.Message)),
	}

	partition, offset, err := c.Producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d", partition, offset)

	return ctx.JSON(http.StatusOK, "Message Sent")
}
