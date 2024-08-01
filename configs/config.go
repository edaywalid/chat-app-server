package configs

import "github.com/spf13/viper"

type Config struct {
	DBUrl     string `mapstructure:"DB_URL"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config = &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
