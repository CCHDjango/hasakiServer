/*
hasaki-quant server center的爬虫调度中心
所有爬虫程序都会通过该爬虫中心进行调度处理

定时轮询爬虫，在该代码中设置定时轮询的时间，方便修改
在数据中台中设置定时轮询的间隔
*/
package main

type Crawl struct{

}
type crawlInterface interface{
	setting(target string,saveMethod string)
	startGovNewsCroll(saveMethod string,checkTime int)
}

func (c *Crawl) setting(target string,saveMethod string,checkTime int){
	// function : 爬虫中心的设置，设置需要爬取的站点，保存方式,每次爬虫的轮询间隔
	// param target : 需要爬取的对象站点
	// param saveMethod : 保存数据的方式
	// param checkTime : 爬虫轮询间隔

	switch target {
	case "govNewsCroll":
		go_sync.Add(1)
		go c.startGovNewsCroll(saveMethod,checkTime)
		go_sync.Wait()
	case "xinLangCaiJing":
		break
	}
}

func (c *Crawl) startGovNewsCroll(saveMethod string,checkTime int){
	// function : 启动爬取中国政府官网滚动新闻
	defer go_sync.Done()
	for {
		sleep(checkTime)
	}
}