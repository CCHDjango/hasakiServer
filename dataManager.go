/*
hasaki-quant server center data manager code
所有数据操作全部在dataManager中进行统一调度
包括行情数据和爬虫数据，websocket获取数据并返回到外部

行请处理:
从gateway对象获取行请，行请通过websocket发送到用户策略中，行请保存到数据库中,最新一条行请保存在内存中
方便以后给用户发送价格警告

舆情处理:
从crawlCenter中获取的舆情新闻，hasaki-quant数据中台还没有能力做NLP处理，所以
新闻舆情会直接保存到数据库，用户可以指定某些字符串，新闻中有相关字符串则给用户地址
发送新闻

订单处理:
gateway订单信息触发后，传递订单信息到dataManager，由dataManager传给策略，
订单详情保存到数据库
*/
package main

import "strings"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type DataManager struct{
	Gateway Gateway    // TODO : golang不需要这种设计
	GatewayQuoteChan chan map[string]interface{}
	QuoteChan chan map[string]interface{}
	NewsChan chan string
	CrawlNewsChan chan string
	Session *mgo.Session
	Mysql *sql.DB
}
type dataManagerInterface interface{
	setting(quoteChan ,gatewayQuoteChan chan map[string]interface{},newsChan ,crawlNewsChan chan string)
	dataMain(mgoPath string,mySqlPath string)
	saveAsMongoDB(session *mgo.Session ,datasetName string,tableName string,content map[string]interface{})
	findAllInMgo(session *mgo.Session,datasetName string,tableName string) bool
	deleteInMgo(session *mgo.Session,datasetName string,tableName string,identity string) error
	findAllInSql(sqlDB *sql.DB,order string)
}

func (d *DataManager) oldSetting(gateway *Gateway){
	// function : 设置数据控制中心
	// param gateway : hasaki - quant 网关对象
	d.Gateway=*gateway
}

/*------------------------------获取行请数据--------------------------------*/
func (d *DataManager) recvQuote()map[string]interface{}{
	// function : 获取到行请并做分发，注意阻塞，行请分发到保存在数据库，并把行请传到websocket发送队列里
	// return : 行请数据
	var quote map[string]interface{}
	quote = <- d.Gateway.QuoteChan

	// 保存到数据库
	go_sync.Add(1)
	go d.saveAsMongoDB(d.Session,"","",quote)
	go_sync.Wait()

	// 传给websocket
	d.QuoteChan <- quote

	return quote
}


func (d *DataManager) dataMain(mgoPath string,mySqlPath string){
	// function : 数据调度模块的启动接口
	// param mySqlPath : mysql路径 示例 : root:userName@tcp(address:port)/datasetName?charset=utf8
	session,err:=mgo.Dial(mgoPath)
	if err!=nil{
		printError(err)
	}
	d.Session=session
	//var sqlDB *sql.DB
	mysql,err:=sql.Open("mysql",mySqlPath)
	if err!=nil{
		printError(err)
	}
	d.Mysql=mysql
	// TODO : 这是个sql的示例,第二个参数是sql命令字符串
	d.findAllInSql(mysql,"")

	// TODO : 这是一个示例,之后换成业务代码
	content:=map[string]interface{}{"title":"","content":"","date":"","id":"","from":""}
	d.saveAsMongoDB(session,"crawl","govNews",content)
}

/*--------------------------数据模块设置-----------------------------------*/
func (d *DataManager) setting(quoteChan ,gatewayQuoteChan chan map[string]interface{},newsChan ,crawlNewsChan chan string){
	// function : 数据模块的设置部分
	// param quoteChan : 数据通道，这个是数据模块处理好行情数据之后，分发到websocket模块的
	// param gatewayQuote : 数据管道，从gateway网关把外部数据传到数据处理模块
	// param newsChan : 新闻舆情管道，数据模块处理后的新闻舆情分发到websocket模块
	// param crawlNewsChan : 新闻舆情管道，从爬虫模块传新闻舆情数据到数据处理模块
	d.QuoteChan=quoteChan
	d.GatewayQuoteChan=gatewayQuoteChan
	d.NewsChan=newsChan
	d.CrawlNewsChan=crawlNewsChan
}

/*--------------------------mongo数据库操作-----------------------------------*/
func (d *DataManager) saveAsMongoDB(session *mgo.Session ,datasetName string,tableName string,content map[string]interface{}){
	// function : 保存新闻舆情数据到mongo数据库
	// 读表
	// param session : mgo数据库的操作对象
	// param datasetName : 选择mgo的数据库名字
	// param tableName : 指定mgo数据库的表
	// param conotent : 需要插入到数据库的内容 参考示例 : {"title":title,"content":content,"date":time,"id":id,"from":dataFrom}
	defer go_sync.Done()
	c:=session.DB(datasetName).C(tableName)
	c.Insert(content)
	print("insert data in mgo") // 这个打印不是必要的,如果打印次数过多，这个会导致日志臃肿
}

func (d *DataManager) findAllInMgo(session *mgo.Session,datasetName string,tableName string)(bool){
	// function : 根据标识来查询符合条件的所有数据

	// TODO : 这个需要从外部传进来
	type TempStruct struct{
		Date string `bson:"date"`
		Content string `bson:"content"`
		Title string `bson:"title"`
		Id string `bson:"id"`
		From int `bson:"from"`
	}
	var result []TempStruct
	err:=session.DB(datasetName).C(tableName).Find(nil).All(&result)
	if err!=nil{
		printError(err)
		return false
	}
	return true
}

func (d *DataManager) deleteInMgo(session *mgo.Session,datasetName string,tableName string,identity string)(error){
	// function : 根据表示来删除内容

	// TODO : 这个需要从外部传进来
	type TempStruct struct{
		Date string `bson:"date"`
		Content string `bson:"content"`
		Title string `bson:"title"`
		Id string `bson:"id"`
		From int `bson:"from"`
	}
	var tempS []TempStruct
	c:=session.DB(datasetName).C(tableName)
	err:=c.Find(nil).Limit(1).All(&tempS)
	if len(tempS)==0{
		// 如果一开始数据库就没有，那么就跳过
		return nil
	}
	// 判断数据库里面的数据是否与当天的时间一致
	if strings.Index(tempS[0].Date,identity)==-1{
		c.RemoveAll(bson.M{"date": tempS[0].Date})
	}
	
	return err
}

/*-------------------------------mysql数据库操作------------------------------*/
func (d *DataManager) findAllInSql(sqlDB *sql.DB,order string){
	// function : 查询mysql数据库
}

/*------------------------------操作本地数据-------------------------------*/
