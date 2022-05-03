package main

import (
	"truphone/http"
	"truphone/model"
)

type Config struct {
	Address string
	Devices model.Devices
}

type Service struct{ Config }

func NewService(c Config) (*Service, error) {
	if c.Devices == nil {
		c.Devices = model.NewDevices()
	}

	if c.Address == "" {
		c.Address = ":8080"
	}

	return &Service{c}, nil
}

func (s *Service) Run() error {
	return http.New(s.Address, s.Devices)
}
