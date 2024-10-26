package routes

import (
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, mod *modules.Modules) {

	beef := r.Group("/beef")
	{
		beef.GET("/summary", mod.Beef.Ctl.Summary)
	}
}
