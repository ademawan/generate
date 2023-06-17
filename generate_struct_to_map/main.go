package main

import (
	"fmt"
	"generate-struct-to-map/rogerdev_generator"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type (
	UserRequestCreate struct {
		Email     string  `json:"email" test:"-"`
		Name      string  `json:"name" test:"-"`
		Username  string  `json:"username" test:"-"`
		Phone     string  `json:"phone" test:"-"`
		Password  string  `json:"password" `
		Photo     string  `json:"photo" `
		Latitude  float64 `json:"latitude" `
		Longitude float64 `json:"longitude" `
	}
	MerchantRequestCreate struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		OtpToken string `json:"otp_token"`
	}
)
type Base struct {
	id   int
	name string
}

type Extended struct {
	Base
	Email    string
	Password string
}

//	type ResponseSubscriberProfile struct {
//		Profiles Profiles `json:"profile"`
//	}
type ExampleStructToMap struct {
	Name  string `json:"user_name,omitempty"`
	Email string `json:"user_email"`
	Price int    `json:"price"`
}
type InvoiceMerchantPanelHttp struct {
	InvoiceTypeId    string `json:",omitempty"`
	Name             string `json:",omitempty"`
	Msisdn           string `json:",omitempty"`
	Email            string `json:",omitempty"`
	CallableContact  string `json:",omitempty"`
	PackageId        string `json:",omitempty"`
	PackageName      string `json:",omitempty"`
	PackagePrice     int    `json:",omitempty"`
	PackagePeriod    string `json:",omitempty"`
	DeviceType       string `json:",omitempty"`
	InvoiceNumber    string `json:",omitempty"`
	ImageInvoice     string `json:",omitempty"`
	ImageStarterpack string `json:",omitempty"`
}
type Invoice struct {
	Id                   string
	TransactionNumber    string
	OrderIdEsb           string
	OrderId              string
	PaymentId            string
	Msisdn               string
	PaymentMethod        string
	TotalDiscount        int
	TotalBill            int
	TotalTax             int
	Qty                  int
	Cashback             int
	Admin_Fee            int
	UserId               string
	Tag                  string
	Product              Product
	StatusTransaction    string
	CreatedAt            string
	UpdatedAt            string
	OrderNumber          string
	URLPayload           string
	TransactionPackageId string
}
type Product struct {
	OrderId              string
	DeviceId             string
	Msisdn               string
	Email                string
	TotalBill            int
	UserId               string
	Tag                  string
	OrderNumber          string
	UrlPayload           string
	TransactionPackageId string
	TransactionStatus    string
	OrderDate            string
	ActivePeriode        string
	ProductLists         []ProductList
}
type ProductList struct {
	ID            string
	Package_Id    string
	Name          string
	Price         int
	Validity      string
	Bid           string
	Autorenewal   string
	PaymentMethod string
	Tax           int
	Discount      int
	Description   string
	Quantity      int
	Admin_Fee     int
}

type InvoiceWna struct {
	Id                string
	TransactionNumber string
	OrderId           string
	OrderIdEsb        string
	UserId            string
	PaymentId         string
	WnaId             string
	PaymentMethod     string
	PurchaseMode      string
	ItemName          string
	ItemPrice         int
	Description       string
	CaseId            string
	ItemQty           int
	CustomerName      string
	Gender            string
	Passport          string
	Imei1             string
	Imei2             string
	Nationality       string
	DateOfBirth       string
	TotalDiscount     int
	TotalBill         int
	TotalTax          int
	Cashback          int
	DeviceId          string
	Email             string
	PaidOn            string
	StatusTransaction string
	CreatedAt         string
	UpdatedAt         string
}

type OrderWna struct {
	OrderId           string
	OrderIdEsb        string
	UserId            string
	WnaId             string
	TransactionNumber string
	PaymentMethod     string
	PurchaseMode      string
	ItemName          string
	ItemPrice         int
	Description       string
	CaseId            string
	CustomerName      string
	Gender            string
	Passport          string
	Imei1             string
	Imei2             string
	TotalDiscount     int
	TotalBill         int
	TotalTax          int
	DeviceId          string
	Nationality       string
	DateOfBirth       string
	Email             string
	TransactionStatus string
}
type InvoiceBill struct {
	Id                   string
	TransactionNumber    string
	OrderId              string
	PaymentId            string
	Msisdn               string
	PaymentMethod        string
	CustomerName         string
	Periode              string
	TotalDiscount        int
	TotalBill            int
	TotalTax             int
	Qty                  int
	Cashback             int
	Tag                  string
	Name                 string
	UserId               string
	DeviceId             string
	Email                string
	StatusTransaction    string
	CreatedAt            string
	UpdatedAt            string
	OrderNumber          string
	URLPayload           string
	TransactionPackageId string
}
type InvoiceCredit struct {
	Id                   string
	TransactionNumber    string
	OrderId              string
	PaymentId            string
	CreditId             string
	Msisdn               string
	PaymentMethod        string
	CreditName           string
	Periode              string
	TotalDiscount        int
	TotalBill            int
	TotalTax             int
	Qty                  int
	Cashback             int
	UserId               string
	Tag                  string
	DeviceId             string
	Email                string
	StatusTransaction    string
	CreatedAt            string
	UpdatedAt            string
	OrderNumber          string
	URLPayload           string
	TransactionPackageId string
}
type InvoiceHalo struct {
	ID                string `json:"-"`
	TransactionNumber string `json:"transaction_number"`
	OrderID           string `json:"order_id"`
	Msisdn            string `json:"msisdn"`
	HaloMsisdn        string `json:"halo_msisdn"`
	ItemName          string `json:"item_name"`
	ItemPrice         int    `json:"item_price"`
	ItemQty           int    `json:"item_qty"`
	TotalBill         int    `json:"total_bill"`
	StatusTransaction string `json:"status_transaction"`
	Description       string `json:"description"`
	Email             string `json:"email"`
	Period            string `json:"period"`
	Fullname          string `json:"fullname"`
	Pob               string `json:"pob"`
	Dob               string `json:"dob"`
	CustomerValue     string `json:"identity_value"`
	Nokk              string `json:"nokk"`
	IdentityAddress   string `json:"identity_address"`
	CreatedAt         string `json:"-"`
	UpdatedAt         string `json:"-"`

	Status           bool   `json:"status"`
	Message          string `json:"message"`
	ValidationStatus string `json:"validation_status"`
	// TransactionID string `json:"transaction_id"`
	// StatusDesc    string `json:"status_desc"`
}

type InvoiceHaloPlus struct {
	Id                   string
	TransactionNumber    string
	OrderId              string
	PaymentId            string
	DeviceId             string
	Msisdn               string
	PaymentMethod        string
	TotalDiscount        int
	TotalBill            int
	TotalTax             int
	Cashback             int
	Product              Product
	StatusTransaction    string
	CreatedAt            string
	UpdatedAt            string
	URLPayload           string
	TransactionPackageId string
	OrderDate            string
	TypePayment          string
	ActivePeriode        string
	AuthId               string
}

func main() {
	// generator := rogerdev_generator.New()
	// res := generator.RandStringRunes(10)
	// fmt.Println(res)

	// check := generator.GenerateName(6, true)
	// fmt.Println(check)
	// var data []byte
	// GetFileJson("", "file", "json", &data)
	// // fmt.Println(data)
	// dataMap := make(map[string]interface{})
	// json.Unmarshal(data, &dataMap)
	// // fmt.Println(dataMap)

	// st := UserRequestCreate{}
	// GenerateMapToStruct(dataMap, "coba", "test")
	// getPropertyInfo(&st)
	generateToMap(&MerchantRequestCreate{})
	generateToStruct(&InvoiceBill{})

}
func getPropertyInfo(s interface{}) {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f, _ := reflectType.FieldByName("Email")
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Lookup("test"))
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func generateToMap(s interface{}) {
	generator := rogerdev_generator.New()
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	//initial package
	tmp := `package file
	
	import "fmt"
	`

	//write function
	tmp += `
	func example(){

	

	`

	tmp += `data := map[string]interface{}{
		`

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f, _ := reflectType.FieldByName(reflectType.Field(i).Name)
		// fmt.Println(f.Tag)
		fmt.Println(f.Tag.Lookup("json"))
		field, check := f.Tag.Lookup("json")
		if check {
			fieldArray := strings.Split(field, ",")
			if len(fieldArray) == 0 {
				field = fmt.Sprintf("%v", reflectType.Field(i).Name)
			} else if len(fieldArray) == 1 {
				if fieldArray[0] == "omitempty" {
					field = fmt.Sprintf("%v", reflectType.Field(i).Name)
				}

			} else if len(fieldArray) > 1 {
				for _, val := range fieldArray {
					if val != "omitempty" {
						if len(val) != 0 {
							field = val

						}

					}

				}
			}
			if isExists := strings.Contains(field, ","); isExists {
				field = fmt.Sprintf("%v", reflectType.Field(i).Name)

			}
		} else {
			field = fmt.Sprintf("%v", reflectType.Field(i).Name)
		}

		tmp += `"` + field + `":`
		typeData := fmt.Sprintf("%v", reflectType.Field(i).Type)
		switch typeData {
		case "string":
			tmp += `"`
			if isExists := strings.Contains(strings.ToLower(fmt.Sprintf("%v", reflectType.Field(i).Name)), "name"); isExists {
				tmp += generator.GenerateName(10, true)
			} else if isExists := strings.Contains(strings.ToLower(fmt.Sprintf("%v", reflectType.Field(i).Name)), "email"); isExists {
				tmp += generator.GenerateEmail()
			} else {
				tmp += generator.GenerateString(5)
			}
			tmp += `"`

		case "int":
			tmp += strconv.Itoa(generator.GeneratePrice())
		case "float64":
		case "interface{}":

		}
		tmp += `,
		`

		// fmt.Println("nama      :", reflectType.Field(i).Name)
		// fmt.Println("tipe data :", reflectType.Field(i).Type)
		// fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		// fmt.Println("")
	}
	//endMapVariable
	tmp += `
	}
	`

	tmp += `fmt.Println(data)
	`

	//endFunction
	tmp += `}`
	_ = ioutil.WriteFile("file/example_struct_to_map"+".go", []byte(tmp), 0644)

}

func generateToStruct(s interface{}) {
	generator := rogerdev_generator.New()
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()

	//temporary for type struct
	// var tmpStruct []string

	//initial package
	tmp := `package file
	
	import "fmt"
	`

	//create struct
	tmp += ` 
	type ExampleStruct struct{

	
	`

	for i := 0; i < reflectValue.NumField(); i++ {

		// field, _ := f.Tag.Lookup("json")
		field := fmt.Sprintf("%v", reflectType.Field(i).Name)
		typeData := fmt.Sprintf("%v", reflectType.Field(i).Type)
		// if isExists := strings.Contains(typeData, "."); isExists {
		// 	dt := strings.Split(typeData, ".")

		// 	if exist := strings.Contains(typeData, "[]"); exist {

		// 	} else {
		// 		tmpStruct = append(tmpStruct, dt[1])
		// 		typeData = dt[1]
		// 	}

		// }
		// fmt.Println(typeData)
		tmp += field + ` ` + typeData + `
		`
	}

	//end create struct
	tmp += `}
	`

	//write function
	tmp += `
	func exampleToStruct(){

	

	`

	tmp += `data := ExampleStruct{}
		`

	for i := 0; i < reflectValue.NumField(); i++ {

		field := fmt.Sprintf("%v", reflectType.Field(i).Name)

		if isExists := strings.Contains(field, ","); isExists {
			field = fmt.Sprintf("%v", reflectType.Field(i).Name)

		}
		tmp += `data.` + field + `=`
		typeData := fmt.Sprintf("%v", reflectType.Field(i).Type)
		switch typeData {
		case "string":
			tmp += `"`
			if isExists := strings.Contains(strings.ToLower(fmt.Sprintf("%v", reflectType.Field(i).Name)), "name"); isExists {
				tmp += generator.GenerateName(10, true)
			} else if isExists := strings.Contains(strings.ToLower(fmt.Sprintf("%v", reflectType.Field(i).Name)), "email"); isExists {
				tmp += generator.GenerateEmail()
			} else {
				tmp += generator.GenerateString(5)
			}
			tmp += `"`

		case "int":
			tmp += strconv.Itoa(generator.GeneratePrice())
		case "float64":
		case "bool":
			tmp += "true"
		case "interface{}":

		}

		tmp += `
		`

		// fmt.Println("nama      :", reflectType.Field(i).Name)
		// fmt.Println("tipe data :", reflectType.Field(i).Type)
		// fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		// fmt.Println("")
	}

	tmp += `
	fmt.Println(data)
	`

	//endFunction
	tmp += `}`
	_ = ioutil.WriteFile("file/example_to_struct"+".go", []byte(tmp), 0644)

}

func GenerateMapToStruct(dataMap map[string]interface{}, nameOfStruct, nameOfFile string) {
	if len(nameOfStruct) == 0 {
		fmt.Println("the name of struct can't be empty!")
		return
	}

	jsonStructFormat := `json:"`

	data := `package main
	
	func Anonymous(){
	type `

	res_1 := unicode.IsUpper(rune(nameOfStruct[0]))
	if !res_1 {

		data += strings.ToUpper(string(nameOfStruct[0])) + string(nameOfStruct[1:len(nameOfStruct)])
	} else {
		data += nameOfStruct
	}
	data += ` struct{
	`
	for key, val := range dataMap {

		switch v := val.(type) {
		case int:
			// v is an int here, so e.g. v + 1 is possible.
			data += strings.ToUpper(string(key[0])) + string(key[1:len(key)])
			data += ` int ` + "`" + jsonStructFormat + strings.ToLower(string(key[0])) + string(key[1:len(key)]) + `"` + "`\n"
			fmt.Printf(key+":Integer: %v\n", v)
		case float64:
			data += strings.ToUpper(string(key[0])) + string(key[1:len(key)])
			data += ` float64 ` + "`" + jsonStructFormat + strings.ToLower(string(key[0])) + string(key[1:len(key)]) + `"` + "`\n"
			// v is a float64 here, so e.g. v + 1.0 is possible.
			fmt.Printf(key+":Float64: %v\n", v)
		case string:
			data += strings.ToUpper(string(key[0])) + string(key[1:len(key)])
			data += ` string ` + "`" + jsonStructFormat + strings.ToLower(string(key[0])) + string(key[1:len(key)]) + `"` + "`\n"
			// v is a string here, so e.g. v + " Yeah!" is possible.
			fmt.Printf(key+":String: %v\n", v)
		default:
			data += strings.ToUpper(string(key[0])) + string(key[1:len(key)])
			data += ` interface{} ` + "`" + jsonStructFormat + strings.ToLower(string(key[0])) + string(key[1:len(key)]) + `"` + "`\n"
			// And here I'm feeling dumb. ;)
			fmt.Printf(key + ":I don't know, ask stackoverflow.\n")
		}
	}

	data += `
	}
	}`
	_ = ioutil.WriteFile(nameOfFile+".go", []byte(data), 0644)
}
