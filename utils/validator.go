package utils

import (
	"errors"
	"net"
	"strconv"
)

func ValidateIP(ip string) error {
	if net.ParseIP(ip) == nil {
		return errors.New("invalid IP address")
	}
	return nil
}

func ValidatePort(port string) error {
	p, err := strconv.Atoi(port)
	if err != nil || p < 1 || p > 65535 {
		return errors.New("invalid port number")
	}
	return nil
}
