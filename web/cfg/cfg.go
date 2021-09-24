package cfg

// C store global config
var C Config

type Config struct {
	APP ApplicationConfig `mapstructure:"application"`
	DB  MySqlConfig       `mapstructure:"database"`
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

}

func localFromApollo() {

}