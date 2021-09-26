package cfg

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// C store global config
var C Config

type Config struct {
	APP ApplicationConfig `mapstructure:"application"`
	DB  map[string]MySqlConfig       `mapstructure:"database"`
}

func Load(envFile string)  {
	LoadEnv(envFile)

	if E.Drive == DriveFile {
		localFromFile()
	} else if E.Drive == DriveApollo {
		localFromApollo()
	}
}

//todo to know more about viper
func localFromFile() {
	appViper := viper.New()
	appViper.SetConfigName(E.File.Name)
	appViper.AddConfigPath(E.File.Path)
	appViper.AddConfigPath("./web/dist/cfg")

	if err := appViper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("fail to load application configs")
	}

	if err := appViper.Unmarshal(&C); err != nil{
		log.Fatal().Err(err).Msg("unable to decode application configs")
	}
}

func localFromApollo() {
	//app := lunar.New()

}