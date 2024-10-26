package beef

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BeefController struct {
	beefSvc *BeefService
}

func newController(beefSvc *BeefService) *BeefController {
	return &BeefController{
		beefSvc: beefSvc,
	}
}

func (ctl *BeefController) Summary(c *gin.Context) {

	data, err := ctl.beefSvc.Summary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
