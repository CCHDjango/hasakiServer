## Hasaki Server Center

服务不会自动连接交易所，用户需要服务去连接交易所需要在配置文件中登记

启动websocket服务，接受用户的连接，如果用户没有连接，那么服务就暂停请求交易所，为了节约资源

#### connect exchange

服务中心连接交易所

#### crawl the news

服务中心爬取舆情数据

#### start websocket

服务中心启动websocket服务接收策略的发单并把行情新闻数据发给策略

#### config

用户使用配置文件定义服务
