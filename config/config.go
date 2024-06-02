package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	S3 struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"accessKeyID"`
		SecretAccessKey string `yaml:"secretAccessKey"`
		UseSSL          bool   `yaml:"useSSL"`
		ResourcesBucket string `yaml:"resourcesBucket"`
	} `yaml:"S3"`
	Database struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"dbname"`
	} `yaml:"database"`
}

const (
	configName = "config.yaml"
)

func GetConfig() (*Config, error) {
	f, err := os.Open(configName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
