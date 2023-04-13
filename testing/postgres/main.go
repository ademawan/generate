package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	pgUser_test   = "postgres"
	pgPass_test   = "X2023RoG@1"
	pgHost_test   = "103.13.207.248"
	pgPort_test   = "5432"
	pgDbname_test = "testing"
)

type Role struct {
	Id       int64
	Name     string
	KeyName  string
	IsActive bool
}

func main() {
	conn, err := Connect()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	ctx := context.TODO()
	// dataref := Role{}
	query := `SELECT * FROM role`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var dataa []interface{}

	for rows.Next() {
		err := rows.Scan(&dataa)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(dataa)

	// datas, err := ScanAll(rows, &dataref)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(datas)
}

func Connect() (*pgxpool.Pool, error) {
	configPGUser := pgUser_test
	configPGPass := pgPass_test
	configPGHost := pgHost_test
	configPGPort := pgPort_test
	configPGDbname := pgDbname_test
	url := "postgres://" + configPGUser + ":" + configPGPass + "@" + configPGHost + ":" + configPGPort + "/" + configPGDbname
	return pgxpool.Connect(context.Background(), url)
}

func ScanAll(rows pgx.Rows, structReference interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(structReference)
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("must pass a pointer, not a value, to StructScan destination") // @todo add new error message
	}
	reflectValue := v.Elem()
	var reflectType = reflectValue.Type()
	// v = reflect.Indirect(v)
	var queryField = make([]string, 0)
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("test")

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

		}
	}
	cols := queryField

	data := make([]interface{}, 0)
	for rows.Next() {
		var m map[string]interface{}
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m = make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		data = append(data, m)
	}
	// for i := 0; i < v.NumField(); i++ {
	// 	field := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]

	// 	if item, ok := m[field]; ok {
	// 		if v.Field(i).CanSet() {
	// 			if item != nil {
	// 				switch v.Field(i).Kind() {
	// 				case reflect.String:
	// 					v.Field(i).SetString(string(item.([]uint8)))
	// 				case reflect.Float32, reflect.Float64:
	// 					v.Field(i).SetFloat(item.(float64))
	// 				case reflect.Ptr:
	// 					if reflect.ValueOf(item).Kind() == reflect.Bool {
	// 						itemBool := item.(bool)
	// 						v.Field(i).Set(reflect.ValueOf(&itemBool))
	// 					}
	// 				case reflect.Struct:
	// 					v.Field(i).Set(reflect.ValueOf(item))
	// 				default:
	// 					fmt.Println(t.Field(i).Name, ": ", v.Field(i).Kind(), " - > - ", reflect.ValueOf(item).Kind()) // @todo remove after test out the Get methods
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return data, nil
}

func ScanAllTagJson(rows pgx.Rows, structReference interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(structReference)
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("must pass a pointer, not a value, to StructScan destination") // @todo add new error message
	}
	reflectValue := v.Elem()
	var reflectType = reflectValue.Type()
	// v = reflect.Indirect(v)
	var queryField = make([]string, 0)
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("json")

		queryField = append(queryField, val)
	}
	cols := queryField

	data := make([]interface{}, 0)
	for rows.Next() {
		var m map[string]interface{}
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m = make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		data = append(data, m)
	}
	// for i := 0; i < v.NumField(); i++ {
	// 	field := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]

	// 	if item, ok := m[field]; ok {
	// 		if v.Field(i).CanSet() {
	// 			if item != nil {
	// 				switch v.Field(i).Kind() {
	// 				case reflect.String:
	// 					v.Field(i).SetString(string(item.([]uint8)))
	// 				case reflect.Float32, reflect.Float64:
	// 					v.Field(i).SetFloat(item.(float64))
	// 				case reflect.Ptr:
	// 					if reflect.ValueOf(item).Kind() == reflect.Bool {
	// 						itemBool := item.(bool)
	// 						v.Field(i).Set(reflect.ValueOf(&itemBool))
	// 					}
	// 				case reflect.Struct:
	// 					v.Field(i).Set(reflect.ValueOf(item))
	// 				default:
	// 					fmt.Println(t.Field(i).Name, ": ", v.Field(i).Kind(), " - > - ", reflect.ValueOf(item).Kind()) // @todo remove after test out the Get methods
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	return data, nil
}

func ScanAllx(rows pgx.Rows, structPointerReference interface{}) ([]interface{}, error) {

	data := make([]interface{}, 0)

	v := reflect.ValueOf(structPointerReference)
	if v.Kind() != reflect.Ptr {
		return data, errors.New("pointer is required")
	}
	reflectValue := v.Elem()
	var reflectType = reflectValue.Type()
	var queryField = make([]string, 0)
	for i := 0; i < reflectValue.NumField(); i++ {

		field := reflectType.Field(i).Name
		f, _ := reflectType.FieldByName(field)
		val, _ := f.Tag.Lookup("json")

		queryField = append(queryField, val)

	}

	for rows.Next() {
		var m map[string]interface{}
		columns := make([]interface{}, len(queryField))
		columnPointers := make([]interface{}, len(queryField))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return data, err
		}

		m = make(map[string]interface{})
		for i, colName := range queryField {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		data = append(data, m)
	}

	return data, nil
}
