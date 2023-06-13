package service

import (
	"errors"
	"strings"
)

type ClusterConfig struct {
	Name  string
	Host  string
	Token string
}

func (c *ClusterConfig) Validate() error {

	if strings.TrimSpace(c.Name) == "" {
		return errors.New("name cluster config name empty or nil")
	}
	if strings.TrimSpace(c.Host) == "" {
		return errors.New("host cluster config host empty or nil")
	}
	if strings.TrimSpace(c.Token) == "" {
		return errors.New("token cluster config token empty or nil")
	}

	return nil
}
