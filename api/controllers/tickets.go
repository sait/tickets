package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Flix14/soporte-tickets/api/models"
	"github.com/gin-gonic/gin"
)

//CreateTicket c
func CreateTicket(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ticket := models.Ticket{}
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ticketCreado, err := ticket.SaveTicket()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, ticketCreado)
}

//GetTickets g
func GetTickets(c *gin.Context) {
	ticket := models.Ticket{}

	tickets, err := ticket.FindAllTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

//GetTicket g
func GetTicket(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ticket := models.Ticket{}
	ticketObtenido, err := ticket.FindTicketByID(uint32(uid))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, ticketObtenido)
}

//GetTicketsByUsuario g
func GetTicketsByUsuario(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ticket := models.Ticket{}

	tickets, err := ticket.FindTicketsByUsuario(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

//ChangeEstadoTicket c
func ChangeEstadoTicket(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	ticket := models.Ticket{}
	err = json.Unmarshal(body, &ticket)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	uid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ticketChanged, err := ticket.ChangeEstadoTicket(uint32(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, ticketChanged)
}
