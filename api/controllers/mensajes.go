package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sait/tickets/api/models"
)

//CreateMsg c
func CreateMsg(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	mensaje := models.Mensaje{}
	err = json.Unmarshal(body, &mensaje)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	mensajeCreado, err := mensaje.SaveMsg()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, mensajeCreado)
}

//GetMsgsByTicket g
func GetMsgsByTicket(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	mensaje := models.Mensaje{}

	mensajes, err := mensaje.FindAllMsgByTicket(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, mensajes)
}
