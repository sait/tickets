package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Flix14/soporte-tickets/api/models"
	"github.com/gin-gonic/gin"
)

//CreateUsuario c
func CreateUsuario(c *gin.Context) {

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	usuario := models.Usuario{}
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	usuarioCreado, err := usuario.SaveUsuario()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, usuarioCreado)
}

//GetUsuarios g
func GetUsuarios(c *gin.Context) {
	usuario := models.Usuario{}
	if c.Query("email") != "" {
		GetUsuarioByEmail(c)
	} else {
		usuarios, err := usuario.FindAllUsuarios()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
		c.JSON(http.StatusOK, usuarios)
	}
}

//GetUsuario g
func GetUsuario(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	usuario := models.Usuario{}
	usuarioObtenido, err := usuario.FindUsuarioByID(uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, usuarioObtenido)
}

//GetUsuarioByEmail g
func GetUsuarioByEmail(c *gin.Context) {
	email := c.Query("email")
	usuario := models.Usuario{}
	usuarioObtenido, err := usuario.FindUsuarioByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, usuarioObtenido)
}

//UpdateUsuario u
func UpdateUsuario(c *gin.Context) {
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
	usuario := models.Usuario{}
	err = json.Unmarshal(body, &usuario)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	usuarioActualizado, err := usuario.UpdateUsuario(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, usuarioActualizado)
}

//DeleteUsuario d
func DeleteUsuario(c *gin.Context) {
	usuario := models.Usuario{}

	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	usuarioEliminado, err := usuario.DeleteUsuario(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, usuarioEliminado)
}
