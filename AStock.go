/*
hasaki-quant A股接口，通过凤凰api获取
API : http://api.finance.ifeng.com/akdaily/?code=sh000001&type=last
sh000001是上证指数，一次请求，能获取所有上证指数从2017年到现在的数据

取回的数据在record中，列表中的数据的排列分别为:
date,open,high,close,low,volume,chg(涨跌额),p_chg(涨跌幅),
ma5(5日均价),ma10(10日均价),ma20(20日均价),vma5(5日均量)
vma10(10日均量),vma20(20日均量),turnover换手率(指数没有)

注意：这里的请求是请求回日线的行情，所以每天请求一次就可以了
同时周末不需要查询，查询的数据再和数据库最新的数据作对比查重
*/
package main

// 单条行情的数据结构体
type StockADataStruct struct{
	Date string
	Open float32
	High float32
	Close float32
	Low float32
	Volume float32
	Chg float32
	P_chg float32
	Ma5 float32
	Ma10 float32
	Ma20 float32
	Vma5 float32
	Vma10 float32
	Vma20 float32
	Turnover float32
}
// 接口，不一定要实现
type stockInterface interface{
	stockAMain()
	dayCheck()bool
	stockASave()
}

func stockAMain(){
	// function : A股启动入口
}

func dayCheck()(bool){
	// function : 判断是否是周末，周末不启动
	// 工作日是9点半启动，中午11点半休息，1点半到3点再启动
	// return : 如果是正常开市时间就返回true,否则返回false
	openMarket:=true
	return openMarket
}

func stockASave(){
	// function : A股数据保存到数据库
}