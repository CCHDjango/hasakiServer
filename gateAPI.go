/*
连接gateio的API

每次发单前，请求一次行情，然后根据最新的行情发出订单

最后一次修改时间 : 2020-1-15

*/
package main

import (
	"crypto/hmac"
	"crypto/sha512"
	// "encoding/hex"
	// "encoding/json"
	"net/http"
	// "net/url"
	// "sort"
	"io/ioutil"
	"strings"
	"fmt"
)

var gateKEY  = "gate.io api key"; // gate.io api key
var gateSECRET = "gate.io api secret";  // gate.io api secret

func gateMain() {
	// Method call
	// all pairs
	var ret string = gateSell("EOS_USDT","3.7800","1")
	fmt.Println(ret)
}

// all support pairs
func gateGetPairs() string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/pairs"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Market Info
func gateMarketinfo() string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/marketinfo"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Market Details
func gateMarketlist() string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/marketlist"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// tickers
func gateTickers() string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/tickers"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// ticker
func gateTicker(ticker string) string {
	// function : 返回tick数据
	// param ticker : 示例eos_usdt
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/ticker" + "/" + ticker
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Depth
func gateOrderBooks() string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/orderBooks"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Depth of pair
func gateOrderBook(params string) string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/orderBook/" + params
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Trade History
func gateTradeHistory(params string) string {
	var method string = "GET"
	var url string = "http://data.gateio.life/api2/1/tradeHistory/" + params
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// Get account fund balances
func gateBalances() string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/balances"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// get deposit address
func gateDepositAddress(currency string) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/depositAddress"
	var param string = "currency=" + currency
	var ret string = httpDo(method,url,param)
	return ret
}

// get deposit withdrawal history
func gateDepositsWithdrawals(start string, end string) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/depositsWithdrawals"
	var param string = "start=" + start + "&end=" + end
	var ret string = httpDo(method,url,param)
	return ret
}

// Place order buy
func gateBuy(currencyPair string, rate string, amount string) string {
	// function : 下单函数
	// param currencyPair : 交易品种  EOS_USDT
	// param rate : 交易价格
	// param amount : 成交量
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/buy"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = httpDo(method,url,param)
	return ret
}

// Place order sell
func gateSell(currencyPair string, rate string, amount string) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/sell"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = httpDo(method,url,param)
	return ret
}

// Cancel order
func gateCancelOrder(orderNumber string, currencyPair string ) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/cancelOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}

// Cancel all orders
func gateCancelAllOrders( types string, currencyPair string ) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/cancelAllOrders"
	var param string = "type=" + types + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}

// Get order status
func gateGetOrder( orderNumber string, currencyPair string ) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/getOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}

// Get my open order list
func gateOpenOrders() string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/openOrders"
	var param string = ""
	var ret string = httpDo(method,url,param)
	return ret
}

// 获取我的24小时内成交记录
func gateMyTradeHistory( currencyPair string, orderNumber string) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/tradeHistory"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = httpDo(method,url,param)
	return ret
}

// Get my last 24h trades
func gateWithdraw( currency string, amount string, address string) string {
	var method string = "POST"
	var url string = "https://api.gateio.life/api2/1/private/withdraw"
	var param string = "currency=" + currency + "&amount=" + amount + "&address=" + address
	var ret string = httpDo(method,url,param)
	return ret
}

func getSign( params string) string {
    key := []byte(gateSECRET)
    mac := hmac.New(sha512.New, key)
    mac.Write([]byte(params))
    return fmt.Sprintf("%x", mac.Sum(nil))
}

/**
*  http request
*/
func httpDo(method string,url string, param string) string {
    client := &http.Client{}

    req, err := http.NewRequest(method, url, strings.NewReader(param))
    if err != nil {
        // handle error
    }
    var sign string = getSign(param)

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("key", gateKEY)
    req.Header.Set("sign", sign)

    resp, err := client.Do(req)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

 	return string(body);
}