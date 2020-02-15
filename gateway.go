/*
hasaki-quant 数据中台连接交易所的网关
所有交易所的连接都通过gateway去调度
网关包括请求所有交易所的价格行情，所有接口的数据最后在gateway中进行抽象

网关还包含统一发单的功能，如果交易所被墙，那么在外网部署另一个hasaki-server，然后国内的服务中台的网关
连接墙外的数据中台进行数据订阅
所以hasaki-quant data server center做一个分支，若是在墙外启动，则在服务中台的配置文件中进行设置，只执行部分功能

网关接受下单需要的参数: market,symbol,price,volume 订阅行情需要的参数 : market,symbol,frequency
*/
package main

// globals var
type Gateway struct{

}
type gatewayInterface interface{
	setting(market string,symbol string,frequency string,apiKey string,secret string)
	sendOrder(market string,symbol string,price string,volume string)
}

// 设置订阅数据
func (g *Gateway) setting(market string,symbol string,frequency string,apiKey string,secret string){
	// params market : exchange name like gateio
	// params symbol : exchange`s symbol like btc
	// params frequency : K line `s frequency like 1min,1hour,1day
	// params apiKey : user `s api key it create in exchange
	// params secret : user `s secret key
	switch market {
	case "gateio":
		break
	case "huobi":
		break
	default:
		print("gateway setting switch default")
	}

}

// 接收外部的数据并下单下单
func (g *Gateway) sendOrder(market string,symbol string,price string,volume string){
	switch market {
	case "gateio":
		break
	case "huobi":
		break
	default:
		print("gateway send order switch default")
	}
}