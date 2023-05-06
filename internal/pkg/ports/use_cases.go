package ports

import "Pharmacy/internal/pkg/entity"

type LocalUseCase interface {
	GetLocales(format string) ([]*entity.Local, error)
	GetLocalesByComuna(comuna string, format string) ([]*entity.Local, error)
}
