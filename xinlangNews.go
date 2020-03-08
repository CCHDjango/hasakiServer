/*
新浪新闻，国际新闻爬虫
爬虫的地址不是PC端的新浪地址，而是用手机的新浪新闻地址：
https://news.sina.cn/gj?vt=1&pos=8       vt=1的时候是简化版

能直接拿到题目，但是要进去到页面才能拿到时间，链接的标签是
a class="f_card_m_f_a_r"

进入链接之后，新闻标题是: h1 class="art_tit_h1"
时间的标签是：time class="art_time"

只要拿到标题和时间即可，文章内容不需要爬取，注意是一个页面，
爬取到内容直接保存到数据库即可
*/
package main

import "time"
import "fmt"

func xinlanNewsMain(){
	// function : 启动新浪国际新闻爬虫
	// 半个小时更新一次
	for {
		fmt.Println("爬虫开始")
		time.Sleep(time.Minute * 30)
	}
}

func xinlangGlobalLink(){
	// function : 在新闻滚动页面获取到每个新闻的链接

}

func xinlangGlobalTitle(){
	// function : 获取新闻的标题
}

func xinlangGlobalTime(){
	// function : 获取新闻的时间
}

func xinlangGlobalContent(){
	// function : 获取新闻的简要内容
}

func xinlangGlobalCheckSame(){
	// function : 把爬取到的内容和数据库的最近内容
	// 进行对比，如果有重复那就直接结束该次保存即可
}