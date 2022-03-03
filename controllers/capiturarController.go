package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roberto/go-api-rest/database"
	"github.com/roberto/go-api-rest/models"
)

func TodasCapituras(c *gin.Context) {
	var capituras []models.Capitura
	database.DB.Find(&capituras)
	c.JSON(200, capituras)
}

func BuscarCapituraPorID(c *gin.Context) {
	var capitura models.Capitura
	id := c.Params.ByName("id")
	database.DB.First(&capitura, id)
	if capitura.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Capitura não encontrada"})
		return
	}
	c.JSON(http.StatusOK, capitura)
}

func CriarNovaCapitura(c *gin.Context) {
	var capitura models.Capitura
	if err := c.ShouldBindJSON(&capitura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeCapitura(&capitura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Create(&capitura)
	c.JSON(http.StatusOK, capitura)
}

func DeletarCapitura(c *gin.Context) {
	var capitura models.Capitura
	id := c.Params.ByName("id")
	database.DB.Delete(&capitura, id)
	c.JSON(http.StatusOK, gin.H{"data": "Capitura deletada com sucesso"})
}

func EditarCapitura(c *gin.Context) {
	var capitura models.Capitura
	id := c.Params.ByName("id")
	database.DB.First(&capitura, id)
	if err := c.ShouldBindJSON(&capitura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaDadosDeCapitura(&capitura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Save(&capitura)
	c.JSON(http.StatusOK, capitura)
}
