## Hasaki Server Center

服务不会自动连接交易所，用户需要服务去连接交易所需要在配置文件中登记

启动websocket服务，接受用户的连接，如果用户没有连接，那么服务就暂停请求交易所，为了节约资源

交易服务和策略容器连接流程：

1，策略容器连接交易服务，然后向交易服务发送订阅信息和验证信息，个人用户不需要验证，使用hasaki服务的容器才需要验证
用户传过来的账户信息不要持久化处理，也就是数据中台不能把用户的api key写到文件，策略端才保存用户的api key的数据。

2，交易服务接受到策略容器的消息后，再去请求交易所,接收交易所的行情，转发回策略容器

3，交易服务收到策略容器的发单，并转发到交易所,交易所返回的订单信息，则传到策略，另外数据中台提供信息通知的功能，
用户可以自定义绑定重要信息发送到钉钉机器人(建议使用钉钉机器人，非常方便)

4，数据服务除了一个websocket连接用于实时交易之外，还提供一个restful数据请求服务.
hasaki server自身的从外界获取到的数据会保存在数据库中，用户可以通过外部请求下载过往的历史数据用于做数据分析或者策略数据预热

#### connect exchange

服务中心连接交易所

gateio : gateio数字货币交易所

#### crawl the news

服务中心爬取舆情数据

govNewsCroll : 中华人民共和国政府网

#### start websocket

服务中心启动websocket服务接收策略的发单并把行情新闻数据发给策略

消息推送分两种:  1,服务等待请求，收到请求之后再一次性推送. 2,订阅模式，绑定钉钉机器人等，定时推送.

#### data saving

行情数据和新闻舆情数据都默认保存在数据库Mysql和mongoDB,服务默认是一直持久化保存数据，可选只保存一天的数据，因为如果
部署到云服务的话，云服务的数据库很贵，所以保存一天的数据，而由外部的硬盘来保存所有的数据则是可行的方案，若用户是在自己设备上
运行服务，则我们认为，用用户自身的硬盘的是安全且无限的。

#### config

用户使用配置文件定义服务，配置文件可以指定网关和爬虫只执行部分功能，方便在墙外部署同一个服务中台，墙外的服务中台可以不用去连接国内的
交易所和网站，节省资源。

config.json包括:每个交易所的apikey设置，订阅交易所，订阅新闻爬取种类