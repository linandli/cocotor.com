package torchlib

import (
	"fmt"
	"reflect"
	//"cocotor.com/utils/torchlib"
)

func GetPlants() map[string]string {
//http.get(['org', 'plants'])
	res := Get("org/plants", nil, 60)
	fmt.Println("r----------->", reflect.TypeOf(res).String())
	m3 := make(map[string]string)
	return m3
}