package main

import (
	"context"
	"net/http"

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
	// 起代理會打不到～地址也要改一下喔喔喔喔喔
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	defer conn.Close()
	client := proto.NewSumServiceClient(conn)
	result, err := client.Sum(context.Background(), &proto.SumRequest{A: a, B: b})
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

func main() {
	g := gin.Default()
	p := ProxyRouter{}
	NewProxyRouter(g, p)
}
