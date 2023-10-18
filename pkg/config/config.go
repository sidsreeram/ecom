package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	DBURL            string `mapstructure:"DB_URL"`
	TWILIOACCOUNTSID string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TWILIOAUTHTOKEN  string `mapstructure:"TWILIO_AUTHTOKEN"`
	TWILIOSERVICESID string `mapstructure:"TWILIO_SERVICES_ID"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("ECOMMERCE_PROJECT/cmd/.env")
	viper.SetConfigFile("db.env")
	viper.ReadInConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// Ensure that the DBURL field is also set
	if config.DBURL == "" {
		return config, fmt.Errorf("DB_URL is required in the configuration")
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
