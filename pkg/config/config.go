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
	RAZORPAYID       string `mapstructure:"RAZORPAY_ID"`
	RAZORPAYSECRET   string `mapstructure:"RAZORPAY_SECRET"`
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
	if config.RAZORPAYID ==""{
		return config,fmt.Errorf("razorpay id is required i the configuration")
	}
	if config.RAZORPAYSECRET ==""{
		return config,fmt.Errorf("razorpaysecret is required i the configuration")
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
