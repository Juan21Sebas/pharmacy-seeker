package usecase

import (
	"Pharmacy/internal/pkg/entity"
	"Pharmacy/internal/pkg/ports"
)

type LocalUseCase struct {
	repo ports.LocalUseCase
}

func NewLocalUseCase(repo ports.LocalRepository) *LocalUseCase {
	return &LocalUseCase{repo: repo}
}

func (uc *LocalUseCase) GetLocales(format string) ([]*entity.Local, error) {
	return uc.repo.GetLocales(format)
}

func (uc *LocalUseCase) GetLocalesByComuna(comuna string, format string) ([]*entity.Local, error) {
	return uc.repo.GetLocalesByComuna(comuna, format)
}
