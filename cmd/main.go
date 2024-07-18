package main

import (
	"context"
	"github.com/trangnkp/debezium/internal/app"
	"github.com/trangnkp/debezium/internal/config"
)

func main() {
	cfg := config.New(context.Background())
	myApp := app.New(cfg)
	myApp.Run()
}
