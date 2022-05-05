package main

import (
	"context"
	"fmt"

	proto "grpc/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type message struct {
	A int64
	B int64
}

type ProxyRouter struct{}

func NewProxyRouter(g *gin.Engine, p ProxyRouter) {
	g.POST("/proxy/sum.SumService/Sum", p.connectToSumService)
	g.Run(":8881")
}

func (p ProxyRouter) connectToSumService(g *gin.Context) {
	info := message{}
	g.BindJSON(&info)
	a := info.A
	b := info.B
	fmt.Println(a, b)
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := proto.NewSumServiceClient(conn)
	result, err := client.Sum(context.Background(), &proto.SumRequest{A: 2, B: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func main() {
	g := gin.Default()
	p := ProxyRouter{}
	NewProxyRouter(g, p)
}
