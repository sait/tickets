package router

import (
	"fmt"
	"net/http"

	"github.com/Flix14/soporte-tickets/api/controllers"
	"github.com/gin-gonic/gin"
)

//InitializeRoutes d
func InitializeRoutes(r *gin.Engine) {
	r.Use(CORSMiddleware())
	//Usuario routes
	r.GET("/usuarios", isAuthorized(controllers.GetUsuarios))
	r.GET("/usuarios/:id", isAuthorized(controllers.GetUsuario))
	r.POST("/usuarios", controllers.CreateUsuario)
	r.PUT("/usuarios/:id", isAuthorized(controllers.UpdateUsuario))
	r.DELETE("/usuarios/:id", isAuthorized(controllers.DeleteUsuario))

	//Agente routes
	r.GET("/agentes", isAuthorized(controllers.GetAgentes))
	r.GET("/agentes/:id", isAuthorized(controllers.GetAgente))
	r.POST("/agentes", isAuthorized(controllers.CreateAgente))
	r.PUT("/agentes/:id", isAuthorized(controllers.UpdateAgente))
	r.DELETE("/agentes/:id", isAuthorized(controllers.DeleteAgente))

	//Ticket routes
	r.GET("/tickets", isAuthorized(controllers.GetTickets))
	r.GET("/tickets/:id", isAuthorized(controllers.GetTicket))
	r.GET("/usuarios/:id/tickets", isAuthorized(controllers.GetTicketsByUsuario))
	r.POST("/tickets", controllers.CreateTicket)
	r.PATCH("/tickets/:id", isAuthorized(controllers.ChangeEstadoTicket))
	r.PUT("/tickets/:id", isAuthorized(controllers.ChangeAgenteTicket))

	//Mensaje routes
	r.GET("/tickets/:id/mensajes", isAuthorized(controllers.GetMsgsByTicket))
	r.POST("/tickets/:id/mensajes", isAuthorized(controllers.CreateMsg))

	//Auth
	r.GET("/auth/agentes", GetAuthKeyAgente)
	r.GET("/auth/usuarios", GetAuthKeyUsuario)

	//Test
	r.GET("/read_cookie", func(context *gin.Context) {
		val, err := context.Cookie("c1")
		if err != nil {
			fmt.Println(err)
		}
		context.String(200, "Cookie:%s", val)

	})
	r.GET("/write_cookie", func(context *gin.Context) {
		context.SetCookie("c1", "Shimin Li", 10, "/", "localhost", http.SameSiteNoneMode, false, true)
	})
	r.GET("/clear_cookie", func(context *gin.Context) {
		context.SetCookie("c1", "Shimin Li", -1, "/", "localhost", http.SameSiteNoneMode, false, true)
	})

}

//CORSMiddleware c
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Token, Email, Client")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Client, Email")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
