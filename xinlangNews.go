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
import "strings"
import "github.com/PuerkitoBio/goquery"

var xinlangGlobalNewsTitle string = ""   // 新浪国际新闻新闻标题
var xinlangGlobalNewsTime string = ""    // 新浪国际新闻文章时间
var xinlangGlobalNewsContent string = "" // 新浪国际新闻文章内容

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
	var xlGlobalNewAddress string = "https://news.sina.cn/gj?vt=1&pos=8"
	var linkList []string      // 新浪国际新闻滚动新闻的链接列表
	resp,err:=getHTMLResponse(xlGlobalNewAddress)
	if err!=nil{
		fmt.Println("get xinlangGlobalHTML list failed")
	}
	doc,errd:=goquery.NewDocumentFromReader(resp.Body)
	if errd!=nil{
		fmt.Println("explain html list doc failed")
	}
	doc.Find("a").Each(func(i int,s *goquery.Selection){
		href,isExist := s.Attr("href")
		if isExist==true{
			// 过滤http
			if strings.Index(href,"https")!=-1{
				linkList=append(linkList,href)
			}
		}
	})
	fmt.Println(linkList)
	for _,linkValue:=range(linkList){
		xinlangGlobalHTML(linkValue)
	}
	
}

func xinlangGlobalHTML(link string){
	// function : 通过链接获取新闻滚动的html
	resp,err:=getHTMLResponse(link)
	if err!=nil{
		fmt.Println("get xinlangGlobalHTML failed")
	}
	doc,errd:=goquery.NewDocumentFromReader(resp.Body)
	if errd!=nil{
		fmt.Println("explain html doc failed")
	}
	doc.Find("div").Each(func(i int,s *goquery.Selection){
		title := s.Find("h1").Text()
		content := s.Find("p").Text()
		time:=nowTime("x")      // TODO : 这里拿的时间应该是包含小时分钟的
		if len(title)>4 && len(content)>4{
			fmt.Println(title,"--",content,"++",time,"\n")
		}
		
	})
	// xinlangGlobalCheckSame()
}

func xinlangGlobalTitle(title string){
	// function : 处理新闻的标题
	xinlangGlobalNewsTitle=title
}

func xinlangGlobalTime(time string){
	// function : 处理新闻的时间
	xinlangGlobalNewsTime=time
}

func xinlangGlobalContent(content string){
	// function : 处理新闻的简要内容
	xinlangGlobalNewsContent=content
}

func xinlangGlobalCheckSame(title string,time string){
	// function : 把爬取到的内容和数据库的最近内容
	// 进行对比，如果有重复那就直接结束该次保存即可
}