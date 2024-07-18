package connector

import (
	"bytes"
	"github.com/trangnkp/debezium/internal/config"
	"log"
	"net/http"
	"os"
)

type Connector struct {
	cfg *config.ConnectorConfig
}

func New(cfg *config.ConnectorConfig) *Connector {
	return &Connector{
		cfg: cfg,
	}
}

func (c *Connector) Check() error {
	response, err := http.Get(c.cfg.Host + c.cfg.Name)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		plan, err := os.ReadFile(c.cfg.PlanFile)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = http.Post(c.cfg.Host, "application/json", bytes.NewBuffer(plan))
		if err != nil {
			return err
		}
	}

	return nil
}
