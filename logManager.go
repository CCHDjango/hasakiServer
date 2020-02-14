/*
hasaki-quant server log system
管理数据中台运行的所有日志数据
包括系统日志，debug日志，错误日志，日志保存，日志发送到用户前端,终端日志
*/
package main

type Log struct{

}
type logInterface interface{
	setting(ding bool,email bool,phoneMsg bool)
	ding(dingAddress string)
	email(emailAddress string)
}

// 钉钉推送消息
func (l *Log) ding(dingAddress string){

}

// 发送到用户的邮箱
func (l *Log) email(emailAddress string){

}

// 设置服务推送给用户前端的方式
func (l *Log) setting(ding bool,email bool,phoneMsg bool){
	// param ding : 是否启动钉钉
	// param email : 推送钉钉开关
	// param phoneMsg : 推送到手机短信的开关
}
