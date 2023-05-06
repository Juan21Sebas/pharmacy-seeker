package repositories

import (
	"Pharmacy/internal/pkg/entity"
	"Pharmacy/internal/pkg/ports"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// En el repositorio, se implementa la interfaz LocalRepository y se hace la petición HTTP a la API
type localRepository struct {
	url string
}

func NewLocalRepository(url string) ports.LocalRepository {
	return &localRepository{url: url}
}

func (r *localRepository) GetLocales(format string) ([]*entity.Local, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locales []*entity.Local

	resbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error 4: %s", err)
		return nil, err
	}
	switch format {
	case "json":
		resbody = bytes.TrimPrefix(resbody, []byte("\xef\xbb\xbf"))
		if err := json.Unmarshal(resbody, &locales); err != nil {
			return nil, err
		}
		return locales, nil
	case "xml":
		resp, err := r.GetLocalesXML()
		return resp, err
	default:
		return nil, fmt.Errorf("formato '%s' no válido", format)
	}
}

func (r *localRepository) GetLocalesXML() ([]*entity.Local, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return nil, err
	}

	// Reemplazar el carácter "&" por "&amp;"
	escapedBuf := bytes.ReplaceAll(b.Bytes(), []byte("&"), []byte("&amp;"))

	var locales []*entity.Local
	if err := xml.Unmarshal(escapedBuf, &locales); err != nil {
		return nil, err
	}

	return locales, nil
}

func (r *localRepository) GetLocalesByComuna(comuna string, format string) ([]*entity.Local, error) {
	locales, err := r.GetLocales(format)
	if err != nil {
		return nil, err
	}

	var filteredLocales []*entity.Local
	for _, local := range locales {
		if local.ComunaNombre == comuna {
			filteredLocales = append(filteredLocales, local)
		}
	}

	return filteredLocales, nil

}
