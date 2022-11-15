package service

import (
	"go_api/repository"
)

type PingService interface {
	Ping() string
}

type pingService struct {
	pingRepository repository.PingRepository
}

func NewPingService(pingRepo repository.PingRepository) PingService {
	return &pingService{
		pingRepository: pingRepo,
	}
}
func (service *pingService) Ping() string {
	return "nghilc"
}
