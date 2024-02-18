package main

import (
	"fmt"
	"github.com/periclescesar/rinha-2024-q1-go/configs"
	httpHandler "github.com/periclescesar/rinha-2024-q1-go/internal/clientes/delivery/http"
)

func main() {
	configs.InitConfigs(".env")
	fmt.Println("Init API...")

	r := httpHandler.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", configs.Configs.ApiConf.Port))
	if err != nil {
		panic("Server error")
	}
}
