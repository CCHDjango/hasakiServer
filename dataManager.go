/*
hasaki-quant server center data manager code
所有数据操作全部在dataManager中进行统一调度
包括行情数据和爬虫数据，websocket获取数据并返回到外部
*/
package main

import "strings"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func dataMain(mgoPath string,mySqlPath string){
	// function : 数据调度模块的启动接口
	// param mySqlPath : mysql路径 示例 : root:userName@tcp(address:port)/datasetName?charset=utf8
	session,err:=mgo.Dial(mgoPath)
	if err!=nil{
		printError(err)
	}
	//var sqlDB *sql.DB
	mysql,err:=sql.Open("mysql",mySqlPath)
	if err!=nil{
		printError(err)
	}
	// TODO : 这是个sql的示例
	findAllInSql(mysql,"")

	// TODO : 这是一个示例,之后换成业务代码
	content:=map[string]interface{}{"title":"","content":"","date":"","id":"","from":""}
	saveAsMongoDB(session,"crawl","govNews",content)
}

/*--------------------------mongo数据库操作-----------------------------------*/
func saveAsMongoDB(session *mgo.Session ,datasetName string,tableName string,content map[string]interface{}){
	// function : 保存新闻舆情数据到mongo数据库
	// 读表
	// param session : mgo数据库的操作对象
	// param datasetName : 选择mgo的数据库名字
	// param tableName : 指定mgo数据库的表
	// param conotent : 需要插入到数据库的内容 参考示例 : {"title":title,"content":content,"date":time,"id":id,"from":dataFrom}
	c:=session.DB(datasetName).C(tableName)
	c.Insert(content)
	print("insert data in mgo") // 这个打印不是必要的,如果打印次数过多，这个会导致日志臃肿
}

func findAllInMgo(session *mgo.Session,datasetName string,tableName string)(bool){
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

func deleteInMgo(session *mgo.Session,datasetName string,tableName string,identity string)(error){
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
func findAllInSql(sqlDB *sql.DB,order string){
	// function : 查询mysql数据库
}