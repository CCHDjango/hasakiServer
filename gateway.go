/*
hasaki-quant 数据中台连接交易所的网关
所有交易所的连接都通过gateway去调度
网关包括请求所有交易所的价格行情，所有接口的数据最后在gateway中进行抽象

网关还包含统一发单的功能，如果交易所被墙，那么在外网部署另一个hasaki-server，然后国内的服务中台的网关
连接墙外的数据中台进行数据订阅
所以hasaki-quant data server center做一个分支，若是在墙外启动，则在服务中台的配置文件中进行设置，只执行部分功能
*/
package main

