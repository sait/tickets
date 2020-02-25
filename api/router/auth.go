package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sait/tickets/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte("soporteticketssait1524863970")

//GetAuthKeyUsuario g
func GetAuthKeyUsuario(c *gin.Context) {
	validToken, err := GenerateJWTUsuario(c.Query("email"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	usuario := models.Usuario{}
	val, err := usuario.ValidateCredentials(c.Query("email"), c.Query("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if val != nil {
		c.Header("Token", validToken)
		c.Next()
		c.JSON(http.StatusOK, val)
		return
	}
	c.JSON(http.StatusNotFound, "")
}

//GetAuthKeyAgente g
func GetAuthKeyAgente(c *gin.Context) {
	validToken, err := GenerateJWTAgente(c.Query("email"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	agente := models.Agente{}
	val, err := agente.ValidateCredentials(c.Query("email"), c.Query("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if val != nil {
		c.Header("Token", validToken)
		c.Next()
		c.JSON(http.StatusOK, val)
		return
	}
	c.JSON(http.StatusNotFound, "")
}

//GenerateJWTUsuario g
func GenerateJWTUsuario(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["client"] = "Usuario"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

//GenerateJWTAgente g
func GenerateJWTAgente(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["client"] = "Agente"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func isAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header["Token"] != nil {

			token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}
			if token.Valid {
				client, _ := token.Claims.(jwt.MapClaims)["client"].(string)
				email := token.Claims.(jwt.MapClaims)["email"].(string)
				c.Header("client", client)
				c.Header("email", email)
				c.Next()
				endpoint(c)
			} else {
				c.JSON(http.StatusUnauthorized, "No autorizado")
			}

		} else {
			c.JSON(http.StatusUnauthorized, "No autorizado")
		}
	}
}
