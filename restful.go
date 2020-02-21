/*
hasaki server restful 
return the news data and quote data to user
这个restful请求会全部返回数据库中所有的数据给用户，
用户在启动服务的时候，已经设定好了，数据量的保存模式
保存所有数据或者只保留一天的数据，如果是保存一天的数据应该就不需要数据库了
因为一天的数据的量并不大
*/
package main

import "github.com/gin-gonic/gin"

func newsServer(context *gin.Context){
	// function : 返回新闻舆情的服务
	context.JSON(200,gin.H{
		"code":200,
		"success":true,
		"news":"hasaki",
	})
}

func quoteServer(context *gin.Context){
	// function : 返回历史历史行情的服务
}

func startRestful(){
	// Engin指针
    router := gin.Default()

	router.GET("/newsServer", newsServer)
	router.GET("/quoteServer",quoteServer)
    // 指定地址和端口号
	router.Run("0.0.0.0:8888")                    // 如果是云服务改成0.0.0.0:8888
}