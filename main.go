/*
hasaki-quant trade module
hasaki-quant server center 总调度程序
从配置文件获取到用户配置，然后启动其他模块
流程:
1,先从配置文件config.json中获取到用户的key，密钥和指定获取数据接口的参数
2,启动系统的日志中心               logManager.go
3,启动数据模块                     dataManager.go
4,启动行情网关                     gateway.go
5,启动爬虫模块                     crawlCenter.go
6,启动websocket服务并接收用户的请求websocket.go
last Code At : 2020-2-13
*/
package main

import "sync"

// globals
var go_sync sync.WaitGroup   // 全局线程协程管理

func main(){
	//websocketMain()
	go startRestful()
}