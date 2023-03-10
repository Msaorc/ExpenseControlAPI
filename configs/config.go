package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var conf *config

type config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwtExperesIn  int    `mapstructure:"JWT_EXPERESIN"`
	TokenAuth     *jwtauth.JWTAuth
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfigs(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	conf.TokenAuth = jwtauth.New("HS256", []byte(conf.JwtSecret), nil)
	return conf, nil
}
