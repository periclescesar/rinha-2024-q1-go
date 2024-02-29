package httpHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/periclescesar/rinha-2024-q1-go/internal/clientes/repositories"
	"net/http"
	"strconv"
)

func getAccountStatement(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "cliente não encontrado",
		})
		return
	}

	repo := &repositories.PostgresClientRepository{}

	accStatement, err := repo.GetAccountStatement(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, accStatement)
}

func makeTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "cliente não encontrado",
		})
		return
	}

	if id == 6 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "cliente não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"limite": 100000,
		"saldo":  -9098,
	})
}
