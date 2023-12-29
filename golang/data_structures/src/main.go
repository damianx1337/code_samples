package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

func main() {
    MyMap := make(map[string][]map[string]string)
		details := make(map[string]string)
		details["prop_id"] = "tesstprop"
		details["rep_id"] = "tesstrep"
		details["ad_time"] = "tesstad"

    MyMap["key1"] = append(MyMap["key1"], details)
    MyMap["key1"] = append(MyMap["key1"], details)
    MyMap["key1"] = append(MyMap["key1"], details)
    MyMap["key2"] = append(MyMap["key1"], details)
    fmt.Println(MyMap["key1"])
    fmt.Println(reflect.TypeOf(MyMap["key1"]))

		j, err := json.Marshal(MyMap)
    if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(j))
		}

}
