package main

import (
	"fmt"
	"reflect"
	"encoding/json"
)

type Example struct {
    Id []int
    Name []string
}

func (data *Example) AppendOffer(id int, name string) {
    data.Id = append(data.Id, id)
    data.Name = append(data.Name, name)
}

var MyMap map[string]*Example

func main() {
    obj := &Example{[]int{}, []string{}}
    obj.AppendOffer(1, "SomeText")
    MyMap = make(map[string]*Example)
    MyMap["key1"] = obj
    MyMap["key2"] = obj
    MyMap["key1"].Name = append(MyMap["key1"].Name, "vlavla")
    fmt.Println(MyMap["key1"].Name)
    fmt.Println(reflect.TypeOf(MyMap["key1"]))

		j, err := json.Marshal(MyMap)
    if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(j))
		}
}
