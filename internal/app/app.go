package app

import (
	"github.com/trangnkp/debezium/internal/config"
	"github.com/trangnkp/debezium/internal/connector"
	"github.com/trangnkp/debezium/internal/consumer"
	"log"
	"strings"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (app *App) Run() {
	conn := connector.New(app.cfg.Connector)
	err := conn.Check()
	if err != nil {
		panic(err)
	}

	cons, err := consumer.New(app.cfg)
	if err != nil {
		panic(err)
	}

	err = cons.SubscribeTopic(app.getTopic(app.cfg.Connector))
	if err != nil {
		panic(err)
	}

	log.Println("Listening ...")
	for {
		msg, err := cons.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Message:", msg.TopicPartition, string(msg.Value))
	}
}

func (app *App) getTopic(cfg *config.ConnectorConfig) string {
	return strings.Join([]string{cfg.TopicPrefix, cfg.Schema, cfg.DbName}, ".")
}
