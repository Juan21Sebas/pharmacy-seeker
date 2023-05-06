package repositories

import (
	"Pharmacy/internal/infrastructure/repositories/mocks"
	"Pharmacy/internal/pkg/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	mockRepo := new(mocks.LocalRepository)

	r := localRepository{
		url: "https://farmanet.minsal.cl/index.php/ws/getLocales",
	}

	expectedLocales := &entity.Local{
		Fecha:                      "06-05-23",
		LocalId:                    "1",
		LocalNombre:                "CRUZ VERDE ",
		ComunaNombre:               "QUILLOTA",
		LocalidadNombre:            "QUILLOTA",
		LocalDireccion:             "OHIGGINS 195, LOCAL 1",
		FuncionamientoHoraApertura: "08:30:00",
		FuncionamientoHoraCierre:   "18:30:00",
		LocalTelefono:              "+5633332269467",
		LocalLat:                   "-32.8793428949969",
		LocalLng:                   "-71.2467871500868",
		FuncionamientoDia:          "sabado",
		FkRegion:                   "6",
		FkComuna:                   "69",
		FkLocalidad:                "32",
	}
	mockRepo.On("GetLocales").Return(expectedLocales, nil)
	var formart string
	formart = "json"
	locales, err := r.GetLocales(formart)

	assert.Nil(t, err)
	assert.Equal(t, expectedLocales, locales[0])
}

func TestGetLocalesByComuna(t *testing.T) {
	mockRepo := new(mocks.LocalRepository)

	r := localRepository{
		url: "https://farmanet.minsal.cl/index.php/ws/getLocales",
	}

	expectedLocales := &entity.Local{
		Fecha:                      "06-05-23",
		LocalId:                    "1",
		LocalNombre:                "CRUZ VERDE ",
		ComunaNombre:               "QUILLOTA",
		LocalidadNombre:            "QUILLOTA",
		LocalDireccion:             "OHIGGINS 195, LOCAL 1",
		FuncionamientoHoraApertura: "08:30:00",
		FuncionamientoHoraCierre:   "18:30:00",
		LocalTelefono:              "+5633332269467",
		LocalLat:                   "-32.8793428949969",
		LocalLng:                   "-71.2467871500868",
		FuncionamientoDia:          "sabado",
		FkRegion:                   "6",
		FkComuna:                   "69",
		FkLocalidad:                "32",
	}
	mockRepo.On("GetLocales").Return(expectedLocales, nil)
	var formart string
	formart = "json"
	locales, err := r.GetLocalesByComuna("QUILLOTA", formart)

	assert.Nil(t, err)
	assert.Equal(t, expectedLocales, locales[0])
}

func TestNewLocalRepository(t *testing.T) {
	url := "https://farmanet.minsal.cl/index.php/ws/getLocales"
	repo := NewLocalRepository(url)
	if repo.(*localRepository).url != url {
		t.Errorf("Expected url %s but got %s", url, repo.(*localRepository).url)
	}
}
