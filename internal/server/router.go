package server

import (
	"go-kafka-example/config"
	"go-kafka-example/internal/server/handlers"
	"log"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
)

type Router struct {
	hc handlers.HealthCheckController
	mess handlers.MessageController
}

func InitRouter (e *echo.Echo, cfg *config.Config, producer sarama.SyncProducer) *Router {
	r := &Router{
		hc: handlers.HealthCheckController{},
		mess: handlers.MessageController{Cfg: cfg, Producer: producer},
	}
	if cfg == nil {
		log.Fatalf("[Router Init]:Received nil configuration")
	}

	if producer == nil {
		log.Fatalf("[Router Init]:Received nil sarama.SyncProducer producer")
	}
	r.initHeathCheckRoutes(e)
	r.initMessageRoutes(e)

	return r
}

func (r *Router) initHeathCheckRoutes(e *echo.Echo) {
	e.GET("api/v1/health-check", r.hc.HealthCheck)
}

func (r *Router) initMessageRoutes(e *echo.Echo) {
	e.POST("api/v1/send-message", r.mess.SendMessage)
}