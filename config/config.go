package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		General			`yaml:"general"`
		Addition  		`yaml:"addition"`
		Subtraction 	`yaml:"subtraction"`
		Multiplication  `yaml:"multiplication"`
		Division		`yaml:"division"`
	}

	General struct {
		TimerSecond	  	uint8	`yaml:"timer_second" default:"0"`
		EnableScores	bool 	`yaml:"enable_scores" default:"true"`
		EnablePenalty 	bool 	`yaml:"enable_penalty" default:"false"`
	}


	Addition struct {
		Enable  		bool	`yaml:"enable" default:"true"`
		LeftDigits 		uint8 	`yaml:"left_digits" default:"3"`
		RightDigits 	uint8 	`yaml:"right_digits" default:"2"`
	}

	Subtraction struct {
		Enable  		bool	`yaml:"enable" default:"true"`
		NegativeResult	bool	`yaml:"negative_result" default:"false"`
		LeftDigits 		uint8 	`yaml:"left_digits" default:"3"`
		RightDigits 	uint8 	`yaml:"right_digits" default:"2"`
	}

	Multiplication struct {
		Enable  		bool	`yaml:"enable" default:"false"`
		LeftDigits 		uint8 	`yaml:"left_digits" default:"3"`
		RightDigits 	uint8 	`yaml:"right_digits" default:"2"`
	}

	Division struct {
		Enable  		bool	`yaml:"enable" default:"false"`
		DividendDigits 	uint8 	`yaml:"dividend_digits" default:"3"`
		DivisorDigits 	uint8 	`yaml:"divisor_digits" default:"2"`
	}
)

func NewConfig(path string) (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
