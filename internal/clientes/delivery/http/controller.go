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

type TransactionRequest struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

func makeTransaction(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "cliente não encontrado",
		})
		return
	}

	repo := &repositories.PostgresClientRepository{}

	accStatement, err := repo.GetAccountStatement(id)

	if err != nil && accStatement == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": err.Error(),
		})
		return
	}

	var transactionReq TransactionRequest

	if err := c.ShouldBind(&transactionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "request mal formatado",
		})
	}

	if transactionReq.Tipo == "d" && accStatement.Balance.Total+transactionReq.Valor > accStatement.Balance.Limit {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"mensagem": "limite insuficiente",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"limite": 100000,
		"saldo":  -1000,
	})
}
