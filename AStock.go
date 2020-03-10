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

