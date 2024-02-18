package httpHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	clientes := r.Group("/clientes/:id")
	{
		clientes.POST("/transacoes", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(404, gin.H{
					"message": "cliente não encontrado",
				})
				return
			}

			if id == 6 {
				c.JSON(404, gin.H{
					"message": "cliente não encontrado",
				})
				return
			}

			c.JSON(200, gin.H{
				"limite": 100000,
				"saldo":  -9098,
			})
		})

		clientes.GET("/extrato", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				c.JSON(404, gin.H{
					"message": "cliente não encontrado",
				})
				return
			}

			if id == 6 {
				c.JSON(404, gin.H{
					"message": "cliente não encontrado",
				})
				return
			}

			c.JSON(200, gin.H{
				"saldo": gin.H{
					"total":        -9098,
					"data_extrato": "2024-01-17T02:34:41.217753Z",
					"limite":       100000,
				},
				"ultimas_transacoes": gin.H{
					"valor":        1000,
					"tipo":         "c",
					"descricao":    "descricao",
					"realizada_em": "2024-01-17T02:34:38.543030Z",
				},
			})
		})
	}
	return r
}
