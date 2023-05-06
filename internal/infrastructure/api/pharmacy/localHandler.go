package pharmacy

import (
	"Pharmacy/internal/pkg/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocalHandler struct {
	usecase usecase.LocalUseCase
}

func NewLocalHandler(usecase usecase.LocalUseCase) *LocalHandler {
	return &LocalHandler{usecase: usecase}
}

func (h *LocalHandler) GetLocales(c *gin.Context) {
	//order := &entity.{}
	format := c.Request.Header.Get("format")
	res, err := h.usecase.GetLocales(format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusNotFound, "No se encontraron datos ")
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (h *LocalHandler) GetLocalesByComuna(c *gin.Context) {
	comuna := c.Param("comuna")

	format := c.Request.Header.Get("format")

	locales, err := h.usecase.GetLocalesByComuna(comuna, format)
	if err != nil {
		return
	}
	if len(locales) == 0 {
		c.JSON(http.StatusNotFound, "No se encontraron datos ")
	} else {
		c.JSON(http.StatusCreated, locales)
	}

}
