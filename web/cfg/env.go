package cfg

import (
	"ecomb-go-demo/web/pkg/localtime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

type EnvConfig struct {
	Env      string       `mapstructure:"env"`
	Mode     string       `mapstructure:"mode"`
	Timezone string       `mapstructure:"timezone"`
	I18n     string       `mapstructure:"i18n"`
	Locales  []string     `mapstructure:"locales"`
	Drive    string       `mapstructure:"drive"`
	File     FileConfig   `mapstructure:"file"`
	Apollo   ApolloConfig `mapstructure:"apollo"`
	Port     int          `mapstructure:"port"`
}

type FileConfig struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

type ApolloConfig struct {
	AppID         string        `mapstructure:"app_id"`
	Cluster       string        `mapstructure:"cluster"`
	Namespace     string        `mapstructure:"namespace"`
	Server        string        `mapstructure:"server"`
	WatchInterval time.Duration `mapstructure:"watch_interval"`
}

const (
	DriveFile = "file"
	DriveApollo = "apollo"
)

var E EnvConfig

func LoadEnv(envFile string) {
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("useing config file: %s", viper.ConfigFileUsed())
	} else {
		log.Fatal().Err(err).Msg("fail to load env configs")
	}

	if err := viper.Unmarshal(&E); err != nil {
		log.Fatal().Err(err).Msg("unable to decode env configs")
	}

	log.Debug().Msgf("%v", viper.AllSettings())

	localtime.SetLocation(E.Timezone)

	LoadI18n(E.I18n, E.Locales)
}
