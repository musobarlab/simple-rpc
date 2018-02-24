package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/wuriyanto48/simple-rpc/service"
)

type Client struct {
	c *rpc.Client
}

func (c *Client) Add(x, y float64) service.Result {
	inputs := &service.Inputs{X: x, Y: y}

	var result service.Result
	err := c.c.Call("Calculator.Add", inputs, &result)

	if err != nil {
		result = service.Result{Error: err}
	}

	return result
}

func (c *Client) Div(x, y float64) service.Result {
	inputs := &service.Inputs{X: x, Y: y}

	var result service.Result
	err := c.c.Call("Calculator.Div", inputs, &result)

	if err != nil {
		result = service.Result{Error: err}
	}

	return result
}

func (c *Client) Mul(x, y float64) service.Result {
	inputs := &service.Inputs{X: x, Y: y}

	var result service.Result
	err := c.c.Call("Calculator.Mul", inputs, &result)

	if err != nil {
		result = service.Result{Error: err}
	}

	return result
}

func (c *Client) Sub(x, y float64) service.Result {
	inputs := &service.Inputs{X: x, Y: y}

	var result service.Result
	err := c.c.Call("Calculator.Sub", inputs, &result)

	if err != nil {
		result = service.Result{Error: err}
	}

	return result
}

func main() {

	fmt.Println("Simple RPC Client")

	// connect to localhost:1234 using HTTP protocol (The port on which rpc server is listening)
	client, err := rpc.DialHTTP("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("Error Dialing:", err)
	}

	calculator := &Client{client}

	addRes := calculator.Add(10.0, 5.0)
	divRes := calculator.Div(10.0, 0)
	mulRes := calculator.Mul(10.0, 5.0)
	subRes := calculator.Sub(10.0, 5.0)

	fmt.Println(addRes.Z)
	fmt.Println(divRes)
	fmt.Println(mulRes.Z)
	fmt.Println(subRes.Z)
}
