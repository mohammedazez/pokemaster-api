package healthCheck

import (
	port "pokemaster-api/core/port/healthCheck"
)

type Service struct {
	repoMysql port.Repository
}

func NewService(repoMysql port.Repository) port.Service {
	return &Service{
		repoMysql: repoMysql,
	}
}

func (s *Service) Get() map[string]interface{} {
	getMysql := s.repoMysql.Get()

	result := map[string]interface{}{
		"mysql": getMysql,
	}

	return result
}
