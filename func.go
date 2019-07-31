package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/md5"
	"encoding/hex"
	"time"
	"sort"
)

type Result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func main(){
	//httpGet()
	ss := "123456"
	fmt.Printf("MD5(%s): %s\n",ss,MD5(ss))
	times := getTimeStr()
	fmt.Println(times)

	params := map[string]interface{}{
		"name":"张三",
		"age":19,
		"pwd":"12345",
	}
	sign := createSign(params)
	fmt.Println(sign)
	//var res Result
	//res.Code = 200
	//res.Msg = "success"
	//toJson(&res)
	//setJson(&res)
	//toJson(&res)
}
func setJson(res *Result){
	res.Code = 500
	res.Msg = "err"
}
func toJson(res *Result){
	jsonData,err := json.Marshal(res)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("json data: ",string(jsonData))

	// 反序列化
	var res1 Result
	json.Unmarshal([]byte(jsonData),&res1)
	fmt.Println(res1.Msg)
}

func httpGet(){
	response,err := http.Get("http://127.0.0.1:81/api/news_options")
	if err != nil{
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Println("错误:读取body", err)
	}
	fmt.Println(string(body))
}

func MD5(str string)string{
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getTimeStr()  string{
	return  time.Now().Format("2006-01-02 15:04:05")
}

func createSign(params map[string]interface{})  string{
	var key []string
	var str string
	var secret = "123456"
	for k := range params{
		key = append(key,k)
	}

	sort.Strings(key)

	for i := 0;i<len(key) ; i++  {
		if i== 0{
			str = fmt.Sprintf("%v=%v",key[i],params[key[i]])
		}else{
			str += fmt.Sprintf("&%v=%v",key[i],params[key[i]])
		}
	}
	str = MD5(secret + MD5(str))
	return  str
}