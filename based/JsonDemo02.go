package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetEnterpriseByIDResponse struct {
	Ret  int                           `json:"ret"`
	Msg  string                        `json:"msg"`
	Data GetEnterpriseByIDResponseData `json:"data"`
}

type GetEnterpriseByIDResponseData struct {
	Ep GetEnterpriseByIDResponseDataEp `json:"ep"`
}

type GetEnterpriseByIDResponseDataEp struct {
	ID                   int64     `json:"id"`
	EpID                 int64     `json:"ep_id"`
	UserID               int64     `json:"user_id"`
	StaffID              int64     `json:"staff_id"`
	Name                 string    `json:"name"`
	Address              string    `json:"address"`
	Email                string    `json:"email"`
	Tel                  string    `json:"tel"`
	LicenseNo            string    `json:"license_no"`
	LicensePicPath       string    `json:"license_pic_path"`
	StaffNum             int       `json:"staff_num"`
	Status               int       `json:"status"`
	Verify               int       `json:"verify"`
	VerifyTime           string `json:"verify_time"`
	VerifyDesc           string    `json:"verify_desc"`
	CreateTime           string `json:"create_time"`
	ModifyTime           string `json:"modify_time"`
	CityID               int       `json:"city_id"`
	RegRef               string    `json:"reg_ref"`
	ClientType           int       `json:"client_type"`
	OrderTime            string `json:"order_time"`
	RechargeTime         string `json:"recharge_time"`
	Balance              string    `json:"balance"`
	EpTypeID             int64     `json:"ep_type_id"`
	Remark               string    `json:"remark"`
	InvoiceTime          string `json:"invoice_time"`
	Level                int       `json:"level"`
	IsInvoice            int       `json:"is_invoice"`
	PaySortType          int       `json:"pay_sort_type"`
	IsNucReason          int       `json:"is_nuc_reason"`
	StaffScale           int       `json:"staff_scale"`
	CreateUser           string    `json:"create_user"`
	IsTax                int       `json:"is_tax"`
	EnableSMSToRecipient int       `json:"enable_sms_to_recipient"`
	EnableMyFleetOnly    int       `json:"enable_my_fleet_only"`
	MembershipTier       int       `json:"membership_tier"`
	IsSendTopupReminder  int       `json:"is_send_topup_reminder"`
	IsSendStatement      int       `json:"is_send_statement"`
}

// 设置时间格式
const (
	timeFormat = "2006-01-02 15:04:05"
)


type UnixTime time.Time

func (t UnixTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	newTime, err := time.ParseInLocation("\""+timeFormat+"\"", string(data), time.Local)
	if err != nil {
		return nil
	}
	*t = UnixTime(newTime)
	return
}


func main() {
	var result GetEnterpriseByIDResponse

	str := "{\"ret\": 0,\"msg\": \"Success\",\"data\": {\"ep\": {\"id\": 40440,\"ep_id\": 640273088,\"user_id\": null,\"staff_id\": 0,\"name\": \"autotest_coyote_sg\",\"address\": \"\",\"email\": \"autotest_coyote_corp@lalamove.com.sg\",\"tel\": null,\"license_no\": \"\",\"license_pic_path\": \"\",\"staff_num\": 1,\"status\": 2,\"verify\": 1,\"verify_time\": \"2021-01-15 17:19:27\",\"verify_desc\": \"\",\"create_time\": \"2021-01-15 17:19:27\",\"modify_time\": \"2022-05-06 17:00:36\",\"city_id\": 31001,\"reg_ref\": \"\",\"client_type\": 0,\"order_time\": \"2021-01-15 17:19:27\",\"recharge_time\": null,\"balance\": \"\",\"ep_type_id\": 0,\"remark\": \"\",\"invoice_time\": \"0000-00-00 00:00:00\",\"level\": 2,\"is_invoice\": 0,\"pay_sort_type\": 1,\"is_nuc_reason\": 2,\"staff_scale\": 1,\"create_user\": \"chris.susanto\",\"is_tax\": 0,\"enable_sms_to_recipient\": 1,\"enable_my_fleet_only\": 0,\"membership_tier\": 2,\"is_send_topup_reminder\": 2,\"is_send_statement\": 2,\"industry\": null,\"monthly_delivery_volume\": null,\"first_contact_way\": 1,\"on_behalf\": 0,\"platform\": null,\"frequent_vehicle_type\": \"\",\"tax_id\": \"\",\"tax_id_attachment_url\": \"\",\"enable_fleet_toggle\": 0,\"license_no_attachment_url\": \"\"},\"invoice\": []}\n}"


	//json.NewDecoder(strings.NewReader(str)).Decode(&result)

	//json.Unmarshal([]byte(body), &result)

	//fmt.Println(result)


	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ",result)
}
