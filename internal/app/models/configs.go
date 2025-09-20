package models

import (
	"golang.org/x/oauth2"
	"time"
)

type Configs struct {
	LogParams       LogParams       `json:"log_params"`
	AppParams       AppParams       `json:"app_params"`
	PostgresParams  PostgresParams  `json:"postgres_params"`
	RedisParams     RedisParams     `json:"redis_params"`
	KafkaParams     KafkaParams     `json:"kafka_params"`
	ProvidersParams ProvidersConfig `json:"providers"`
	Cors            Cors            `json:"cors"`
	Clients         ClientsConfig   `json:"clients"`
	AuthParams      AuthParams      `json:"auth_params"`
	OAuth2          oauth2.Config   `json:"oauth2"`
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`
	LogInfo          string `json:"log_info"`
	LogError         string `json:"log_error"`
	LogWarn          string `json:"log_warn"`
	LogDebug         string `json:"log_debug"`
	MaxSizeMegabytes int    `json:"max_size_megabytes"`
	MaxBackups       int    `json:"max_backups"`
	MaxAge           int    `json:"max_age"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"local_time"`
}

type AppParams struct {
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
	AppVersion string `json:"app_version"`
	PortRun    string `json:"port_run"`
	GinMode    string `json:"gin_mode"`
	AppID      int    `json:"app_id"`
	Env        string `json:"env"`
}

type PostgresParams struct {
	User         string `json:"user"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Database     string `json:"database"`
	UserDatabase string `json:"user_database"`
	SSLMode      string `json:"sslmode"`
}

type RedisParams struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type KafkaParams struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Topic           string `json:"topic"`
	GroupID         string `json:"group_id"`
	AutoOffsetReset string `json:"auto_offset_reset"`
}

type AuthParams struct {
	JwtTtlMinutes int `json:"jwt_ttl_minutes"`
	JwtTtlHours   int `json:"jwt_ttl_hours"`
}

type Client struct {
	ClientAddress string        `json:"address"`
	Timeout       time.Duration `json:"timeout"`
	RetriesCount  int           `json:"retries_count"`
	Insecure      bool          `json:"insecure"`
}

type ClientsConfig struct {
	SSO Client `json:"sso"`
}

type GoogleProvider struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Redirect     string `json:"redirect"`
}

type ProvidersConfig struct {
	GoogleProvider GoogleProvider `json:"google_provider"`
}

type Cors struct {
	AllowOrigins     []string `json:"allow_origins"`
	AllowMethods     []string `json:"allow_methods"`
	AllowHeaders     []string `json:"allow_headers"`
	ExposeHeaders    []string `json:"expose_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
}
