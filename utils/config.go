package utils

import (
	"log"
	"os"
	"strings"
	"syscall"
	"time"
)

// APIConfig ...
type APIConfig struct {
	Env         string // "development", "production"
	HeaderValue struct {
		ClientID string
	}
	EnableTracing  bool
	DisableLogging bool
}

// ConfigDB ...
type ConfigDB struct {
	Postgres struct {
		DBMS     string
		Username string
		Password string
		DBName   string
	}
}

// NewAPIConfig ...
func NewAPIConfig() *APIConfig {
	config := new(APIConfig)
	config.HeaderValue.ClientID = os.Getenv("SBCNTR_CLIENT_ID_HEADER")

	// 環境変数[SBCNTR_ENABLE_TRACING]を見てトレースを有効にする。対応しているTracingはAWS_XRAYのみ。
	enableKey := os.Getenv("SBCNTR_ENABLE_TRACING")
	if strings.ToLower(enableKey) == "true" || enableKey == "1" {
		config.EnableTracing = true
	} else {
		config.EnableTracing = false
	}

	// 環境変数[SBCNTR_DISABLE_LOGGING]を見てリクエストログを無効にする
	disableLoggingKey := os.Getenv("SBCNTR_DISABLE_LOGGING")
	if strings.ToLower(disableLoggingKey) == "true" || disableLoggingKey == "1" {
		config.DisableLogging = true
	} else {
		config.DisableLogging = false
	}

	config.Env = os.Getenv("APP_ENV")
	if config.Env == "" {
		config.Env = "development"
	}

	// 環境変数[SBCNTR_ENABLE_AUTO_SIGTERM]を見て自プロセスへ SIGTERM を送出する。Cloud Run のグレースフル停止検証用。
	enableAutoSigterm := os.Getenv("SBCNTR_ENABLE_AUTO_SIGTERM")
	if strings.ToLower(enableAutoSigterm) == "true" || enableAutoSigterm == "1" {
		scheduleAutoSigterm()
	}

	return config
}

// 120秒後に自プロセスへ SIGTERM を送出する。Cloud Run のグレースフル停止検証用。
func scheduleAutoSigterm() {
	log.Println("[utils] SBCNTR_ENABLE_AUTO_SIGTERM=true: SIGTERM will be sent in 120s")
	go func() {
		time.Sleep(120 * time.Second)
		if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
			log.Printf("[utils] failed to send SIGTERM: %v", err)
		}
	}()
}

// NewConfigDB ...
func NewConfigDB() *ConfigDB {
	config := new(ConfigDB)

	config.Postgres.DBMS = "postgres"
	config.Postgres.Username = os.Getenv("DB_USERNAME")
	config.Postgres.Password = os.Getenv("DB_PASSWORD")
	config.Postgres.DBName = os.Getenv("DB_NAME")

	return config
}
