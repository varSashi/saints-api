package usecase

import (
	"saints-api/model"
	"saints-api/repository"
)

type SaintUsecase struct {
	repository repository.SaintRepository
}

func NewSaintUseCase(repo repository.SaintRepository) SaintUsecase {
	return SaintUsecase{
		repository: repo,
	}
}

func (pu *SaintUsecase) GetSaints() ([]model.Saint, error) {
	return pu.repository.GetSaints()
}

func (pu *SaintUsecase) CreateSaint(saint model.Saint) (model.Saint, error) {

	saintId, err := pu.repository.CreateSaint(saint)
	if err != nil {
		return model.Saint{}, err
	}

	saint.ID = saintId

	return saint, nil
}

func (pu *SaintUsecase) GetSaintById(id_saint int) (*model.Saint, error) {

	saint, err := pu.repository.GetSaintById(id_saint)
	if(err != nil){
		return nil, err
	}

	return saint, nil
}