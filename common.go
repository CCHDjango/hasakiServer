// Server common function
package main

import "time"
import "strings"
import "strconv"
import "fmt"
import "encoding/json"
import "io/ioutil"
import "io"
import "math"
import "log"
import "os"
import "path/filepath"
import "net/http"
import "bufio"

/*-----------------------------网络请求-----------------------------*/
func getHTMLResponse(address string)(*http.Response,error){
	// function : 通过一个网址获取该网页的html对象
	// param address : 网页链接
	// return : 一个http的返回对象
	resp,err:=http.Get(address)
	if err != nil{
		fmt.Println("请求获取html失败 :",err)
	}

	return resp,err
}

func savePic(url string,savePath string){
	// function : 保存图片到本地
	// param savePath : 图片保存路径  xxx/xxx.jpg
	// param url : 图片网上的地址
    imgPath := savePath
    imgUrl := url
    
    res, err := http.Get(imgUrl)
    if err != nil {
        fmt.Println("common function save picture error")
        return
	}
	if strings.Index(res.Status,"200")==-1{
		fmt.Println("save pictrue function return status code : ",res.Status)
		return
	}
    defer res.Body.Close()
    // 获得get请求响应的reader对象
    reader := bufio.NewReaderSize(res.Body, 64 * 1024)
    
    file, err := os.Create(imgPath)
    if err != nil {
        panic(err)
    }
    // 获得文件的writer对象
    writer := bufio.NewWriter(file)
    io.Copy(writer, reader)
}

/*-----------------------------路径相关----------------------------*/
func exePath()string{
	// function : 返回golang二进制文件的启动位置,golang代码打包之后，这个就会
	// 返回当前程序的路径，如果是通过go run 来运行则会返回一个临时启动的执行路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err!=nil{
		printError(err)
		return ""
	}
	return dir
}

/*-----------------------------定制日志----------------------------*/
func logging(module string,content string,logFile string){
	// function : 记录普通级别的信息
	// param module : trace/info/warning/error
	// param logFile : 需要写入到文件的路径
	switch module{
	case "trace":
		trace:=log.New(ioutil.Discard,"trace :",log.Ldate|log.Ltime|log.Lshortfile)
		trace.Println(content)
		break
	case "info":
		info:=log.New(os.Stdout,"info :",log.Ldate|log.Ltime|log.Lshortfile)
		info.Println(content)
		break
	case "warning":
		warning:=log.New(os.Stdout,"warning :",log.Ldate|log.Ltime|log.Lshortfile)
		warning.Println(content)
		break
	case "error":
		// 如果用户传过来的地址是个空值的情况就直接打印内容即可
		if len(logFile)<2{
			print(content)
			return
		}
		file,err:=os.OpenFile(logFile,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
		if err!=nil{
			log.Fatalln("common.go logging function failed to open error log file:",err)
		}
		err2:=log.New(io.MultiWriter(file,os.Stderr),"error :",log.Ldate|log.Ltime|log.Lshortfile)
		err2.Println(content)
		break
	default:
		return
	}
}

//-----------------------------------------------------------------------
func print(content interface{}){
	// function : 简化打印
	fmt.Println(content)
}

func printError(content error){
	// function : 错误提示
	panic(content)
}

func printLocal(content string){
	// function : 写入日志到本地txt文件，日志文件的路径默认在执行文件的同一目录
	dir:=exePath()
	if len(dir)<2{
		return
	}
	err:=mkdir(dir+"/log")
	if err!=nil{
		fmt.Println(err)
	}
}

func dateJudge(firstDate string,lastDate string)(bool){
	/*
	function : 判断前后时间是否一致，如果一致则返回真，否则假
	param firstDate : 第一个时间 格式：2019-12-21 20:05
	param lastDate : 第二个时间
	return : 如果一致则返回真，否则假
	*/
	if firstDate==lastDate{
		return true
	}else{
		return false
	}
}

/*----------------------------------文件操作--------------------------------*/
func readJson(path string)([]map[string]interface{}){
	/*
	function : 读取json文件并转换成[]map类型
	param path : 需要读取的json文件的路径
	return : 返回
	*/
	byteValue, err := ioutil.ReadFile(path)
	if err!=nil{
		fmt.Println("读取json文件报错，读取的文件地址：",path)
	}

	var result []map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("解析json 错误",err)
	}
	
	return result
}

func writeJson(path string,mapData interface{}){
	// function : 写入json 文件
	// param mapData : map类型数据
	if jsonData,err:=json.Marshal(mapData);err==nil{
		err := ioutil.WriteFile(path, jsonData, 0644)
		fmt.Println("写入json",err)
	}else{
		fmt.Println("结果错误",err)
	}
}

func readCSV(path string)[]map[string]interface{}{
	// function : 从本地加载CSV文件，并返回到上层,读取csv
	// param path : CSV文件路径
	var tempData []map[string]interface{}
	return tempData
}

func mkdir(filePath string)error{
	// function : 创建文件夹
	// param filePath : 文件夹路径
	err:=os.Mkdir(filePath,os.ModePerm)
	return err
}

/* ----------------------------时间相关的操作----------------------*/
func sleep(n int){
	/*
	function : 睡眠函数
	param n : 睡眠的秒数
	秒:time.Second
	分钟:time.Minute
	小时:time.Hour
	微妙:time.Nanosecond
	*/
	if n<0{
		return
	}

	for i:=0;i<=n;i++{
		time.Sleep(time.Second)
	}
}

func sleepMin(n int){
	// function : 睡眠多少分钟
	// param n : 睡眠的分钟数
	if n<0{
		return
	}
	for i:=0;i<=n;i++{
		time.Sleep(time.Minute)
	}
}

func sleepHour(h int){
	// function : 睡眠的小时数
	if h<0{
		return
	}
	for i:=0;i<h;i++{
		time.Sleep(time.Hour)
	}
}

func nowTime(m string)(string){
	// function : 返回需要的当前时间字符串
	// return : 返回需要的时间字符串
	now:=time.Now().Format("2006-01-02 15:04:05")
	if m=="day"{
		return string([]byte(now[:10]))              // 返回的示例：2019-12-30
	}else{
		now:=time.Now().Unix()
		return string(now)
	}
}

func integralPoint()(bool){
	// function : 判断当前时间是否为整点
	// 如果不为整点则返回false
	time:=nowTime("day")
	if strings.Index(time,"00:00:")!=-1{
		return true
	}else{
		return false
	}
}

/*--------------------------------类型转换-------------------------------*/
func strJoin(first string,last string,mid string)(string){
	/*
	function : 拼接合成字符串
	param first : 第一段字符串
	param mid : 中间分割的字符串
	param last : 最后一段字符串
	return : 返回拼接好的字符串
	*/
	return strings.Join([]string{first,last},mid)
}

func intToStr(num int)(string){
	/*
	function : 整型转换成字符串
	param num : 整型
	return : 。。
	*/
	return strconv.Itoa(num)
}

func structToJson(s interface{})string{
	// function : 结构体转json字符串
	// return : 返回json字符串
	var data string
	if dataU,err:=json.MarshalIndent(s,"","    ");err==nil{
		// fmt.Println(fmt.Sprintf("%T",dataU)) []uint8
		data=string(dataU)
	}else{
		fmt.Println("结构体转json结果错误",err)
	}
	return data
}

func float64ToInt(f float64)int{
	// function : float64转int
	return int(math.Floor(f))
}

func uint8ToStr(uint8Data []uint8) string {
	// function : uint8转字符串
	return string(uint8Data)
}

func jsonToMap(jsonStr string) map[string]interface{}{
	// function : json转map
	var returnData map[string]interface{}
	err:=json.Unmarshal([]byte(jsonStr),&returnData)
	if err!=nil{
		fmt.Println("json转map报错",jsonStr,"错误原因:",err)
	}
	return returnData
}

func uint8ToMap(uint8Data []uint8) interface{}{
	// uint8转map
	var returnData map[string]interface{}
	err:=json.Unmarshal(uint8Data,&returnData)
	if err!=nil{
		fmt.Println("uint8转map",uint8Data,err)
	}
	return returnData
}

func mapToByte(mapData map[string]interface{}) []byte {
	// function : map 转 byte数组
	mJson,_:=json.Marshal(mapData)
	mJsonStr:=string(mJson)
	mByte:=[]byte(mJsonStr)
	return mByte
}

