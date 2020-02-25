package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sait/tickets/api/models"
)

//CreateAgente c
func CreateAgente(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	agente := models.Agente{}
	err = json.Unmarshal(body, &agente)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	agenteCreado, err := agente.SaveAgente()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, agenteCreado)
}

//GetAgentes g
func GetAgentes(c *gin.Context) {
	agente := models.Agente{}

	agentes, err := agente.FindAllAgentes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, agentes)
}

//GetAgente g
func GetAgente(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	agente := models.Agente{}
	agenteObtenido, err := agente.FindAgenteByID(uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, agenteObtenido)
}

//UpdateAgente u
func UpdateAgente(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	agente := models.Agente{}
	err = json.Unmarshal(body, &agente)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	agenteActualizado, err := agente.UpdateAgente(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, agenteActualizado)
}

//DeleteAgente d
func DeleteAgente(c *gin.Context) {
	agente := models.Agente{}

	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	agenteEliminado, err := agente.DeleteAgente(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, agenteEliminado)
}
