package httpHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "online")
	})

	clientes := r.Group("/clientes/:id")
	{
		clientes.POST("/transacoes", makeTransaction)

		clientes.GET("/extrato", getAccountStatement)
	}
	return r
}
