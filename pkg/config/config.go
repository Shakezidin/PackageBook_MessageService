package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	EMAIL    string
	PASSWORD string
	PORT     string
}

func Configuration() (*Conf, error) {
	godotenv.Load("../../.env")
	conf := &Conf{
		EMAIL:    os.Getenv("EMAIL"),
		PASSWORD: os.Getenv("PASSWORD"),
		PORT:     os.Getenv("PORT"),
	}

	return conf, nil
}
