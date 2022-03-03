package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/roberto/go-api-rest/controllers"
	"github.com/roberto/go-api-rest/database"
	"github.com/roberto/go-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaCapituraMock() {
	capitura := models.Capitura{Nome: "Capitura Teste"}
	database.DB.Create(&capitura)
	ID = int(capitura.ID)
}

func DeletaCapituraMock() {
	var capitura models.Capitura
	database.DB.Delete(&capitura, ID)
}

func TestListaTodasCapiturasHanlder(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaCapituraMock()
	defer DeletaCapituraMock()
	r := SetupDasRotasDeTeste()
	r.GET("/capituras", controllers.TodasCapituras)
	req, _ := http.NewRequest("GET", "/capituras", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaCapituraPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaCapituraMock()
	defer DeletaCapituraMock()
	r := SetupDasRotasDeTeste()
	r.GET("/capituras/:id", controllers.BuscarCapituraPorID)
	pathDaBusca := "/capituras/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var capituraMock models.Capitura
	json.Unmarshal(resposta.Body.Bytes(), &capituraMock)
	assert.Equal(t, "Capitura Teste", capituraMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaCapituraHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaCapituraMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/capituras/:id", controllers.DeletarCapitura)
	pathDeBusca := "/capituras/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmCapituraHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaCapituraMock()
	defer DeletaCapituraMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/capituras/:id", controllers.EditarCapitura)
	capitura := models.Capitura{Nome: "Nome do Aluno Teste"}
	valorJson, _ := json.Marshal(capitura)
	pathParaEditar := "/capituras/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var capituraMockAtualizada models.Capitura
	json.Unmarshal(resposta.Body.Bytes(), &capituraMockAtualizada)
	assert.Equal(t, "Nome do Aluno Teste", capituraMockAtualizada.Nome)
}
