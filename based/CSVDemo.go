package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Type                 string `basic:TYPE`
	Action               string `basic:ACTION`
	OrderID              string `basic:ORDER ID`
	OrderStatus          string `basic:ORDER STATUS`
	Credit               string `basic:CREDITS (+/-)`
	EPayment             string `basic:E-PAYMENT (+/-)`
	Cash                 string `basic:CASH (+/-)`
	Charge               string `basic:CHARGE`
	FinalPrice           string `basic:FINAL PRICE`
	User                 string `basic:USER`
	UserId               string `basic:USER ID`
	OrderDateTime        string `basic:ORDER DATE/TIME`
	Create               string `basic:CREATED`
	CompletedTime        string `basic:COMPLETED TIME`
	ServiceType          string `basic:SERVICE TYPE`
	OrderPath            string `basic:ORDER PATH`
	Addresses            string `basic:ADDRESSES`
	Distance             string `basic:Distance (km)`
	SpecialRequest       string `advanced:SPECIAL REQUEST ITEM`
	SpecialRequestAmount string `basic:SPECIAL REQUEST`
	OrderRemark          string `basic:ORDER REMARK`
	OrderContact         string `basic:ORDER CONTACT`
	StartingPrice        string `basic:STARTING PRICE`
	ExtraMileagePrice    string `basic:EXTRA MILEAGE PRICE`
	TotalRefundAmount    string `basic:REFUND`
	RefundDateLocal      string `basic:REFUND DATE`
	DriverId             string `advanced:DRIVER ID`
	CouponAmount         string `advanced:COUPON AMOUNT`
	PriorityFee          string `advanced:PRIORITY FEE`
	SurchargeFee         string `advanced:SURCHARGE FEE`
	SurchargeDiscount    string `advanced:SURCHARGE DISCOUNT`
	CsSubsidy            string `advanced:CS SUBSIDY`
	AddOnFee             string `advanced:ADD ON FEE`
}

type Info struct {
	Id      string `csv:""`
	Name    string `csv:""`
}

const (
	REPORT_TYPE_BASIC    = "basic"
	REPORT_TYPE_ADVANCED = "advanced"
)

func GetReportTypeMap() map[string]int {
	return map[string]int{
		REPORT_TYPE_BASIC: 0,
		REPORT_TYPE_ADVANCED: 2,
	}
}

func main() {
	_, tierMap, sortTierIndex := GetHeaders(REPORT_TYPE_BASIC)
	keys := make([]int, 0)
	for k := range tierMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, key := range keys {
	//for i, _ := range tierMap {
		fmt.Println(i, key)
	}

	fmt.Println(tierMap)

	fmt.Println("sortTierIndex", sortTierIndex)

	//
	//clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	//defer clientsFile.Close()
	//
	//clients := []*Client{}
	//infos := []*Info{}
	//
	////if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
	////	panic(err)
	////}
	//for _, client := range clients {
	//	fmt.Println("Hello", client.Name)
	//}
	//
	//if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
	//	panic(err)
	//}
	//
	////clients = append(clients, &Client{Id: "User Name", Name: "John"})
	//
	//clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	//clients = append(clients, &Client{Id: "13", Name: "Fred"})
	//clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	//clients = append(clients, &Client{Id: "15", Name: "Danny"})
	//csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	//
	//
	////err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	//
	//infos = append(infos, &Info{Id: "userId", Name: "j"})
	//infos = append(infos, &Info{Id: "Company Name", Name: "fafsa"})
	//infos = append(infos, &Info{Id: "Corporate ID", Name: "1231"})
	//
	//err = gocsv.MarshalWithoutHeaders(&infos, clientsFile)
	//
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(csvContent) // Display all clients as CSV string
	//
	//
	//
	//var reader = csv.NewReader(strings.NewReader(csvContent))
	//xlsxFile := excelize.NewFile()
	//var row int64 = 1
	//var col = 65
	//
	//for {
	//	fields, err := reader.Read()
	//
	//	for _, field := range fields {
	//		fmt.Println(field)
	//		err := xlsxFile.SetCellValue("Sheet1", string(rune(col)) + strconv.FormatInt(row, 10), field)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		col++
	//		if col >= 90 {
	//			fmt.Println(col)
	//		}
	//	}
	//	row++
	//	col = 65
	//	if err == io.EOF {
	//		break
	//	}
	//	//if err != nil {
	//	//	return nil, err
	//	//}
	//}

}

func GetHeaders(reportType string) ([]string, map[int]string, int) {
	reportTypeMap := GetReportTypeMap()
	csvReflectType := reflect.ValueOf(Client{}).Type()
	var num = csvReflectType.NumField()
	var headers []string
	var tierMap = make(map[int]string)
	reportTypeValue := reportTypeMap[reportType]
	sortTierIndex := 0
	for i := 0; i < num; i++ {
		tag := string(csvReflectType.Field(i).Tag)
		s := strings.Split(tag, ":")
		if reportTypeMap[s[0]] <= reportTypeValue {
			headers = append(headers, s[1])
			tierMap[i] = s[0]
		}
		if s[1] == "CREATED" {
			sortTierIndex = i
		}
	}
	return headers, tierMap, sortTierIndex
}