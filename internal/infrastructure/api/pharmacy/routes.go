package pharmacy

import (
	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine, r *LocalHandler) {

	c.GET("/GetLocales", r.GetLocales)
	c.GET("/GetLocales/:comuna", r.GetLocalesByComuna)
}
