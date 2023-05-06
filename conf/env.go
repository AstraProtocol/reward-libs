package conf

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SCInfo struct {
	SCAddress string `mapstructure:"TEST_SC_ADDR"`
	SCPK      string `mapstructure:"TEST_SC_PK"`
}

type RedisConfiguration struct {
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
	RedisTtl     string `mapstructure:"REDIS_TTL"`
}

type DatabaseConfiguration struct {
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

type NotificationConfiguration struct {
	EndPoint         string `mapstructure:"NOTIFY_ENDPOINT"`
	AccessToken      string `mapstructure:"ACCESS_TOKEN"`
	FrontendClaimURL string `mapstructure:"FE_CLAIM_URL"`
}

type EvmRpcEndpointConfiguration struct {
	ChainID       string `mapstructure:"CHAIN_ID"`
	EVMChainID    int    `mapstructure:"EVM_CHAIN_ID"`
	EndPoint      string `mapstructure:"EVMRPC_ENDPOINT"`
	PrefixAddress string `mapstructure:"PREFIX_ADDRESS"`
	TokenSymbol   string `mapstructure:"TOKEN_SYMBOL"`
}

type KafkaConfiguration struct {
	KafkaURL string `mapstructure:"KAFKA_URL" yaml:"kafkaUrl" toml:"kafkaUrl" xml:"kafkaUrl" json:"kafkaUrl,omitempty"`
	User     string `mapstructure:"KAFKA_USER" yaml:"user" toml:"user" xml:"user" json:"user,omitempty"`
	Password string `mapstructure:"KAFKA_PASSWORD" yaml:"password" toml:"password" xml:"password" json:"password,omitempty"`
}

type Configuration struct {
	Database     DatabaseConfiguration       `mapstructure:",squash"`
	Redis        RedisConfiguration          `mapstructure:",squash"`
	Notify       NotificationConfiguration   `mapstructure:",squash"`
	EvmRpc       EvmRpcEndpointConfiguration `mapstructure:",squash"`
	Kafka        KafkaConfiguration          `mapstructure:",squash"`
	TestSC       SCInfo                      `mapstructure:",squash"`
	AppName      string                      `mapstructure:"APP_NAME"`
	AppAddr      string                      `mapstructure:"APP_ADDR"`
	ConfigFile   string                      `mapstructure:"CONFIG_FILE"`
	DbUrl        string                      `mapstructure:"DB_URL"`
	Env          string                      `mapstructure:"ENV"`
	CouponSecret string                      `mapstructure:"COUPON_SECRET"`
}

var configuration Configuration

func init() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile(fmt.Sprintf("../%s", envFile))
		if err := viper.ReadInConfig(); err != nil {
			logrus.Printf("Error reading config file \"%s\", %v", envFile, err)
		}
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logrus.Printf("Unable to decode config into map, %v", err)
	}

}

func GetConfiguration() *Configuration {
	return &configuration
}

func GetRedisConnectionURL() string {
	return configuration.Redis.RedisAddress
}
