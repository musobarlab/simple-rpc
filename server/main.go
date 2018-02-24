package main

import (
	"fmt"

	httpserver "github.com/wuriyanto48/simple-rpc/server/servers"
	"github.com/wuriyanto48/simple-rpc/service"
)

func main() {

	fmt.Println("RPC CALCULATOR SERVER")

	calculator := service.NewCalculator("Calculator")

	s := httpserver.NewHTTPRPC(calculator)

	s.Run()
}
