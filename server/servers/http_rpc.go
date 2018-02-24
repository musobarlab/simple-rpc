package servers

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/wuriyanto48/simple-rpc/service"
)

//httprpc struct
type httprpc struct {
	serv service.Service
}

//NewHTTPRPC httprpc constructor
func NewHTTPRPC(serv service.Service) *httprpc {
	return &httprpc{serv: serv}
}

//Run , server httprpc
func (h *httprpc) Run() {

	s := rpc.NewServer()

	//register service with its name
	s.RegisterName(service.CalculatorName, h.serv)

	// show list of all services on http handler, visit localhost:9000/services
	s.HandleHTTP("/", "/services")

	// Listen for incoming tcp packets on specified port
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	//server
	http.Serve(listener, nil)
}
