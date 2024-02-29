package httpHandler

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func Test_getAccountStatement(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		id   string
	}{
		{
			name: "client not found",
			id:   "6",
		},
		{
			name: "client not found",
			id:   "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Params = []gin.Param{gin.Param{Key: "id", Value: tt.id}}
			getAccountStatement(c)
		})
	}
}

func Test_makeTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		id   string
	}{
		{
			name: "client not found",
			id:   "6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Params = []gin.Param{gin.Param{Key: "id", Value: tt.id}}
			makeTransaction(c)
		})
	}
}
