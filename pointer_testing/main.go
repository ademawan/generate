package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DetailTransaction struct {
	Month       string `json:"month"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
type Response struct {
	Code int                 `json:"code"`
	Data []DetailTransaction `json:"data"`
}

func main() {
	data1 := &DetailTransaction{}

	data2 := &DetailTransaction{}
	data2.Description = "halo"
	datas := []DetailTransaction{}
	if data1 != nil {
		datas = append(datas, *data1)

	}
	userData := new(DetailTransaction)
	userData2 := new(DetailTransaction)
	userData.Description = "hallo"
	userData2.Description = "halo2"
	res := &Response{}
	res.Code = 200
	res.Data = datas
	fmt.Println(&data1, &data2, data1.Description, data2.Description)
	fmt.Println(&userData, &userData2, userData.Description, userData2.Description)
	WriteFile("file1.json", data1)
	WriteFile("file2.json", res)
	

}

func WriteFile(fileName string, data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fileName, file, 0644)
}
