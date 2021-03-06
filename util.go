package zendeskapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
	"strings"
)

func StructToSchema(s interface{}) string {

	re := regexp.MustCompile(".*\\.")
	name := re.ReplaceAllString(reflect.TypeOf(s).String(), "")

	schema := "{\"data\": {\"key\":" + fmt.Sprintf("\"%s\",\"schema\": {\"properties\": {", name)

	v := reflect.ValueOf(s)

	properties := make([]string, 0)

	for i := 0; i < v.NumField(); i++ {

		if v.Field(i).Kind() == reflect.Struct {
			continue
		}

		var t string

		switch v.Field(i).Kind() {
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr,
			reflect.Float32,
			reflect.Float64:
			t = "number"
		case reflect.Bool:
			t = "boolean"
		default:
			t = "string"

		}

		property := fmt.Sprintf("\"%s\":{\"type\":\"%s\",\"description\": \"autogenerated\"}", v.Type().Field(i).Name, t)

		properties = append(properties, property)
	}

	schema += strings.Join(properties, ",") + "}}}}"

	schema = strings.ToLower(schema)

	return schema

}

func printPrettyResponse(r *http.Response) {

	d, _ := httputil.DumpResponse(r, true)

	log.Println(string(d))

}

func printPrettyRequest(r *http.Request) {

	d, _ := httputil.DumpRequest(r, true)

	log.Println(string(d))

}
func printPrettyStruct(d interface{}) {

	jsn, _ := json.MarshalIndent(d, " ", "")

	log.Println(string(jsn))

}

func bufferJSON(s interface{}) (*bytes.Buffer, error) {

	j, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer([]byte(j)), nil

}
