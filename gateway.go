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
	subscribeQuote(market string,symbol string,frequency string)
	gateRestfulQuote(symbol string)
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
		gateKEY=apiKey
		gateSECRET=secret
		break
	case "huobi":
		break
	default:
		print("gateway setting switch default")
	}

	// 订阅行情
	go_sync.Add(1)
	go g.subscribeQuote(market,symbol,frequency)
	go_sync.Wait()
}

// 接收外部的数据并下单下单
func (g *Gateway) sendOrder(direction string,market string,symbol string,price string,volume string){
	switch market {
	case "gateio":
		go_sync.Add(1)
		go g.gateSendOrder(direction,symbol,price,volume)
		go_sync.Wait()
		break
	case "huobi":
		break
	default:
		print("gateway send order switch default")
	}
}

/*--------------------------控制交易所接口--------------------------*/
func (g *Gateway) subscribeQuote(market string,symbol string,frequency string){
	// function : 根据市场和品种和周期订阅行情,通过goroutine启动然后通过channel把数据传出去
	// param market : example - gateio
	// param symbol : example - btc_usdt
	// param frequency : example - 1min
	defer go_sync.Done()
	switch market {
	case "gateio":
		g.gateRestfulQuote(symbol)
		break
	case "huobi":
		break
	}
}

// gateio 交易所的接口
func (g *Gateway) gateRestfulQuote(symbol string){
	// function : 通过gateio restful接口查询行情，这里拿到的是tick的行情
	// 默认是10秒查询一次行情，拿到行情后，把行情传到dataManager
	// param symbol : 订阅的品种 example - btc_usdt
	for {
		result:=gateTicker(symbol)
		print(result)
		sleep(10)
	}

}

func (g *Gateway) gateSendOrder(direction string,symbol string,price string,volume string){
	// function : 发单到gateio交易所
	// param symbol : 下单品种 example - BTC_USDT 统一大写
	// param price : 下单价格
	// param volume : 下单数量
	// param direction : 下单方向 buy or sell

	// 下单前需要先给api赋值
	defer go_sync.Done()
	if direction=="buy"{
		gateBuy(symbol,price,volume)
	}else{
		gateSell(symbol,price,volume)
	}

}