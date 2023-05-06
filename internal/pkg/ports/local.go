package ports

import "Pharmacy/internal/pkg/entity"

type LocalRepository interface {
	GetLocales(format string) ([]*entity.Local, error)
	GetLocalesByComuna(comuna string, format string) ([]*entity.Local, error)
}
