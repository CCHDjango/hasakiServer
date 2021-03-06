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

func govCrollNews(context *gin.Context){
	// function : 返回数据库全部新闻舆情
	context.JSON(200,gin.H{
		"code":200,
		"success":true,
		"news":"hasaki",
	})
}

func gateioQuote(context *gin.Context){
	// function : 返回数据库全部的历史历史行情
}

func xinlangNews(context *gin.Context){
	// function : 返回新浪国际新闻滚动列表
}

func goldQuote(context *gin.Context){
	// function : 黄金行情数据
}

func DJI(context *gin.Context){
	// function : 美国道指数据
}

func startRestful(){
	// Engin指针
    router := gin.Default()

	router.GET("/govCrollNews", govCrollNews)
	router.GET("/gateioQuote",gateioQuote)
	router.GET("/goldQuote",goldQuote)
	router.GET("/xinlangNews",xinlangNews)
	router.GET("/DJI",DJI)
    // 指定地址和端口号
	router.Run("0.0.0.0:8888")                    // 如果是云服务改成0.0.0.0:8888
}