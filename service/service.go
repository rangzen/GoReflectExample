package service

import (
	"context"
	"github.com/rangzen/GoReflectExample/repository"
)

type Service struct {
	Repository repository.Repository
}

func (s *Service) EtlOrder(ctx context.Context) string {
	dataA := s.Repository.DataA(ctx)
	dataB := s.Repository.DataB(ctx)
	dataC := s.Repository.DataC(ctx)

	return dataA + dataB + dataC
}

func (s *Service) EtlReverse(ctx context.Context) string {
	dataA := s.Repository.DataA(ctx)
	dataB := s.Repository.DataB(ctx)
	dataC := s.Repository.DataC(ctx)

	return dataC + dataB + dataA
}
