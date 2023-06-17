package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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
type Tc struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Alamat string `json:",omitempty"`
}

// type ResponseSubscriberProfile struct {
// 	Profiles Profiles `json:"profile"`
// }

func main() {

	// e := Extended{}
	// e.Email = "me@mail.com"
	// e.Password = "secret"

	// for i := 0; i < reflect.TypeOf(e).NumField(); i++ {
	// 	if reflect.ValueOf(e).Field(i).Kind() != reflect.Struct {
	// 		fmt.Println(reflect.ValueOf(e).Field(i))
	// 	}
	// }

	// fmt.Println("hallp")
	// // getPropertyInfo(&UserRequestCreate{})
	// // generateQueryInsert(&UserRequestCreate{})
	// generateQueryInsert(&UserRequestCreate{}, "query_insert", "user")

	// var data2 interface{}
	// GetFile("", "query_insert", "go", &data2)
	// fmt.Println(data2)
	// te := strings.Split(data2.(string), "//")
	// fmt.Println(te[0])

	var data []byte
	GetFileJson("", "file", "json", &data)
	var d []byte

	GetFileJson("", "tc", "json", &d)

	// fmt.Println(data)
	dataMap := make(map[string]interface{})
	tc := Tc{}
	json.Unmarshal(d, &tc)
	fmt.Println(tc)
	// json.Unmarshal(data, &dataMap)
	// fmt.Println(dataMap)

	GenerateMapToStruct(dataMap, "coba", "test")

}
func getPropertyInfo(s *UserRequestCreate) {
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

func generateQueryInsert(data interface{}, nameOfFile, nameOfTable string) {
	var reflectValue = reflect.ValueOf(data)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	query := `INSERT INTO `
	query += nameOfTable + `(`
	var queryValues = make([]string, 0)
	var queryField = make([]string, 0)
	goFieldPointerReference := make([]string, 0)
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("test")

		//get pointer of field
		if val != "-" {
			goFieldPointerReference = append(goFieldPointerReference, nameOfTable+`.`+reflectType.Field(i).Name)
		}
		//
		newField := ""
		if val != "-" {
			for j := 0; j < len(field); j++ {

				res_1 := unicode.IsUpper(rune(field[j]))
				if res_1 {
					if res_1 && j != 0 {
						newField += "_"
					}
					newField += strings.ToLower(string(field[j]))
				} else {
					newField += string(field[j])
				}

			}
			queryField = append(queryField, newField)
			queryValues = append(queryValues, "?")
		}
	}
	query += strings.Join(queryField, ",")
	query += `)VALUES(`
	query += strings.Join(queryValues, ",")
	query += `)`
	fmt.Println(query)
	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

	goFieldPointerReferenceString := `
	package main
	//
	func Create(){
	query := ` + `"` + query + `"
	`
	goFieldPointerReferenceString += `stmt, err := r.db.Prepare(query)
	if err != nil {
		return err, r.log

	}
	_, err = stmt.Exec(`
	goFieldPointerReferenceString += strings.Join(goFieldPointerReference, ",")
	goFieldPointerReferenceString += `)
	`
	goFieldPointerReferenceString += `if err != nil {
		r.log.Message += "|Exec|" + err.Error()
		return err, r.log
	}
	
	}`
	_ = ioutil.WriteFile(nameOfFile+".go", []byte(goFieldPointerReferenceString), 0644)
	// fmt.Println(goFieldPointerReferenceString)
}

func generateQueryInsertFileGo(data interface{}, nameOfFile, nameOfTable string) {
	var reflectValue = reflect.ValueOf(data)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	query := `INSERT INTO `
	query += nameOfTable + `(`
	var queryValues = make([]string, 0)
	var queryField = make([]string, 0)

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("test")

		newField := ""
		if val != "-" {
			for i := 0; i < len(field); i++ {
				res_1 := unicode.IsUpper(rune(field[i]))
				if res_1 {
					if res_1 && i != 0 {
						newField += "_"
					}
					newField += strings.ToLower(string(field[i]))
				} else {
					newField += string(field[i])
				}

			}

			queryField = append(queryField, newField)
			queryValues = append(queryValues, "?")
		}
	}
	query += strings.Join(queryField, ",")
	query += `)VALUES(`
	query += strings.Join(queryValues, ",")
	query += `)`
	fmt.Println(query)
	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

}

// func generateQueryInsert(data interface{}, nameOfFile, nameOfTable string) {
// 	var reflectValue = reflect.ValueOf(data)

// 	if reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}

// 	var reflectType = reflectValue.Type()
// 	query := `INSERT INTO `
// 	query += nameOfTable + `(`
// 	var queryValues = make([]string, 0)
// 	var queryField = make([]string, 0)
// 	goFieldPointerReference := make([]string, 0)
// 	for i := 0; i < reflectValue.NumField(); i++ {
// 		field := reflectType.Field(i).Name
// 		f, _ := reflectType.FieldByName(field)
// 		val, _ := f.Tag.Lookup("test")

// 		//get pointer of field
// 		goFieldPointerReference = append(goFieldPointerReference, `&`+nameOfTable+`.`+reflectType.Field(i).Name)
// 		//
// 		newField := ""
// 		if val != "-" {
// 			for j := 0; j < len(field); j++ {

// 				res_1 := unicode.IsUpper(rune(field[j]))
// 				if res_1 {
// 					if res_1 && j != 0 {
// 						newField += "_"
// 					}
// 					newField += strings.ToLower(string(field[j]))
// 				} else {
// 					newField += string(field[j])
// 				}

// 			}
// 			queryField = append(queryField, newField)
// 			queryValues = append(queryValues, "?")
// 		}
// 	}
// 	query += strings.Join(queryField, ",")
// 	query += `)VALUES(`
// 	query += strings.Join(queryValues, ",")
// 	query += `)`
// 	fmt.Println(query)
// 	_ = ioutil.WriteFile(nameOfFile+".sql", []byte(query), 0644)

// 	goFieldPointerReferenceString := strings.Join(goFieldPointerReference, ",")
// 	_ = ioutil.WriteFile(nameOfFile+".go", []byte(goFieldPointerReferenceString), 0644)

// }
func GetFile(baseDir, fileName, ext string, data *interface{}) {
	jsonFile, err := os.Open(baseDir + fileName + "." + ext)
	if err != nil {

		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	// json.Unmarshal(byteValue, &data)
	fmt.Println("hallo")
	*data = string(byteValue)

}
func GetFileJson(baseDir, fileName, ext string, data *[]byte) {
	jsonFile, err := os.Open(baseDir + fileName + "." + ext)
	if err != nil {

		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	// json.Unmarshal(byteValue, &data)
	*data = byteValue

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
