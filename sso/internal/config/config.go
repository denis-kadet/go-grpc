package config

import "time"

type Config struct {
	//https://stackoverflow.com/questions/10858787/what-are-the-uses-for-struct-tags-in-go
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storagePath" env-required:"true"`
	TokenTTL    time.Duration `yaml:"tokenTTL" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"GRPC"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {

}
