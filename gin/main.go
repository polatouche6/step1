package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//test1()
	gin.SetMode(gin.DebugMode)
	//if global.GlobalConfig.Env == "pro" {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	router := gin.New()

	// 根据code获取一个节点
	router.GET("/section", helloworld)
	//server := &http.Server{
	//	Addr:           "9001",
	//	Handler:        router,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()
	router.Run(":9001")
}

type T struct{}

func (t *T) ServeHTTP(http.ResponseWriter, *http.Request) {
		//c := engine.pool.Get().(*Context)
		//c.writermem.reset(w)
		//c.Request = req
		//c.reset()
		//
		//engine.handleHTTPRequest(c)
		//
		//engine.pool.Put(c)
}

func test1() {
	t := &T{}
	//server := &http.Server{
	//	Addr:           "9001",
	//	Handler:        t,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//server.ListenAndServe()

	http.HandleFunc("/ssss", t.ServeHTTP)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}

func helloworld(ctx *gin.Context) {
	str := fmt.Sprintf("hello world!!")
	ctx.JSON(200, str)
}
