package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roberto/go-api-rest/controllers"
)

func HandleRequestCapiturar() {
	r := gin.Default()
	r.GET("/capituras", controllers.TodasCapituras)
	r.GET("/capituras/:id", controllers.BuscarCapituraPorID)
	r.POST("/capituras", controllers.CriarNovaCapitura)
	r.DELETE("/capituras/:id", controllers.DeletarCapitura)
	r.PATCH("/capituras/:id", controllers.EditarCapitura)
	r.Run(":8000")
}
