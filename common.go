// Server common function
package main

import "time"
import "strings"
import "strconv"
import "fmt"
import "encoding/json"
import "io/ioutil"
import "math"

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
	// function : 从本地加载CSV文件，并返回到上层
	// param path : CSV文件路径
	var tempData []map[string]interface{}
	return tempData
}

func sleep(n int){
	/*
	function : 睡眠函数
	param n : 睡眠的秒数
	*/
	if n<0{
		return
	}

	for i:=0;i<=n;i++{
		time.Sleep(time.Second)
	}
}

func nowTime(m string)(string){
	// function : 返回需要的当前时间字符串
	// return : 返回需要的时间字符串
	now:=time.Now().Format("2006-01-02 15:04:05")
	if m=="day"{
		return string([]byte(now[:10]))              // 返回的示例：2019-12-30
	}
	return "x"
}

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
