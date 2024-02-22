package config

import (
	"os"
)

type Conf struct {
	EMAIL    string
	PASSWORD string
	PORT     string
}

func Configuration() (*Conf, error) {
	conf := &Conf{
		EMAIL:    os.Getenv("EMAIL"),
		PASSWORD: os.Getenv("PASSWORD"),
		PORT:     os.Getenv("PORT"),
	}

	return conf, nil
}
