package usecase

import (
	"saints-api/model"
	"saints-api/repository"
)

type SaintUsecase struct {
	repository repository.SaintRepository
}

func NewSaintUseCase(repo repository.SaintRepository) SaintUsecase {
	return SaintUsecase{}
}

func (pu *SaintUsecase) GetSaints() ([]model.Saint, error) {
	return []model.Saint{}, nil
}
