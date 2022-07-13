package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/shopspring/decimal"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	var creditOrderCancellationFeeAmount decimal.Decimal
	var decimal100 = decimal.NewFromInt(100)
	var creditCancellationFeeDecimal = decimal.NewFromInt(2990).Div(decimal100)
	creditOrderCancellationFeeAmount = decimal.Sum(creditOrderCancellationFeeAmount, creditCancellationFeeDecimal)
	fmt.Println(creditOrderCancellationFeeAmount)

	mapSerial()

	unjsonStruct()
	//unjsonMap()
	//unjsonSlice()
	//
	//mapJson()
	//structJson()
	//slienceJson()
}

// 封装 map 序列化测试函数
func mapSerial()  {
	// 定义map变量
	var m map[string]interface{}
	// 初始化map,获取空间
	m = make(map[string]interface{})
	// 赋值
	m["name"] = "大晕头"
	m["sal"] = 3141.59
	m["age"] = 27
	m["addr"] = [2]string{"北京朝阳", "天津南开"}
	// 将 map 使用 Marshal() 函数进行序列化
	data, err := json.Marshal(m)      // map本身为引用。
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}
	// 查看序列化后的 json 字符串
	fmt.Println("map序列化后 = ", string(data))
}




type student struct {
	//tagjson序列化后是小写
	NameZhou string 	`json:"name_zhou"`
	Age int			`json:"age"`
	HobbitZhou string	`json:"hobbit_zhou"`
	VerifyTime           UnixTime `json:"verify_time"`
	MembershipTier       int       `json:"membership_tier"`
}

func unjsonStruct(){
	var stu student
	fmt.Println("unjson before： ",stu)
	var str = "{\"name_zhou\":\"qqqq\",\"age\":12,\"hobbit_zhou\":\"ttttt\",\"verify_time\": \"2021-01-15 17:19:27\",\"membership_tier\": 2}"
	//var str = "{\"name_zhou\":\"qqqq\",\"age\":12,\"hobbit_zhou\":\"ttttt\",\"membership_tier\": 2,\"verify_time\": \"2021-01-15 17:19:27\"}"
	err := json.Unmarshal([]byte(str), &stu)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ",stu)
}

func unjsonMap(){
	testMap := make(map[string]interface{})
	fmt.Println("unjson before： ",testMap)
	var str = "{\"addr\":\"beijing\",\"age\":123,\"name\":\"typ\"}"
	err := json.Unmarshal([]byte(str), &testMap)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ",testMap)
}

func unjsonSlice(){
	var sli  = make([]map[string]interface{},2)
	fmt.Println("unjson before： ",sli)
	var str = "[{\"addr\":\"beijing111\",\"age\":1231,\"name\":\"typ111\"},{\"addr\":\"beijing222\",\"age\":1222," +
		"\"name\":\"typ222\"}]"
	err := json.Unmarshal([]byte(str), &sli)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ",sli)
}


func mapJson() {
	testMap := make(map[string]interface{})
	testMap["name"] = "typ"
	testMap["age"] = 123
	testMap["addr"] = "beijing"
	json01, err := json.Marshal(testMap)
	if err != nil {
		fmt.Println("json失败")
	}
	fmt.Println(string(json01))
}
type stu struct {
	//tagjson序列化后是小写
	Name string 	`json:"name"`
	Age int			`json:"age"`
	Hobbit string	`json:"hobbit"`
}

func structJson(){
	stu := stu{
		Name:   "qqqq",
		Age:    12,
		Hobbit: "ttttt",
	}
	json01, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json失败")
	}
	fmt.Println(string(json01))
}

func slienceJson() {
	var sli  = make([]map[string]interface{},2)
	testMap := make(map[string]interface{})
	testMap["name"] = "typ111"
	testMap["age"] = 1231
	testMap["addr"] = "beijing111"

	testMap1 := make(map[string]interface{})
	testMap1["name"] = "typ222"
	testMap1["age"] = 1222
	testMap1["addr"] = "beijing222"
	sli[0] = testMap
	sli[1] = testMap1
	json01, err := json.Marshal(sli)
	if err != nil {
		fmt.Println("json失败")
	}
	fmt.Println(string(json01))
}