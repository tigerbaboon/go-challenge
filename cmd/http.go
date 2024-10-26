package cmd

import (
	"app/app/modules"
	"app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func httpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "http",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()

			r.Use(cors.New(cors.Config{
				AllowAllOrigins:  true,
				AllowMethods:     []string{"*"},
				AllowHeaders:     []string{"*"},
				AllowCredentials: true,
			}))

			mod := modules.Get()

			routes.Router(r, mod)

			r.Run(":8080")
		},
	}
	return cmd
}
