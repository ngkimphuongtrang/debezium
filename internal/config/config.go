package config

import "context"

type Config struct {
	Connector *ConnectorConfig `yaml:"connector" mapstructure:"connector"`
	Kafka     *KafkaConfig     `yaml:"kafka" mapstructure:"kafka"`
}

type ConnectorConfig struct {
	Host        string `yaml:"host" mapstructure:"host"`
	Name        string `yaml:"name" mapstructure:"name"`
	PlanFile    string `yaml:"plan_file" mapstructure:"plan_file"`
	TopicPrefix string `yaml:"topic_prefix" mapstructure:"topic_prefix"`
	Schema      string `yaml:"schema" mapstructure:"schema"`
	DbName      string `yaml:"db_name" mapstructure:"db_name"`
}

type KafkaConfig struct {
	Host string `yaml:"host" mapstructure:"host"`
}

func New(ctx context.Context) *Config {
	cfg, err := LoadConfig(ctx, "")
	if err != nil {
		panic(err)
	}
	return cfg
}
