package config

import "os"

// Config 設定情報をまとめた構造体
type Config struct {
	DatabaseURI string
}

// New Configのコンストラクター
func New() *Config {
	return &Config{
		DatabaseURI: os.Getenv("DATABASE_URI"),
	}
}
