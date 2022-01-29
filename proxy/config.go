package proxy

import (
	"github.com/newPepole123/srv-yy/config"
	"github.com/newPepole123/srv-yy/repository"
)

func WebConfig() *config.WebConfig {
	return repository.WebConfig()
}
